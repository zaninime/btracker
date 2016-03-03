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
		conn.WriteToUDP(response.Accept(), addr)
		return
	}
	conn.WriteToUDP(response.Error("Too fast, try again later"), addr)
}

func processAnnounceRequest(conn *net.UDPConn, addr *net.UDPAddr, pv *udp.ProtocolVars) {
	mainLogger.Debug("got an announce request", "torrent", pv.InfoHashes[0], "peer", pv.PeerID)
	response := udp.AnnounceResponse{pv.TransactionID, int32(config.Tracker.Interval), 0, 0, nil}
	status, err := db.CheckConnectionID(pv.ConnectionID, addr.IP)
	if err != nil {
		mainLogger.Error("couldn't query the database for connection id", "err", err)
		conn.WriteToUDP(response.Error("Internal error"), addr)
		return
	}
	if !status {
		mainLogger.Trace("connection id is invalid")
		conn.WriteToUDP(response.Error("Invalid connection id"), addr)
		return
	}
	mainLogger.Trace("connection id is valid")
	// process normally
	// fetch and return data, then write to db
	if pv.RequestedResults == -1 {
		pv.RequestedResults = 2147483647
	}
	result, err := db.StmtGetPeers.Stmt.Query(base64.StdEncoding.EncodeToString(pv.PeerID), base64.StdEncoding.EncodeToString(pv.InfoHashes[0]), pv.RequestedResults)
	defer result.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			response.Peers = []udp.Peer{}
			conn.WriteToUDP(response.Accept(), addr)
			return
		}
		mainLogger.Error("couldn't query the database for peers", "err", err)
		conn.WriteToUDP(response.Error("Internal error"), addr)
		return
	}
	response.Peers = []udp.Peer{}
	for result.Next() {
		var ipString string
		var port uint16
		err := result.Scan(&ipString, &port)
		if err != nil {
			mainLogger.Error("couldn't scan results", "err", err)
		}
		ip := net.ParseIP(ipString)
		response.Peers = append(response.Peers, udp.Peer{IP: ip, Port: port})
	}
	if err := db.StmtGetLeecherPeers.Stmt.QueryRow(base64.StdEncoding.EncodeToString(pv.InfoHashes[0])).Scan(&response.Leechers); err != nil {
		mainLogger.Error("couldn't query the database for leechers", "err", err)
		conn.WriteToUDP(response.Error("Internal error"), addr)
		return
	}
	if err := db.StmtGetSeederPeers.Stmt.QueryRow(base64.StdEncoding.EncodeToString(pv.InfoHashes[0])).Scan(&response.Seeders); err != nil {
		mainLogger.Error("couldn't query the database for seeders", "err", err)
		conn.WriteToUDP(response.Error("Internal error"), addr)
		return
	}
	conn.WriteToUDP(response.Accept(), addr)

	// update database
	var peer struct {
		ID          []byte    `db:"id"`
		TorrentID   []byte    `db:"torrent_id"`
		State       int       `db:"state"`
		IP          string    `db:"ip"`
		Port        uint16    `db:"port"`
		Downloaded  int32     `db:"downloaded"`
		Uploaded    int32     `db:"uploaded"`
		Left        int32     `db:"left"`
		LastUpdated time.Time `db:"last_updated"`
	}
	if err := db.StmtGetPeer.Stmt.QueryRowx(base64.StdEncoding.EncodeToString(pv.PeerID), base64.StdEncoding.EncodeToString(pv.InfoHashes[0])).StructScan(&peer); err != nil {
		if err == sql.ErrNoRows {
			mainLogger.Debug("new peer, running query", "q", db.StmtInsertNewPeer.String)
			var peerState int
			switch pv.Event {
			case udp.EventStarted:
				peerState = 1
			case udp.EventStopped:
				peerState = 0
			case udp.EventCompleted:
				peerState = 2
			default:
				peerState = 1
			}
			if pv.IPAddress == nil {
				pv.IPAddress = addr.IP
			}
			if pv.Port == 0 {
				pv.Port = uint16(addr.Port)
			}
			if _, err2 := db.StmtInsertNewPeer.Stmt.Exec(base64.StdEncoding.EncodeToString(pv.PeerID), base64.StdEncoding.EncodeToString(pv.InfoHashes[0]), peerState, pv.IPAddress, pv.Port, pv.DownloadedBytes, pv.UploadedBytes, pv.LeftBytes); err2 != nil {
				mainLogger.Error("couldn't insert peer into database", "err", err)
			}
			return
		}
		mainLogger.Error("couldn't query database for peer", "err", err)
		return
	}
	switch pv.Event {
	case udp.EventStarted:
		peer.State = 1
	case udp.EventStopped:
		peer.State = 0
	case udp.EventCompleted:
		peer.State = 2
	}
	if _, err := db.StmtUpdatePeer.Stmt.Exec(base64.StdEncoding.EncodeToString(peer.ID), base64.StdEncoding.EncodeToString(peer.TorrentID), peer.State, peer.IP, peer.Port, peer.Downloaded, peer.Uploaded, peer.Left); err != nil {
		mainLogger.Error("couldn't update peer inside database", "err", err)
		return
	}
}
