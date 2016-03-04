package udp

import (
	"encoding/binary"
	"net"
)

type ConnectionRequestResponse struct {
	TransactionID []byte
	ConnectionID  []byte
}

type AnnounceResponse struct {
	TransactionID []byte
	Interval      int32
	Leechers      int32
	Seeders       int32
	Peers         []Peer
}

type ScrapeResponse struct {
	TransactionID []byte
	Stats         []TorrentStats
}

type TorrentStats struct {
	Complete   int32
	Downloaded int32
	Incomplete int32
}

type Peer struct {
	IP   net.IP
	Port uint16
}

func (crr *ConnectionRequestResponse) Accept() []byte {
	pkt := make([]byte, 16)
	binary.BigEndian.PutUint32(pkt, 0)
	copy(pkt[4:8], crr.TransactionID)
	copy(pkt[8:16], crr.ConnectionID)
	return pkt
}

func (crr *ConnectionRequestResponse) Error(msg string) []byte {
	return makeGenericErrorPacket(msg, crr.TransactionID)
}

func (ar *AnnounceResponse) Accept() []byte {
	exceed := 6 * len(ar.Peers)
	pkt := make([]byte, 20+exceed)
	binary.BigEndian.PutUint32(pkt, 1)
	copy(pkt[4:8], ar.TransactionID)
	binary.BigEndian.PutUint32(pkt[8:12], uint32(ar.Interval))
	binary.BigEndian.PutUint32(pkt[12:16], uint32(ar.Leechers))
	binary.BigEndian.PutUint32(pkt[16:20], uint32(ar.Seeders))
	var start int
	for i, p := range ar.Peers {
		start = 20 + i*6
		copy(pkt[start:start+4], p.IP.To4())
		binary.BigEndian.PutUint16(pkt[start+4:start+6], p.Port)
	}
	return pkt
}

func (ar *AnnounceResponse) Error(msg string) []byte {
	return makeGenericErrorPacket(msg, ar.TransactionID)
}

func (sr *ScrapeResponse) Accept() []byte {
	pkt := make([]byte, 8+len(sr.Stats)*12)
	binary.BigEndian.PutUint32(pkt, 2)
	copy(pkt[4:8], sr.TransactionID)
	for i, s := range sr.Stats {
		start := 8 + 12*i
		binary.BigEndian.PutUint32(pkt[start:start+4], uint32(s.Complete))
		binary.BigEndian.PutUint32(pkt[start+4:start+8], uint32(s.Downloaded))
		binary.BigEndian.PutUint32(pkt[start+8:start+12], uint32(s.Incomplete))
	}
	return pkt
}

func (sr *ScrapeResponse) Error(msg string) []byte {
	return makeGenericErrorPacket(msg, sr.TransactionID)
}

func makeGenericErrorPacket(msg string, transactionID []byte) []byte {
	pkt := make([]byte, 8)
	binary.BigEndian.PutUint32(pkt, 3)
	copy(pkt[4:8], transactionID)
	pkt = append(pkt, []byte(msg)...)
	return pkt
}
