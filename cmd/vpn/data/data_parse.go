package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type serversNordVPNData struct {
	ID             int          `json:"id"`
	IPAddress      string       `json:"ip_address"`
	SearchKeywords []string     `json:"search_keywords"`
	Categories     []categories `json:"categories"`
	Name           string       `json:"name"`
	Domain         string       `json:"domain"`
	Price          int          `json:"price"`
	Location       location     `json:"location"`
	Flag           string       `json:"flag"`
	Country        string       `json:"country"`
	Load           int          `json:"load"`
	Features       features     `json:"features"`
}

type location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type categories struct {
	Name string `json:"name"`
}

type features struct {
	Ikev2               bool `json:"ikev2"`                 // 1
	OpenvpnUDP          bool `json:"openvpn_udp"`           // 2
	OpenvpnTCP          bool `json:"openvpn_tcp"`           // 3
	Socks               bool `json:"socks"`                 // 4
	Proxy               bool `json:"proxy"`                 // 5
	Pptp                bool `json:"pptp"`                  // 6
	L2tp                bool `json:"l2tp"`                  // 7
	OpenvpnXorUDP       bool `json:"openvpn_xor_udp"`       // 8
	OpenvpnXorTCP       bool `json:"openvpn_xor_tcp"`       // 9
	ProxyCybersec       bool `json:"proxy_cybersec"`        // 10
	ProxySSL            bool `json:"proxy_ssl"`             // 11
	ProxySSLCybersec    bool `json:"proxy_ssl_cybersec"`    // 12
	Ikev2V6             bool `json:"ikev2v6"`               // 13
	OpenvpnUDPv6        bool `json:"openvpn_udp_v6"`        // 14
	OpenvpnTCPv6        bool `json:"openvpn_tcp_v6"`        // 15
	WireguardUDP        bool `json:"wireguard_udp"`         // 16
	OpenvpnUdpTlsCrypt  bool `json:"openvpn_udp_tls_crypt"` // 17
	OpenvpnTcpTlsCrypt  bool `json:"openvpn_tcp_tls_crypt"` // 18
	OpenvpnDedicatedUdp bool `json:"openvpn_dedicated_udp"` // 19
	OpenvpnDedicatedTcp bool `json:"openvpn_dedicated_tcp"` // 20
	Skylark             bool `json:"skylark"`               // 21
	MeshRelay           bool `json:"mesh_relay"`            // 22
}

type counter struct {
	Ikev2               int `json:"ikev2"`                 // 1
	OpenvpnUDP          int `json:"openvpn_udp"`           // 2
	OpenvpnTCP          int `json:"openvpn_tcp"`           // 3
	Socks               int `json:"socks"`                 // 4
	Proxy               int `json:"proxy"`                 // 5
	Pptp                int `json:"pptp"`                  // 6
	L2tp                int `json:"l2tp"`                  // 7
	OpenvpnXorUDP       int `json:"openvpn_xor_udp"`       // 8
	OpenvpnXorTCP       int `json:"openvpn_xor_tcp"`       // 9
	ProxyCybersec       int `json:"proxy_cybersec"`        // 10
	ProxySSL            int `json:"proxy_ssl"`             // 11
	ProxySSLCybersec    int `json:"proxy_ssl_cybersec"`    // 12
	Ikev2V6             int `json:"ikev2v6"`               // 13
	OpenvpnUDPv6        int `json:"openvpn_udp_v6"`        // 14
	OpenvpnTCPv6        int `json:"openvpn_tcp_v6"`        // 15
	WireguardUDP        int `json:"wireguard_udp"`         // 16
	OpenvpnUdpTlsCrypt  int `json:"openvpn_udp_tls_crypt"` // 17
	OpenvpnTcpTlsCrypt  int `json:"openvpn_tcp_tls_crypt"` // 18
	OpenvpnDedicatedUdp int `json:"openvpn_dedicated_udp"` // 19
	OpenvpnDedicatedTcp int `json:"openvpn_dedicated_tcp"` // 20
	Skylark             int `json:"skylark"`               // 21
	MeshRelay           int `json:"mesh_relay"`            // 22
}

type counterCountry struct {
	All                 map[string]int // 0
	Ikev2               map[string]int // 1
	OpenvpnUDP          map[string]int // 2
	OpenvpnTCP          map[string]int // 3
	Socks               map[string]int // 4
	Proxy               map[string]int // 5
	Pptp                map[string]int // 6
	L2tp                map[string]int // 7
	OpenvpnXorUDP       map[string]int // 8
	OpenvpnXorTCP       map[string]int // 9
	ProxyCybersec       map[string]int // 10
	ProxySSL            map[string]int // 11
	ProxySSLCybersec    map[string]int // 12
	Ikev2V6             map[string]int // 13
	OpenvpnUDPv6        map[string]int // 14
	OpenvpnTCPv6        map[string]int // 15
	WireguardUDP        map[string]int // 16
	OpenvpnUdpTlsCrypt  map[string]int // 17
	OpenvpnTcpTlsCrypt  map[string]int // 18
	OpenvpnDedicatedUdp map[string]int // 19
	OpenvpnDedicatedTcp map[string]int // 20
	Skylark             map[string]int // 21
	MeshRelay           map[string]int // 22
}

func newCounterCountry() counterCountry {
	return counterCountry{
		All:                 make(map[string]int),
		Ikev2:               make(map[string]int),
		OpenvpnUDP:          make(map[string]int),
		OpenvpnTCP:          make(map[string]int),
		Socks:               make(map[string]int),
		Proxy:               make(map[string]int),
		Pptp:                make(map[string]int),
		L2tp:                make(map[string]int),
		OpenvpnXorUDP:       make(map[string]int),
		OpenvpnXorTCP:       make(map[string]int),
		ProxyCybersec:       make(map[string]int),
		ProxySSL:            make(map[string]int),
		ProxySSLCybersec:    make(map[string]int),
		Ikev2V6:             make(map[string]int),
		OpenvpnUDPv6:        make(map[string]int),
		OpenvpnTCPv6:        make(map[string]int),
		WireguardUDP:        make(map[string]int),
		OpenvpnUdpTlsCrypt:  make(map[string]int),
		OpenvpnTcpTlsCrypt:  make(map[string]int),
		OpenvpnDedicatedUdp: make(map[string]int),
		OpenvpnDedicatedTcp: make(map[string]int),
		Skylark:             make(map[string]int),
		MeshRelay:           make(map[string]int),
	}
}

func main() {
	fileName := "cmd/vpn/data/servers_020222.json"

	data := getData(fileName)

	log.Println("records:", len(data))

	count, c := countFeatures(data)
	log.Println("count:", count)
	log.Println("country:", len(c.All), len(c.Ikev2), len(c.OpenvpnUDP), len(c.OpenvpnTCP), len(c.Socks),
		len(c.OpenvpnXorUDP), len(c.OpenvpnXorTCP), len(c.ProxySSL), len(c.ProxySSLCybersec),
		len(c.WireguardUDP), len(c.OpenvpnDedicatedUdp), len(c.OpenvpnDedicatedTcp), len(c.MeshRelay))

	fileNameSaved := "cmd/vpn/data/servers_020222.csv"
	saveData(fileNameSaved, data)

	log.Println("Done...")
}

func saveData(fileName string, data []serversNordVPNData) {
	var dataFile, s string

	s = fmt.Sprintf("%s,%s,%s,", "IPAddress", "Flag", "Country")
	s += fmt.Sprintf("%s,%s,%s,", "Ikev2", "OpenvpnUDP", "OpenvpnTCP")
	s += fmt.Sprintf("%s,%s,%s,", "Socks", "OpenvpnXorUDP", "OpenvpnXorTCP")
	s += fmt.Sprintf("%s,%s,%s,", "ProxySSL", "ProxySSLCybersec", "WireguardUDP")
	s += fmt.Sprintf("%s,%s,%s\n", "OpenvpnDedicatedUdp", "OpenvpnDedicatedTcp", "MeshRelay")
	dataFile = s

	s = ""
	for _, d := range data {
		s = fmt.Sprintf("%s,%s,%s,", d.IPAddress, d.Flag, d.Country)
		s += fmt.Sprintf("%t,%t,%t,", d.Features.Ikev2, d.Features.OpenvpnUDP, d.Features.OpenvpnTCP)
		s += fmt.Sprintf("%t,%t,%t,", d.Features.Socks, d.Features.OpenvpnXorUDP, d.Features.OpenvpnXorTCP)
		s += fmt.Sprintf("%t,%t,%t,", d.Features.ProxySSL, d.Features.ProxySSLCybersec, d.Features.WireguardUDP)
		s += fmt.Sprintf("%t,%t,%t\n", d.Features.OpenvpnDedicatedUdp, d.Features.OpenvpnDedicatedTcp, d.Features.MeshRelay)

		dataFile += s

	}

	saveDataFile(fileName, []byte(dataFile))

}

func saveDataFile(fileName string, data []byte) {
	df, errCreateFile := os.Create(fileName)
	if errCreateFile != nil {
		fmt.Errorf("error create file, %w", errCreateFile)
		return
	}
	defer df.Close()

	_, errWrite := df.Write(data)
	if errWrite != nil {
		fmt.Errorf("error write data %w", errWrite)
		return
	}
	log.Printf("file saved: %s\n", fileName)
}

func countFeatures(data []serversNordVPNData) (counter, counterCountry) {
	var res counter
	c := newCounterCountry()

	for _, d := range data {
		c.All[d.Flag]++

		if d.Features.Ikev2 { // 1
			res.Ikev2++
			c.Ikev2[d.Flag]++
		}
		if d.Features.OpenvpnUDP { // 2
			res.OpenvpnUDP++
			c.OpenvpnUDP[d.Flag]++
		}
		if d.Features.OpenvpnTCP { // 3
			res.OpenvpnTCP++
			c.OpenvpnTCP[d.Flag]++
		}
		if d.Features.Socks { // 4
			res.Socks++
			c.Socks[d.Flag]++
		}
		if d.Features.Proxy { // 5
			res.Proxy++
			c.Proxy[d.Flag]++
		}
		if d.Features.Pptp { // 6
			res.Pptp++
			c.Pptp[d.Flag]++
		}
		if d.Features.L2tp { // 7
			res.L2tp++
			c.L2tp[d.Flag]++
		}
		if d.Features.OpenvpnXorUDP { // 8
			res.OpenvpnXorUDP++
			c.OpenvpnXorUDP[d.Flag]++
		}
		if d.Features.OpenvpnXorTCP { // 9
			res.OpenvpnXorTCP++
			c.OpenvpnXorTCP[d.Flag]++
		}
		if d.Features.ProxyCybersec { // 10
			res.ProxyCybersec++
			c.ProxyCybersec[d.Flag]++
		}
		if d.Features.ProxySSL { // 11
			res.ProxySSL++
			c.ProxySSL[d.Flag]++
		}
		if d.Features.ProxySSLCybersec { // 12
			res.ProxySSLCybersec++
			c.ProxySSLCybersec[d.Flag]++
		}
		if d.Features.Ikev2V6 { // 13
			res.Ikev2V6++
			c.Ikev2V6[d.Flag]++
		}
		if d.Features.OpenvpnUDPv6 { // 14
			res.OpenvpnUDPv6++
			c.OpenvpnUDPv6[d.Flag]++
		}
		if d.Features.OpenvpnTCPv6 { // 15
			res.OpenvpnTCPv6++
			c.OpenvpnTCPv6[d.Flag]++
		}
		if d.Features.WireguardUDP { // 16
			res.WireguardUDP++
			c.WireguardUDP[d.Flag]++
		}
		if d.Features.OpenvpnUdpTlsCrypt { // 17
			res.OpenvpnUdpTlsCrypt++
			c.OpenvpnUdpTlsCrypt[d.Flag]++
		}
		if d.Features.OpenvpnTcpTlsCrypt { // 18
			res.OpenvpnTcpTlsCrypt++
			c.OpenvpnTcpTlsCrypt[d.Flag]++
		}
		if d.Features.OpenvpnDedicatedUdp { // 19
			res.OpenvpnDedicatedUdp++
			c.OpenvpnDedicatedUdp[d.Flag]++
		}
		if d.Features.OpenvpnDedicatedTcp { // 20
			res.OpenvpnDedicatedTcp++
			c.OpenvpnDedicatedTcp[d.Flag]++
		}
		if d.Features.Skylark { // 21
			res.Skylark++
			c.Skylark[d.Flag]++
		}
		if d.Features.MeshRelay { // 22
			res.MeshRelay++
			c.MeshRelay[d.Flag]++
		}
	}
	return res, c
}

func getData(fileName string) []serversNordVPNData {
	file, errOpen := ioutil.ReadFile(fileName)
	if errOpen != nil {
		fmt.Println("error open file stamp data", errOpen.Error())
		return nil
	}

	data := make([]serversNordVPNData, 0)

	errJSON := json.Unmarshal(file, &data)
	if errJSON != nil {
		fmt.Println("error unmarshal data", errJSON.Error())
		return nil
	}

	return data
}
