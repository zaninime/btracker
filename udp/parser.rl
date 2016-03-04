package udp

import (
  //"fmt"
  "net"
  "bytes"
  "encoding/binary"
)

%%{
machine torrent;

protocol_id = 0 . 0 . 0x4 . 0x17 . 0x27 . 0x10 . 0x19 . 0x80;

actionf_connect = 0 . 0 . 0 . 0 %{
  protocolVars.Action = ActionConnectionRequest
};
actionf_announce = 0 . 0 . 0 . 1 %{
  protocolVars.Action = ActionAnnounce
};
actionf_scrape = 0 . 0 . 0 . 2 %{
  protocolVars.Action = ActionScrape;
};

action append_info_hash {
  if protocolVars.InfoHashes == nil {
    protocolVars.InfoHashes = make([][]byte, 0, 74)
  }
  protocolVars.InfoHashes = append(protocolVars.InfoHashes, data[p-20:p])
}

action mark_valid {
  valid = true
}

transaction_id = any{4} %{
  protocolVars.TransactionID = data[p-4:p]
};
connection_id = any{8} %{
  protocolVars.ConnectionID = data[p-8:p]
};
info_hash = any{20} %{
  protocolVars.InfoHashes = [][]byte{data[p-20:p]}
};
peer_id = any{20} %{
  protocolVars.PeerID = data[p-20:p]
};
downloaded_bytes = any{8} %{
  buf = bytes.NewReader(data[p-8:])
  err = binary.Read(buf, binary.BigEndian, &protocolVars.DownloadedBytes)
	if err != nil {
		panic(err)
	}
};
left_bytes = any{8} %{
  err = binary.Read(buf, binary.BigEndian, &protocolVars.LeftBytes)
  if err != nil {
    panic(err)
  }
};
uploaded_bytes = any{8} %{
  err = binary.Read(buf, binary.BigEndian, &protocolVars.UploadedBytes)
  if err != nil {
    panic(err)
  }
};
event = 0 . 0 . 0 . (0 | 1 | 2 | 3) %{
  switch data[p-1] {
  case 0:
    protocolVars.Event = EventNone
  case 1:
    protocolVars.Event = EventCompleted
  case 2:
    protocolVars.Event = EventStarted
  case 3:
    protocolVars.Event = EventStopped
  }
};
ip_address = any{4} %{
  buf = bytes.NewReader(data[p-4:])
  err = binary.Read(buf, binary.BigEndian, &intIP)
  if err != nil {
    panic(err)
  }
  if intIP != 0 {
    protocolVars.IPAddress = net.IPv4(data[p-4], data[p-3], data[p-2], data[p-1])
  }
};
client_key = any{4} %{
  protocolVars.Key = data[p-4:p]
};
results_count = any{4} %{
  buf = bytes.NewReader(data[p-4:])
  err = binary.Read(buf, binary.BigEndian, &protocolVars.RequestedResults)
  if err != nil {
    panic(err)
  }
};
port = any{2} %{
  err = binary.Read(buf, binary.BigEndian, &protocolVars.Port)
  if err != nil {
    panic(err)
  }
};
extensions = any{2}?;

connection_request = protocol_id . actionf_connect . transaction_id;
announce = connection_id . actionf_announce . transaction_id . info_hash . peer_id . downloaded_bytes . left_bytes . uploaded_bytes . event . ip_address . client_key . results_count . port . extensions . any*;
scrape = connection_id . actionf_scrape . transaction_id . (any{20} %append_info_hash)+;

main := (connection_request | announce | scrape) %/mark_valid;
}%%

%% write data;

func Parse(data []byte) *ProtocolVars {
  p := 0
  pe := len(data)
  eof := pe
  var cs int

  var buf *bytes.Reader
  var intIP int32
  var err error
  valid := false

  protocolVars := ProtocolVars{}

  %% write init;
  %% write exec;

  //fmt.Printf("%+v\n", protocolVars)

  if !valid {
    return nil
  }

  switch protocolVars.Action {
  case ActionConnectionRequest:
    return &protocolVars
    //return checkConnectionRequestPkt(&protocolVars)
  case ActionAnnounce:
    return checkAnnouncePkt(&protocolVars)
  case ActionScrape:
    return &protocolVars
    //return checkScrapePkt(&protocolVars)
  }

  // should be useless
  return nil
}
