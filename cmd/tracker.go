package main

import (
	"crypto/rand"
	"net"
	"time"

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
	}
	switch pv.Action {
	case udp.ActionConnectionRequest:
		mainLogger.Debug("processing connection request")
		processConnectionRequest(conn, addr, pv)
	}
}

func processConnectionRequest(conn *net.UDPConn, addr *net.UDPAddr, pv *udp.ProtocolVars) {
	response := udp.ConnectionRequestResponse{pv.TransactionID}
	last, ok := lastConnectionRequest[addr.IP.String()]
	if !ok || last.Add(6*time.Second).Before(time.Now()) {
		lastConnectionRequest[addr.IP.String()] = time.Now()
		connID := make([]byte, 8)
		rand.Read(connID)
		response := udp.ConnectionRequestResponse{pv.TransactionID}
		conn.WriteToUDP(response.Accept(connID), addr)
		return
	}
	conn.WriteToUDP(response.Error("Too fast, try again later"), addr)
}
