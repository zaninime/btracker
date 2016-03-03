package udp

import "testing"

func TestParseConnectionRequest(t *testing.T) {
	pkt := []byte{0, 0, 0x4, 0x17, 0x27, 0x10, 0x19, 0x80, 0, 0, 0, 0, 1, 2, 3, 4}
	if Parse(pkt) == nil {
		t.Error("Expected a valid packet")
	}
}

func TestParseAnnounce(t *testing.T) {
	pkt := []byte{0, 1, 2, 3, 4, 5, 6, 7,
		0, 0, 0, 1,
		1, 2, 3, 4,
		5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
		9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 1, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 0,
		0, 0, 0, 2,
		0, 0, 0, 0,
		7, 7, 7, 7,
		255, 255, 255, 255,
		50, 50,
		0, 0,
	}
	if Parse(pkt) == nil {
		t.Error("Expected a valid packet")
	}
}

func TestParseScrape(t *testing.T) {
	pkt := []byte{0, 1, 2, 3, 4, 5, 6, 7,
		0, 0, 0, 2,
		1, 2, 3, 4,
		5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
		9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	}
	if Parse(pkt) == nil {
		t.Error("Expected a valid packet")
	}
}

func TestInvalid(t *testing.T) {
	pkt := []byte{1, 2, 3, 4, 5}
	if Parse(pkt) != nil {
		t.Error("Expected an invalid packet")
	}
}
