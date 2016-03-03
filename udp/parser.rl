package main

%%{
machine torrent;

protocol_id = 0 . 0 . 0x4 . 0x17 . 0x27 . 0x10 . 0x19 . 0x80;

actionf_connect = 0 . 0 . 0 . 0 %{
  action = ConnectionRequestAction
};
actionf_announce = 0 . 0 . 0 . 1 %{
  action = AnnounceAction;
};
actionf_scrape = 0 . 0 . 0 . 2 %{
  action = ScrapeAction;
  // setup info hash list here
};

action append_info_hash {}
action mark_valid {}

transaction_id = any{4} >{};
connection_id = any{8} >{};
info_hash = any{20} >{};
peer_id = any{20} >{};
downloaded_bytes = any{8} >{};
left_bytes = any{8} >{};
uploaded_bytes = any{8} >{};
event = 0 . 0 . 0 . (0 | 1 | 2 | 3) >{};
ip_address = any{4} >{};
client_key = any{4} >{};
results_count = any{4} >{};
port = any{2} >{};
extensions = any{2} >{};

connection_request = protocol_id . actionf_connect . transaction_id;
announce = connection_id . actionf_announce . transaction_id . info_hash . peer_id . downloaded_bytes . left_bytes . uploaded_bytes . event . ip_address . client_key . results_count . port . extensions . any*;
scrape = connection_id . actionf_scrape . transaction_id . (any{20} >append_info_hash)+;

main := (connection_request | announce | scrape) %/mark_valid;
}%%

%% write data;

func parse() {
  p := 0
  pe := 10
  eof := pe
  var cs int
  var data []byte

  var action int
  valid := false

  %% write init;
  %% write exec;
}
