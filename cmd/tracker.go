package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"net"
	"time"

	"github.com/zaninime/btracker/db"
	"github.com/zaninime/btracker/udp"
)

func listenUDP(conn *net.UDPConn) {
	for {
		mainLogger.Debug("waiting for packet")
		pkt := make([]byte, 2000)
		n, addr, err := conn.ReadFromUDP(pkt)
		if err != nil {
			mainLogger.Error("cannot read from socket", "err", err)
		}
		mainLogger.Debug("got packet", "addr", addr, "len", n)
		go processPkt(conn, addr, pkt[:n])
	}
}

func processPkt(conn *net.UDPConn, addr *net.UDPAddr, data []byte) {
	pv := udp.Parse(data)
	if pv == nil {
		mainLogger.Debug("invalid packet", "addr", addr)
		return
	}
	switch pv.Action {
	case udp.ActionConnectionRequest:
		mainLogger.Debug("processing connection request")
		processConnectionRequest(conn, addr, pv)
	case udp.ActionAnnounce:
		mainLogger.Debug("processing announce")
		processAnnounceRequest(conn, addr, pv)
	case udp.ActionScrape:
		mainLogger.Debug("processing scrape")
		processScrapeRequest(conn, addr, pv)
	}
}

func processConnectionRequest(conn *net.UDPConn, addr *net.UDPAddr, pv *udp.ProtocolVars) {
	response := udp.ConnectionRequestResponse{pv.TransactionID, nil}
	last, ok := lastConnectionRequest[addr.IP.String()]
	if !ok || last.Add(6*time.Second).Before(time.Now()) {
		lastConnectionRequest[addr.IP.String()] = time.Now()
		connID := make([]byte, 8)
		rand.Read(connID)
		expiryTime := time.Now().Add(2 * time.Minute)
		mainLogger.Debug("executing query", "q", db.StmtWriteConnectionID.String, "ip", addr.IP, "id", connID, "expiry", expiryTime)
		if _, err := db.StmtWriteConnectionID.Stmt.Exec(base64.StdEncoding.EncodeToString(connID), addr.IP.String(), expiryTime); err != nil {
			mainLogger.Error("couldn't write connection id to the database", "err", err)
		}
		response.ConnectionID = connID
		// TODO: improve here
		n, _ := conn.WriteToUDP(response.Accept(), addr)
		mainLogger.Debug("responding normally", "len", n)
		return
	}
	conn.WriteToUDP(response.Error("Too fast, try again later"), addr)
}

func processAnnounceRequest(conn *net.UDPConn, addr *net.UDPAddr, pv *udp.ProtocolVars) {
	mainLogger.Debug("got an announce request", "torrent", pv.InfoHashes[0], "peer", pv.PeerID)
	response := udp.AnnounceResponse{pv.TransactionID, int32(config.Tracker.Interval), 0, 0, []udp.Peer{}}
	status, err := db.CheckConnectionID(pv.ConnectionID, addr.IP)
	if err != nil {
		mainLogger.Debug("responding with error", "msg", "Internal error")
		conn.WriteToUDP(response.Error("Internal error"), addr)
		return
	}
	if !status {
		mainLogger.Debug("connection ID is invalid, responding with error", "msg", "Invalid connection ID")
		conn.WriteToUDP(response.Error("Invalid connection ID"), addr)
		return
	}
	mainLogger.Debug("connection ID is valid")

	if pv.RequestedResults == -1 {
		pv.RequestedResults = 2147483647
	} else if pv.RequestedResults > 1 {
		mainLogger.Debug("requesting peers data from database")
		if err := db.PopulatePeersFields(&response, pv.PeerID, pv.InfoHashes[0], pv.RequestedResults); err != nil {
			mainLogger.Debug("responding with error", "msg", "Internal error")
			conn.WriteToUDP(response.Error("Internal error"), addr)
		}
	}
	n, err := conn.WriteToUDP(response.Accept(), addr)
	if err != nil {
		mainLogger.Error("couldn't write to socket", "err", err)
	}
	mainLogger.Debug("responding normally", "len", n)

	// update database
	mainLogger.Debug("getting peer from database")
	var peer *db.Peer
	peer, err = db.GetPeer(pv.PeerID, pv.InfoHashes[0])
	if err == sql.ErrNoRows || (peer == nil && err == nil) {
		// peer not found inside database, creating one
		mainLogger.Debug("peer not present, creating")
		peerStates := map[udp.Event]int{
			udp.EventStarted:   1,
			udp.EventStopped:   0,
			udp.EventCompleted: 2,
			udp.EventNone:      1,
		}
		if pv.IPAddress == nil {
			pv.IPAddress = addr.IP
		}
		peer = &db.Peer{
			ID:          pv.PeerID,
			TorrentID:   pv.InfoHashes[0],
			State:       peerStates[pv.Event],
			IP:          pv.IPAddress.String(),
			Port:        pv.Port,
			Downloaded:  pv.DownloadedBytes,
			Uploaded:    pv.UploadedBytes,
			Left:        pv.LeftBytes,
			LastUpdated: time.Now(),
		}
		mainLogger.Debug("adding peer")
		if err2 := db.InsertPeer(peer); err2 != nil {
			return
		}
		db.IncrementTorrentDownloadedStats(pv.InfoHashes[0])
		if pv.Event == udp.EventCompleted {
			db.IncrementTorrentCompletedStats(pv.InfoHashes[0])
		}
		return
	} else if err != nil {
		return
	}
	peerStates := map[udp.Event]int{
		udp.EventStarted:   1,
		udp.EventStopped:   0,
		udp.EventCompleted: 2,
		udp.EventNone:      peer.State,
	}
	peer.State = peerStates[pv.Event]
	peer.Downloaded = pv.DownloadedBytes
	peer.Uploaded = pv.UploadedBytes
	peer.Left = pv.LeftBytes
	mainLogger.Debug("updating peer")
	if err := db.UpdatePeer(peer); err != nil {
		return
	}
	if pv.Event == udp.EventCompleted {
		db.IncrementTorrentCompletedStats(pv.InfoHashes[0])
	}
}

func processScrapeRequest(conn *net.UDPConn, addr *net.UDPAddr, pv *udp.ProtocolVars) {
	response := udp.ScrapeResponse{pv.TransactionID, make([]udp.TorrentStats, 0, len(pv.InfoHashes))}
	for _, hash := range pv.InfoHashes {
		var stat *udp.TorrentStats
		var err error
		if stat, err = db.GetTorrentStats(hash); err != nil {
			mainLogger.Debug("responding with error", "msg", "Internal error")
			conn.WriteToUDP(response.Error("Internal error"), addr)
			return
		}
		response.Stats = append(response.Stats, *stat)
	}
	n, err := conn.WriteToUDP(response.Accept(), addr)
	if err != nil {
		mainLogger.Error("couldn't write to socket", "err", err)
		return
	}
	mainLogger.Debug("responding normally", "len", n)
}
