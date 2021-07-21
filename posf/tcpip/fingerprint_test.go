package main

//
//func Test_fingerprint(t *testing.T) {
//	data := []byte(`
//[{
//    "ts": 1615394076,
//    "ip_df": 1,
//    "ip_mf": 0,
//    "tcp_window_size": 64240,
//    "tcp_flags": 2,
//    "tcp_ack": 0,
//    "tcp_header_length": 128,
//    "tcp_urp": 0,
//    "tcp_options": "M1460,N,W8,N,N,S,",
//    "tcp_timestamp_echo_reply": 0,
//    "tcp_mss": 1460,
//    "navigatorUserAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4427.0 Safari/537.36",
//    "platform": {
//      "type": "desktop"
//    },
//    "deviceMemory": 2,
//    "hardwareConcurrency": 4,
//    "os": {
//      "name": "Windows",
//      "version": "NT 10.0",
//      "versionName": "10"
//    }
//  },
//  {
//    "ts": 1615406175,
//    "ip_df": 1,
//    "ip_mf": 0,
//    "tcp_window_size": 64800,
//    "tcp_flags": 2,
//    "tcp_ack": 0,
//    "tcp_header_length": 160,
//    "tcp_urp": 0,
//    "tcp_options": "M1350,S,T,N,W7,",
//    "tcp_timestamp_echo_reply": 0,
//    "tcp_mss": 1350,
//    "navigatorUserAgent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
//    "platform": {
//      "type": "desktop"
//    },
//    "deviceMemory": 8,
//    "hardwareConcurrency": 8,
//    "os": {
//      "name": "Linux"
//    }
//  },
//  {
//    "ts": 1615408247,
//    "ip_df": 1,
//    "ip_mf": 0,
//    "tcp_window_size": 64240,
//    "tcp_flags": 2,
//    "tcp_ack": 0,
//    "tcp_header_length": 128,
//    "tcp_urp": 0,
//    "tcp_options": "M1460,N,W8,N,N,S,",
//    "tcp_timestamp_echo_reply": 0,
//    "tcp_mss": 1460,
//    "navigatorUserAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36",
//    "platform": {
//      "type": "desktop"
//    },
//    "deviceMemory": 8,
//    "hardwareConcurrency": 8,
//    "os": {
//      "name": "Windows",
//      "version": "NT 10.0",
//      "versionName": "10"
//    }
//  }
//]`)
//
//	user := Combined{
//		IPDf:            1,
//		IPMf:            0,
//		TCPWindowSize:   64240,
//		TCPFlags:        2,
//		TCPHeaderLength: 128,
//		TCPOptions:      "M1460,N,W8,N,N,S,",
//		TCPMss:          1460,
//	}
//
//	uabase := dataUnmarshal(data)
//
//	_ = user
//
//	scoresOS := fingerprint(uabase, user)
//
//	//fmt.Println(scoresOS)
//
//	resultOS(scoresOS, 3)
//
//}
