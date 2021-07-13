// https://russianblogs.com/article/9204277465/

// для запуска программы необходимо выполнить:
// go build ip.go
// sudo ./ip

package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	//findAllDevices()
	//openDevice()
	//pcapFile()
	//readPcapFile()

	//filterPacket()

	decodePacket()
}

// Получить всю информацию о сетевом устройстве
func findAllDevices() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Devices found:")

	for _, device := range devices {
		fmt.Println("\nName: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println("Devices  flags: ", device.Flags)
		for _, address := range device.Addresses {
			fmt.Println("- IP address : ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
}

// Откройте устройство для захвата в режиме реального времени
func openDevice() {
	var (
		//device            = "eth0"
		device            = "enp7s0"
		snapshotLen int32 = 1024
		promiscuous       = false
		err         error
		timeout     = 30 * time.Second
		handle      *pcap.Handle
	)

	// Open device
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
	}
}

// Результат захвата сохраняется как файл формата pcap
func pcapFile() {
	var (
		deviceName  string = "enp7s0"
		snapshotLen int32  = 1024
		promiscuous bool   = false
		err         error
		timeout     time.Duration = 30 * time.Second
		handle      *pcap.Handle
		packetCount int = 0
	)

	// Open output pcap file and write header
	f, _ := os.Create("file.pcap")
	w := pcapgo.NewWriter(f)
	_ = w.WriteFileHeader(uint32(snapshotLen), layers.LinkTypeEthernet)
	defer f.Close()

	// Open the device for capturing
	handle, err = pcap.OpenLive(deviceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		fmt.Printf("Error opening	device % s: %v ", deviceName, err)
		os.Exit(1)
	}
	defer handle.Close()

	// Start processing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
		_ = w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())

		packetCount++
		// Only capture 100 and then stop
		if packetCount > 100 {
			break
		}
	}
}

// Прочитайте файл формата pcap для просмотра и анализа сетевого пакета
func readPcapFile() {
	// Open file instead of device
	handle, err := pcap.OpenOffline("file.pcap")
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet.Data())
	}
}

// Установите фильтр - Только захватить TCP-порт 80 данных
func filterPacket() {
	var (
		deviceName  string = "enp7s0"
		snapshotLen int32  = 1024
		promiscuous bool   = false
		err         error
		timeout     time.Duration = 30 * time.Second
		handle      *pcap.Handle
	)

	// Open device
	handle, err = pcap.OpenLive(deviceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter
	var filter string = "tcp and port 80"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Only capturing TCP port 80 packets")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}

// Расшифруйте захваченные данные
func decodePacket() {
	var (
		deviceName  string = "lo"
		snapshotLen int32  = 1024
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
	fmt.Println("Only capturing TCP port 2999 packets")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		printPacketInfo(packet)
	}
}

func printPacketInfo(packet gopacket.Packet) {
	// посмотрим, является ли пакет пакетом ethernet
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		fmt.Println("Ethernet layer detected.")
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
		fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
		// Ethernet type is typically IPv4 but could be ARP or other
		fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
		fmt.Println()
	}

	// посмотрим, является ли пакет IP-адресом (хотя тип эфира сообщил нам об этом)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IPv4 layer detected.")
		ip, _ := ipLayer.(*layers.IPv4)

		// IP layer variables:
		// Version (Either 4 or 6)
		// IHL (IP Header Length in 32-bit words)
		// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
		// Checksum, SrcIP, DstIP
		fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		fmt.Println("Protocol: ", ip.Protocol)
		fmt.Println()
	}

	// посмотрим, является ли пакет TCP
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		fmt.Println("TCP layer detected.")
		tcp, _ := tcpLayer.(*layers.TCP)

		// TCP layer variables:
		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
		fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
		fmt.Println("Sequence number: ", tcp.Seq)
		fmt.Println()
	}

	// Повторите все слои, распечатав каждый тип слоя
	fmt.Println("All packet layers:")
	for _, layer := range packet.Layers() {
		fmt.Println("- ", layer.LayerType())
	}

	// При повторном просмотре packet.Layers() выше,
	// если в нем указан слой полезной нагрузки, то это то же самое, что и этот слой приложения.

	// Слой приложений содержит полезную нагрузку
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		fmt.Println("Application layer/Payload found.")
		fmt.Printf("%s\n", applicationLayer.Payload())

		// Поиск строки внутри полезной нагрузки
		if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
			fmt.Println("HTTP found!")
		}
	}

	// Проверьте наличие ошибок
	if err := packet.ErrorLayer(); err != nil {
		fmt.Println("Error decoding some part of the packet:", err)
	}
}
