package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Combined struct {
	Ts                    int    `json:"ts"`
	IPTTL                 int    `json:"ip_ttl"`
	IPDf                  int    `json:"ip_df"`
	IPMf                  int    `json:"ip_mf"`
	TCPWindowSize         int    `json:"tcp_window_size"`
	TCPFlags              int    `json:"tcp_flags"`
	TCPAck                int    `json:"tcp_ack"`
	TCPSeq                int    `json:"tcp_seq"`
	TCPHeaderLength       int    `json:"tcp_header_length"`
	TCPUrp                int    `json:"tcp_urp"`
	TCPOptions            string `json:"tcp_options"`
	TCPWindowScaling      int    `json:"tcp_window_scaling"`
	TCPTimestamp          int    `json:"tcp_timestamp"`
	TCPTimestampEchoReply int    `json:"tcp_timestamp_echo_reply"`
	TCPMss                int    `json:"tcp_mss"`
	NavigatorUserAgent    string `json:"navigatorUserAgent"`
	Platform              struct {
		Type   string `json:"type"`
		Vendor string `json:"vendor"`
		Model  string `json:"model"`
	} `json:"platform"`
	DeviceMemory        float32 `json:"devicememory"`
	HardwareConcurrency int     `json:"hardwareconcurrency"`
	Os                  struct {
		Name        string `json:"name"`
		Version     string `json:"version"`
		VersionName string `json:"versionname"`
	} `json:"os"`
}

func main() {
	fileName := "posf/uabase/database/combined.json"
	data := readFile(fileName)
	fmt.Println("data =", len(data), "bytes")

	result := dataUnmarshal(data)
	fmt.Println("result =", len(result))
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
