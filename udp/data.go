package udp

import "net"

// Action is the UDP request type
type Action int

// Event represents the event reported by the torrent client
type Event int

// UDP request types
const (
	ActionConnectionRequest Action = iota
	ActionAnnounce
	ActionScrape
)

// Possible events
const (
	EventNone Event = iota
	EventStarted
	EventStopped
	EventCompleted
)

// ProtocolVars contains all the parsed fields of the procotol
type ProtocolVars struct {
	Action           Action
	TransactionID    []byte
	ConnectionID     []byte
	InfoHashes       [][]byte
	PeerID           []byte
	DownloadedBytes  int64
	LeftBytes        int64
	UploadedBytes    int64
	Event            Event
	IPAddress        net.IP
	Key              []byte
	RequestedResults int32
	Port             uint16
}
