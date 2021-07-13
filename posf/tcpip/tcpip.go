// TCP IP
// Прослушиваем канал "lo" 127.0.0.1:2999
// - предварительно запустить эмулятор DSP: test/cmd/dsp/main.go
// - и дергаем его из браузера или postman'a (127.0.0.1:2999)

// для запуска программы необходимо выполнить:
// go build tcpip.go
// sudo ./tcpip

package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

func main() {
	fmt.Println("Слушаем интерфейс lo:127.0.0.1:2999")

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
		//fmt.Println(packet)
		packetData(packet)
		//fmt.Println()
	}
}

func packetData(packet gopacket.Packet) {
	// просмотр IP
	//ipLayer := packet.Layer(layers.LayerTypeIPv4)
	//if ipLayer != nil {
	//	fmt.Println("IP:")
	//	ip, _ := ipLayer.(*layers.IPv4)
	//
	//	// IP layer variables:
	//	// Version (Either 4 or 6)
	//	// IHL (IP Header Length in 32-bit words)
	//	// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?), Checksum, SrcIP, DstIP
	//
	//	fmt.Println("protocol: ", ip.Protocol)
	//	fmt.Printf(" from IP  %s to %s\n", ip.SrcIP, ip.DstIP)
	//	fmt.Println()
	//}

	// просмотр TCP
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {

		tcp, _ := tcpLayer.(*layers.TCP)

		// TCP layer variables:
		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS

		if tcp.PSH {
			fmt.Println("TCP:")
			fmt.Printf("from port  %d to %d\n", tcp.SrcPort, tcp.DstPort)
			fmt.Println("tcp_window_size: ", tcp.Window)
			tcpOptions := tcp.Options
			for _, option := range tcpOptions {
				fmt.Println("TCP option type: ", option.OptionType)
				fmt.Println("TCP options data: ", string(option.OptionData))
			}
			fmt.Println()

		} else {
			return
		}
	}

	// Слой приложений содержит полезную нагрузку
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		fmt.Println("Application layer/Payload found:")
		fmt.Printf("%s\n", applicationLayer.Payload())
	}

	// контроль ошибок
	if err := packet.ErrorLayer(); err != nil {
		fmt.Println("Error decoding some part of the packet:", err)
	}
}
