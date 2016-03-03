package udp

func checkConnectionRequestPkt(pv *ProtocolVars) *ProtocolVars {
	// nothing to check, not used
	return pv
}

func checkAnnouncePkt(pv *ProtocolVars) *ProtocolVars {
	//fmt.Printf("%#v\n", pv)
	if pv.DownloadedBytes < 0 || pv.UploadedBytes < 0 || pv.LeftBytes < 0 {
		return nil
	}

	if pv.RequestedResults < 1 && pv.RequestedResults != -1 {
		return nil
	}

	// cesco: not super sure about this, maybe it's "use the port you got the packet from"
	if pv.Port == 0 {
		return nil
	}
	return pv
}

func checkScrapePkt(pv *ProtocolVars) *ProtocolVars {
	// nothing to check, not used
	return pv
}
