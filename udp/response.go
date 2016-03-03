package udp

import "encoding/binary"

type ConnectionRequestResponse struct {
	TransactionID []byte
}

func (crr *ConnectionRequestResponse) Accept(connID []byte) []byte {
	pkt := make([]byte, 4, 16)
	binary.BigEndian.PutUint32(pkt, 0)
	pkt = append(pkt, crr.TransactionID...)
	pkt = append(pkt, connID...)
	return pkt
}

func (crr *ConnectionRequestResponse) Error(msg string) []byte {
	pkt := make([]byte, 4, 8)
	binary.BigEndian.PutUint32(pkt, 3)
	pkt = append(pkt, crr.TransactionID...)
	pkt = append(pkt, []byte(msg)...)
	return pkt
}
