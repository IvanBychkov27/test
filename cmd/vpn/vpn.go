package main

/*
openvpn --genkey secret pre-shared.key



//urlData := "http://apta.fun:2099/TH"
99.adscompass.ru
2000.99.adscompass.ru
3000.99.adscompass.ru
4000.99.adscompass.ru
*/

import "log"
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	//getURL()
	serverNordVPN()
}

func serverNordVPN() {
	log.Println("start...")

	user := "ysETCpBC8JvFzJnt7SjsJxJC"
	password := "qRhXR6k8yW4pcW71c34ReDW3"
	port := "1080"

	ip, flag := "109.202.99.35", "NL"
	//ip, flag := "165.231.210.155", "US"
	//ip, flag := "196.196.244.3", "SE"
	//ip, flag := "196.196.192.11", "IE"

	//ip, flag := "172.83.40.219", "CA"

	_ = flag
	//urlData := "http://eth0.me"
	urlData := "http://4000.99.adscompass.ru/" + flag

	addrNordVPN := fmt.Sprintf("socks5://%s:%s@%s:%s", user, password, ip, port)
	//addrNordVPN := fmt.Sprintf("udp://%s:%s@%s:%s", user, password, ip, port)

	urlAddrNordVPN, err := url.Parse(addrNordVPN)
	if err != nil {
		log.Println(err)
	}
	log.Println("urlAddrNordVPN =", urlAddrNordVPN)

	transport := &http.Transport{
		Proxy: http.ProxyURL(urlAddrNordVPN),
	}

	//adding the Transport object to the http Client
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 10,
	}

	//generating the HTTP GET request

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

// standart
func getURL() {
	//url := "http://ya.ru"
	url := "https://poker-iv.herokuapp.com"
	//url := "http://apta.fun:2099"

	fmt.Println("url:", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	defer resp.Body.Close()

	fmt.Println("StatusCode: ", resp.StatusCode)
	fmt.Println("Done...")
}

func readParams(resp *http.Response) {

}

func readBody(resp *http.Response) {
	fmt.Println("Body:")
	for {
		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		fmt.Println(string(bs[:n]))
		if n == 0 || err != nil {
			break
		}
	}
}
