package main

/*
99.adscompass.ru
2000.99.adscompass.ru
3000.99.adscompass.ru
4000.99.adscompass.ru
*/

import (
	"crypto/tls"
	"log"
)
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	serverNordVPN()
}

func serverNordVPN() {
	log.Println("start...")

	user := "ysETCpBC8JvFzJnt7SjsJxJC"
	password := "qRhXR6k8yW4pcW71c34ReDW3"

	// Socks
	//protocol, ip, port, flag := "socks5", "109.202.99.35", "1080", "NL"
	//protocol, ip, port, flag := "socks5", "165.231.210.171", "1080", "US"
	//protocol, ip, port, flag := "socks5", "196.196.244.3", "1080", "SE"
	//protocol, ip, port, flag := "socks5", "196.196.192.11", "1080", "IE"

	// ProxySSL
	//protocol, ip, port, flag := "https", "172.83.40.219", "89", "CA"
	//protocol, ip, port, flag := "https", "194.99.105.100", "89", "PL"
	//protocol, ip, port, flag := "https", "82.102.20.236", "89", "DK"
	//protocol, ip, port, flag := "https", "82.102.19.137", "89", "BE"
	//protocol, ip, port, flag := "https", "185.206.225.196", "89", "NO"
	//protocol, ip, port, flag := "https", "185.189.114.28", "89", "HU"
	//protocol, ip, port, flag := "https", "185.245.87.59", "89", "US"
	//protocol, ip, port, flag := "https", "185.216.34.100", "89", "AT"
	//protocol, ip, port, flag := "https", "81.92.202.11", "89", "GB"
	//protocol, ip, port, flag := "https", "89.238.186.244", "89", "CZ"
	protocol, ip, port, flag := "https", "82.102.18.252", "89", "FR"

	addrNordVPN := fmt.Sprintf("%s://%s:%s@%s:%s", protocol, user, password, ip, port)
	urlAddrNordVPN, err := url.Parse(addrNordVPN)
	if err != nil {
		log.Println(err)
	}
	log.Println("urlAddrNordVPN =", urlAddrNordVPN)

	transport := &http.Transport{
		Proxy:           http.ProxyURL(urlAddrNordVPN),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	//adding the Transport object to the http Client
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 10,
	}

	//generating the HTTP GET request
	//urlData := "http://eth0.me"
	urlData := "http://4000.99.adscompass.ru/" + flag

	request, err := http.NewRequest("GET", urlData, nil)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("StatusCode:", resp.StatusCode)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error body:", err)
		return
	}

	log.Println("Body:")
	log.Println(string(data))

	log.Println("Done...")
}
