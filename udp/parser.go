
//line parser.rl:1
package udp

import (
  "fmt"
  "net"
  "bytes"
  "encoding/binary"
)


//line parser.rl:112



//line parser.go:18
const torrent_start int = 1
const torrent_first_final int = 138
const torrent_error int = 0

const torrent_en_main int = 1


//line parser.rl:115

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

  
//line parser.go:42
	{
	cs = torrent_start
	}

//line parser.rl:130
  
//line parser.go:49
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 0:
		goto st_case_0
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
	case 140:
		goto st_case_140
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	case 141:
		goto st_case_141
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	}
	goto st_out
	st_case_1:
		if data[p] == 0 {
			goto st2
		}
		goto st137
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 0 {
			goto st3
		}
		goto st136
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 4 {
			goto st123
		}
		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 0 {
			goto tr11
		}
		goto st0
tr11:
//line parser.rl:39

  protocolVars.ConnectionID = data[p-8:p]

	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line parser.go:414
		if data[p] == 0 {
			goto st11
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 0 {
			goto st12
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 1:
			goto st13
		case 2:
			goto st99
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		goto tr17
tr17:
//line parser.rl:18

  protocolVars.Action = ActionAnnounce

	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line parser.go:461
		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		goto tr21
tr21:
//line parser.rl:36

  protocolVars.TransactionID = data[p-4:p]

	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line parser.go:492
		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		goto tr41
tr41:
//line parser.rl:42

  protocolVars.InfoHashes = [][]byte{data[p-20:p]}

	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line parser.go:619
		goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		goto tr61
tr61:
//line parser.rl:45

  protocolVars.PeerID = data[p-20:p]

	goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line parser.go:746
		goto st59
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		goto st60
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		goto st62
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		goto st63
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		goto st64
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		goto st65
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		goto tr69
tr69:
//line parser.rl:48

  buf = bytes.NewReader(data[p-8:])
  err = binary.Read(buf, binary.BigEndian, &protocolVars.DownloadedBytes)
	if err != nil {
		panic(err)
	}

	goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line parser.go:805
		goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		goto tr77
tr77:
//line parser.rl:55

  err = binary.Read(buf, binary.BigEndian, &protocolVars.LeftBytes)
  if err != nil {
    panic(err)
  }

	goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line parser.go:863
		goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		if data[p] == 0 {
			goto tr85
		}
		goto st0
tr85:
//line parser.rl:61

  err = binary.Read(buf, binary.BigEndian, &protocolVars.UploadedBytes)
  if err != nil {
    panic(err)
  }

	goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line parser.go:924
		if data[p] == 0 {
			goto st83
		}
		goto st0
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		if data[p] == 0 {
			goto st84
		}
		goto st0
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		if data[p] <= 3 {
			goto st85
		}
		goto st0
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		goto tr89
tr89:
//line parser.rl:67

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

	goto st86
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
//line parser.go:973
		goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		goto st88
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		goto st89
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		goto tr93
tr93:
//line parser.rl:79

  buf = bytes.NewReader(data[p-4:])
  err = binary.Read(buf, binary.BigEndian, &intIP)
  if err != nil {
    panic(err)
  }
  if intIP != 0 {
    protocolVars.IPAddress = net.IPv4(data[p-4], data[p-3], data[p-2], data[p-1])
  }

	goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
//line parser.go:1011
		goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		goto st93
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		goto tr97
tr97:
//line parser.rl:89

  protocolVars.Key = data[p-4:p]

	goto st94
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
//line parser.go:1042
		goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		goto st97
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		goto tr101
tr101:
//line parser.rl:92

  buf = bytes.NewReader(data[p-4:])
  err = binary.Read(buf, binary.BigEndian, &protocolVars.RequestedResults)
  if err != nil {
    panic(err)
  }

	goto st98
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
//line parser.go:1077
		goto st138
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		goto tr140
tr140:
//line parser.rl:99

  err = binary.Read(buf, binary.BigEndian, &protocolVars.Port)
  if err != nil {
    panic(err)
  }

	goto st139
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
//line parser.go:1099
		goto st139
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		goto tr103
tr103:
//line parser.rl:21

  protocolVars.Action = ActionScrape;

	goto st100
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
//line parser.go:1118
		goto st101
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		goto st102
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		goto st103
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		goto tr107
tr107:
//line parser.rl:36

  protocolVars.TransactionID = data[p-4:p]

	goto st104
tr142:
//line parser.rl:25

  if protocolVars.InfoHashes == nil {
    protocolVars.InfoHashes = make([][]byte, 0, 74)
  }
  protocolVars.InfoHashes = append(protocolVars.InfoHashes, data[p-20:p])

	goto st104
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
//line parser.go:1158
		goto st105
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		goto st106
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		goto st107
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		goto st109
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		goto st110
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		goto st112
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		goto st113
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		goto st114
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		goto st115
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		goto st116
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		goto st117
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		goto st119
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		goto st120
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		goto st121
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		goto st122
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		goto st140
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		goto tr142
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		if data[p] == 23 {
			goto st124
		}
		goto st5
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		if data[p] == 39 {
			goto st125
		}
		goto st6
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		if data[p] == 16 {
			goto st126
		}
		goto st7
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		if data[p] == 25 {
			goto st127
		}
		goto st8
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		if data[p] == 128 {
			goto st128
		}
		goto st9
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		if data[p] == 0 {
			goto tr132
		}
		goto st0
tr132:
//line parser.rl:39

  protocolVars.ConnectionID = data[p-8:p]

	goto st129
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
//line parser.go:1339
		if data[p] == 0 {
			goto st130
		}
		goto st0
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		if data[p] == 0 {
			goto st131
		}
		goto st0
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 0:
			goto st132
		case 1:
			goto st13
		case 2:
			goto st99
		}
		goto st0
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		goto tr136
tr136:
//line parser.rl:15

  protocolVars.Action = ActionConnectionRequest

	goto st133
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
//line parser.go:1384
		goto st134
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		goto st135
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		goto st141
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		goto st0
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
		goto st4
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		goto st136
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof138: cs = 138; goto _test_eof
	_test_eof139: cs = 139; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof100: cs = 100; goto _test_eof
	_test_eof101: cs = 101; goto _test_eof
	_test_eof102: cs = 102; goto _test_eof
	_test_eof103: cs = 103; goto _test_eof
	_test_eof104: cs = 104; goto _test_eof
	_test_eof105: cs = 105; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof121: cs = 121; goto _test_eof
	_test_eof122: cs = 122; goto _test_eof
	_test_eof140: cs = 140; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
	_test_eof126: cs = 126; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
	_test_eof131: cs = 131; goto _test_eof
	_test_eof132: cs = 132; goto _test_eof
	_test_eof133: cs = 133; goto _test_eof
	_test_eof134: cs = 134; goto _test_eof
	_test_eof135: cs = 135; goto _test_eof
	_test_eof141: cs = 141; goto _test_eof
	_test_eof136: cs = 136; goto _test_eof
	_test_eof137: cs = 137; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 139:
//line parser.rl:32

  valid = true

		case 140:
//line parser.rl:25

  if protocolVars.InfoHashes == nil {
    protocolVars.InfoHashes = make([][]byte, 0, 74)
  }
  protocolVars.InfoHashes = append(protocolVars.InfoHashes, data[p-20:p])

//line parser.rl:32

  valid = true

		case 141:
//line parser.rl:36

  protocolVars.TransactionID = data[p-4:p]

//line parser.rl:32

  valid = true

		case 138:
//line parser.rl:99

  err = binary.Read(buf, binary.BigEndian, &protocolVars.Port)
  if err != nil {
    panic(err)
  }

//line parser.rl:32

  valid = true

//line parser.go:1599
		}
	}

	_out: {}
	}

//line parser.rl:131

  fmt.Printf("%+v\n", protocolVars)

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
