package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// http://109.206.162.136:82/api/v1/query?query=node_filesystem_avail_bytes{mountpoint=~"/|/etc/hostname",instance=~"192.168.1.68:9100|192.168.1.69:9100"}/node_filesystem_size_bytes{mountpoint=~"/|/etc/hostname",instance=~"192.168.1.68:9100|192.168.1.69:9100"}*100
// Authorization: Basic c2VydmljZV91c2VyOlFZMmN2aGNwQ0tQQm5Vd3RQZU5KVXBrQw==

type Data struct {
	IP    []string
	Value int
}

func main() {
	networkData := map[string]Data{}

	consulRequestBody := getConsul(`http://127.0.0.1:8500/v1/kv/networks/?raw=1&recurse=1`)
	networks := getNetworks(consulRequestBody)

	for _, network := range networks {
		url := `http://127.0.0.1:8500/v1/kv/` + network + `?raw`
		networkData[network] = Data{IP: getIPs(getConsul(url))}
	}
	//fmt.Println(tntQuerys(networkData))

	urlProm := tntQuerys(networkData)
	prometheusRequestBody := requestInPrometheus(urlProm[0])

	fmt.Println(string(prometheusRequestBody))
}

// получение списка всех networks
func getNetworks(body []byte) []string {
	type data struct {
		Network string `json:"key"`
	}

	consuls := make([]data, 0)
	err := json.Unmarshal(body, &consuls)
	if err != nil {
		fmt.Println("getDataConsul, error unmarshal body", err.Error())
		return nil
	}

	networks := make([]string, 0, len(consuls))
	for _, consul := range consuls {
		if consul.Network == "" {
			continue
		}
		networks = append(networks, consul.Network)
	}
	return networks
}

// получение тела ответа от консула с информацией
func getConsul(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error get url", err.Error())
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error read body", err.Error())
		return nil
	}
	defer resp.Body.Close()

	return body
}

// получаем список уникальных ip без портов из ответа от Consula
func getIPs(body []byte) []string {
	ips := make([]string, 0)
	err := json.Unmarshal(body, &ips)
	if err != nil {
		fmt.Println("getIPs, error unmarshal body", err.Error())
		return nil
	}

	ipData := make(map[string]struct{})
	for _, ip := range ips {
		n := strings.Index(ip, ":")
		if n == -1 {
			continue
		}
		ip = ip[:n]
		ipData[ip] = struct{}{}
	}

	ip := make([]string, 0, len(ipData))
	for i := range ipData {
		ip = append(ip, i)
	}

	return ip
}

// получаем url запросы в prometheus
func tntQuerys(networkData map[string]Data) []string {
	urlQuerys := make([]string, 0, len(networkData))

	for _, data := range networkData {
		ips := ""
		for _, d := range data.IP {
			ips += d + ":9100|"
		}
		ips = ips[:len(ips)-1]
		query := fmt.Sprintf(`http://109.206.162.136:82/api/v1/query?query=node_filesystem_avail_bytes{mountpoint=~"/|/etc/hostname",instance=~"%s"}/node_filesystem_size_bytes{mountpoint=~"/|/etc/hostname",instance=~"%s"}*100`, ips, ips)
		urlQuerys = append(urlQuerys, query)
	}

	return urlQuerys
}

// запрс в Prometheus
func requestInPrometheus(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error NewRequest", err.Error())
		return nil
	}
	req.Header.Set("Authorization", "Basic c2VydmljZV91c2VyOlFZMmN2aGNwQ0tQQm5Vd3RQZU5KVXBrQw==")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error DefaultClient", err.Error())
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error read body", err.Error())
		return nil
	}
	defer resp.Body.Close()

	return body
}
