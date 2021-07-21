// TCP IP
// Прослушиваем канал "lo" 127.0.0.1:2999
// - предварительно запустить эмулятор DSP: test/cmd/dsp/main.go
// - и дергаем его из браузера или postman'a (127.0.0.1:2999)

// для запуска программы необходимо выполнить:
// go build tcpip.go
// sudo ./tcpip

package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Combined struct {
	Ts                    int    `json:"ts"` // Stamp
	IPTTL                 int    `json:"ip_ttl"`
	IPDf                  int    `json:"ip_df"`           // score += 10
	IPMf                  int    `json:"ip_mf"`           // score += 10
	TCPWindowSize         int    `json:"tcp_window_size"` // score += 15
	TCPFlags              int    `json:"tcp_flags"`       // score += 10
	TCPAck                int    `json:"tcp_ack"`
	TCPSeq                int    `json:"tcp_seq"`
	TCPHeaderLength       int    `json:"tcp_header_length"` // score += 10
	TCPUrp                int    `json:"tcp_urp"`
	TCPOptions            string `json:"tcp_options"` // score += 30 || score += 20
	TCPWindowScaling      int    `json:"tcp_window_scaling"`
	TCPTimestamp          int    `json:"tcp_timestamp"`
	TCPTimestampEchoReply int    `json:"tcp_timestamp_echo_reply"`
	TCPMss                int    `json:"tcp_mss"` // score += 15
	NavigatorUserAgent    string `json:"navigatorUserAgent"`
	Platform              struct {
		Type   string `json:"type"`
		Vendor string `json:"vendor"`
		Model  string `json:"model"`
	} `json:"platform"`
	DeviceMemory        float32  `json:"devicememory"`
	HardwareConcurrency int      `json:"hardwareconcurrency"`
	Os                  struct { // data OS
		Name        string `json:"name"`
		Version     string `json:"version"`
		VersionName string `json:"versionname"`
	} `json:"os"`
	HTTPPayload string `json:"http_payload"`
	HTTPuin     string `json:"http_uin"`
}

func main() {
	fmt.Println("Слушаем интерфейс lo:127.0.0.1:2999")

	fileName := "./database/combined.json"
	data := readFile(fileName)
	fmt.Print("dataBase = ", len(data), " bytes, ")

	uabase := dataUnmarshal(data)
	fmt.Println(len(uabase), "records")

	var (
		deviceName  string = "lo"
		snapshotLen int32  = 1024 //65536
		promiscuous bool   = false
		err         error
		timeout     time.Duration = 1 * time.Second
		handle      *pcap.Handle
	)
	// Open device
	handle, err = pcap.OpenLive(deviceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter
	var filter string = "tcp and port 2999"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("filter: only capturing TCP port 2999 packets")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
		//packetData(packet)
		//fmt.Println()

		u, ok := getTCPIP(packet)
		if ok {
			fmt.Println()
			fmt.Println(u)
			scores := fingerprint(uabase, u)
			resultOS(scores, 3)
		}
	}
}

func getTCPIP(packet gopacket.Packet) (Combined, bool) {
	u := Combined{}
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if ipLayer == nil || tcpLayer == nil {
		return u, false
	}

	ip, _ := ipLayer.(*layers.IPv4)
	tcp, _ := tcpLayer.(*layers.TCP)

	if tcp.SrcPort.String() == "2999(remoteware-un)" {
		//fmt.Println("tcp.SrcPort:", tcp.SrcPort.String())
		return u, false
	}

	if !tcp.PSH { // проверка флага PSH
		return u, false
	}
	//if !tcp.SYN { // проверка флага SYN
	//	return u, false
	//}
	//if !tcp.ACK || !tcp.SYN { // проверка флага ACK и SYN
	//	return u, false
	//}

	if tcp.FIN { // проверка флага FIN
		fmt.Println("FIN")
		return u, false
	}

	if ip.Flags.String() == "DF" {
		u.IPDf = 1
		u.IPMf = 0
	}
	if ip.Flags.String() == "MF" {
		u.IPDf = 0
		u.IPMf = 1
	}

	u.TCPWindowSize = int(tcp.Window)
	u.TCPFlags = int(ip.FragOffset) // ?!
	u.TCPHeaderLength = len(tcp.Contents)
	u.TCPOptions, u.TCPMss, u.TCPTimestamp, u.TCPTimestampEchoReply, u.TCPWindowScaling = decodeTCPOptions(tcp.Options)

	// HTTP
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		u.HTTPPayload = string(applicationLayer.Payload())
	}

	// контроль ошибок
	if err := packet.ErrorLayer(); err != nil {
		fmt.Println("Error decoding some part of the packet:", err)
	}

	//fmt.Println(u)

	return u, true
}

//
//func packetData(packet gopacket.Packet) {
//	//просмотр IP
//	ipLayer := packet.Layer(layers.LayerTypeIPv4)
//	if ipLayer != nil {
//		fmt.Println("IP:")
//		ip, _ := ipLayer.(*layers.IPv4)
//
//		// IP layer variables:
//		// Version (Either 4 or 6)
//		// IHL (IP Header Length in 32-bit words)
//		// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?), Checksum, SrcIP, DstIP
//
//		fmt.Println("protocol: ", ip.Protocol)
//		fmt.Printf(" from IP  %s to %s\n", ip.SrcIP, ip.DstIP)
//		fmt.Println()
//	}
//
//	// просмотр TCP
//	tcpLayer := packet.Layer(layers.LayerTypeTCP)
//	if tcpLayer != nil {
//
//		tcp, _ := tcpLayer.(*layers.TCP)
//
//		// TCP layer variables:
//		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
//		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
//
//		if tcp.PSH {
//			fmt.Println("TCP:")
//			fmt.Printf("from port  %d to %d\n", tcp.SrcPort, tcp.DstPort)
//			fmt.Println("tcp_window_size: ", tcp.Window)
//			tcpOptions := tcp.Options
//			for _, option := range tcpOptions {
//				fmt.Println("TCP option type: ", option.OptionType)
//				fmt.Println("TCP options data: ", string(option.OptionData))
//			}
//			fmt.Println()
//
//		} else {
//			return
//		}
//	}
//
//	// Слой приложений содержит полезную нагрузку
//	applicationLayer := packet.ApplicationLayer()
//	if applicationLayer != nil {
//		fmt.Println("Application layer/Payload found:")
//		fmt.Printf("%s\n", applicationLayer.Payload())
//	}
//
//	// контроль ошибок
//	if err := packet.ErrorLayer(); err != nil {
//		fmt.Println("Error decoding some part of the packet:", err)
//	}
//}

// TCP Options (opt_type)
const (
	TCP_OPT_EOL        = 0   // end of option list
	TCP_OPT_NOP        = 1   // no operation
	TCP_OPT_MSS        = 2   // maximum segment size
	TCP_OPT_WSCALE     = 3   // window scale factor, RFC 1072
	TCP_OPT_SACKOK     = 4   // SACK permitted, RFC 2018
	TCP_OPT_SACK       = 5   // SACK, RFC 2018
	TCP_OPT_ECHO       = 6   // echo (obsolete), RFC 1072
	TCP_OPT_ECHOREPLY  = 7   // echo reply (obsolete), RFC 1072
	TCP_OPT_TIMESTAMP  = 8   // timestamps, RFC 1323
	TCP_OPT_POCONN     = 9   // partial order conn, RFC 1693
	TCP_OPT_POSVC      = 10  // partial order service, RFC 1693
	TCP_OPT_CC         = 11  // connection count, RFC 1644
	TCP_OPT_CCNEW      = 12  // CC.NEW, RFC 1644
	TCP_OPT_CCECHO     = 13  // CC.ECHO, RFC 1644
	TCP_OPT_ALTSUM     = 14  // alt checksum request, RFC 1146
	TCP_OPT_ALTSUMDATA = 15  // alt checksum data, RFC 1146
	TCP_OPT_SKEETER    = 16  // Skeeter
	TCP_OPT_BUBBA      = 17  // Bubba
	TCP_OPT_TRAILSUM   = 18  // trailer checksum
	TCP_OPT_MD5        = 19  // MD5 signature, RFC 2385
	TCP_OPT_SCPS       = 20  // SCPS capabilities
	TCP_OPT_SNACK      = 21  // selective negative acks
	TCP_OPT_REC        = 22  // record boundaries
	TCP_OPT_CORRUPT    = 23  // corruption experienced
	TCP_OPT_SNAP       = 24  // SNAP
	TCP_OPT_TCPCOMP    = 26  // TCP compression filter
	TCP_OPT_MAX        = 27  // Quick-Start Response
	TCP_OPT_USRTO      = 28  // User Timeout Option (also, other known unauthorized use) [***][1]	[RFC5482]
	TCP_OPT_AUTH       = 29  // TCP Authentication Option (TCP-AO)	[RFC5925]
	TCP_OPT_MULTIPATH  = 30  // Multipath TCP (MPTCP)
	TCP_OPT_FASTOPEN   = 34  // TCP Fast Open Cookie	[RFC7413]
	TCP_OPY_ENCNEG     = 69  // Encryption Negotiation (TCP-ENO)	[RFC8547]
	TCP_OPT_EXP1       = 253 // RFC3692-style Experiment 1 (also improperly used for shipping products)
	TCP_OPT_EXP2       = 254 // RFC3692-style Experiment 2 (also improperly used for shipping products)
)

func decodeTCPOptions(options []layers.TCPOption) (option string, mss, tcpTimeStamp, tcpTimeStampEchoReply, windowScaling int) {
	var err error
	for _, op := range options {
		switch op.OptionType {
		case TCP_OPT_EOL: // End of options list
			option += "E,"
		case TCP_OPT_NOP: // No operation
			option += "N,"
		case TCP_OPT_MSS: // Maximum segment size
			mssStr := ""
			if len(op.OptionData) > 0 {
				mssStr = string(op.OptionData[0])
				mss, err = strconv.Atoi(mssStr)
				if err != nil {
					//fmt.Println("error: ", err.Error(), mssStr)
					mssStr = ""
					mss = 0
				}
			}
			option += "M" + mssStr + ","

		case TCP_OPT_WSCALE: // Window scaling
			windowScalingStr := ""
			if len(op.OptionData) > 0 {
				windowScalingStr = string(op.OptionData[0])
				windowScaling, err = strconv.Atoi(windowScalingStr)
				if err != nil {
					//fmt.Println("error: ", err.Error(), windowScalingStr)
					windowScalingStr = ""
					windowScaling = 0
				}
			}
			option += "W" + windowScalingStr + ","

		case TCP_OPT_SACKOK: // Selective Acknowledgement permitted
			option += "S,"
		case TCP_OPT_SACK: // Selective ACKnowledgement (SACK)
			option += "K,"
		case TCP_OPT_ECHO:
			option += "J,"
		case TCP_OPT_ECHOREPLY:
			option += "F,"
		case TCP_OPT_TIMESTAMP:
			option += "T,"
			if len(op.OptionData) > 8 {
				tcpTimeStamp, err = strconv.Atoi(string(op.OptionData[0:4]))
				if err != nil {
					fmt.Println("error: ", err.Error(), string(op.OptionData[0:4]))
					tcpTimeStamp = 0
				}
				tcpTimeStampEchoReply, err = strconv.Atoi(string(op.OptionData[4:8]))
				if err != nil {
					fmt.Println("error: ", err.Error(), string(op.OptionData[4:8]))
					tcpTimeStampEchoReply = 0
				}
			}
		case TCP_OPT_POCONN:
			option += "P,"
		case TCP_OPT_POSVC:
			option += "R,"
		default:
			option += "U" + strconv.Itoa(int(op.OptionType)) + ","
		}
	}
	return
}

type ScoreOS struct {
	Stamp               int // Combined.Ts - индификатор для поиска в базе
	Score               int // Рейтинг от 0 до 100
	PlatformName        string
	PlatformVersion     string
	PlatformVersionName string
}

// расчет рейтинга
func fingerprint(uabase []Combined, user Combined) []ScoreOS {
	scores := make([]ScoreOS, 0, len(uabase))
	for _, ua := range uabase {
		score := 0
		if ua.Os.Name == "" {
			continue
		}

		if ua.IPDf == user.IPDf {
			score += 10
		}
		if ua.IPMf == user.IPMf {
			score += 10
		}
		if ua.TCPWindowSize == user.TCPWindowSize {
			score += 15
		}
		if ua.TCPFlags == user.TCPFlags {
			score += 10
		}
		if ua.TCPHeaderLength == user.TCPHeaderLength {
			score += 10
		}
		if ua.TCPMss == user.TCPMss {
			score += 15
		}
		if ua.TCPOptions == user.TCPOptions {
			score += 30
		}

		// проверяем порядок параметров TCP (это слабее, чем равенство параметров TCP)
		if ua.TCPOptions != user.TCPOptions && ua.TCPOptions != "" && user.TCPOptions != "" {
			var orderUA, orderUser string
			for _, e := range strings.Split(ua.TCPOptions, ",") {
				if e != "" {
					orderUA += string(e[0])
				}
			}
			for _, e := range strings.Split(user.TCPOptions, ",") {
				if e != "" {
					orderUser += string(e[0])
				}
			}
			if orderUA == orderUser {
				score += 20
			}
		}

		sc := ScoreOS{
			Stamp:               ua.Ts,
			Score:               score,
			PlatformName:        ua.Os.Name,
			PlatformVersion:     ua.Os.Version,
			PlatformVersionName: ua.Os.VersionName,
		}
		scores = append(scores, sc)
	}
	return scores
}

// вывод n элементов по ТОП рейтингу - Инфо: рейтинг ОС, Имя ОС и версия
func resultOS(scores []ScoreOS, n int) {
	sort.SliceStable(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score // сортировка по убыванию рейтинга
	})

	i := 0
	for _, sc := range scores {
		fmt.Printf("score %3d  OS: %s %s  %s \n", sc.Score, sc.PlatformName, sc.PlatformVersionName, sc.PlatformVersion)
		i++
		if i == n {
			break
		}
	}
}

func readFile(fileName string) []byte {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	if len(data) == 0 {
		fmt.Println("Exit (data = 0)")
		return nil
	}
	return data
}

func dataUnmarshal(data []byte) []Combined {
	d := []Combined{}
	err := json.Unmarshal(data, &d)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	return d
}
