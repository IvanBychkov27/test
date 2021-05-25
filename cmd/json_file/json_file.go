package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

func main() {
	jsonData, err := json.Marshal(dataCollectors())
	chk(err)
	fmt.Println(string(jsonData))

	//nameFile := "/home/ivan/projects/test/.debug/data_request/data_requests_" + strconv.Itoa(int(time.Now().Unix())) + ".dat"
	nameFile := ".debug/data_request/data_requests_" + strconv.Itoa(int(time.Now().Unix())) + ".dat"
	err = ioutil.WriteFile(nameFile, jsonData, 0777)
	chk(err)
	fmt.Println("bytes written in", nameFile)
}

func chk(err error) {
	if err != nil {
		fmt.Println("error: ", err)
	}
}

type Collector struct {
	ID         int
	UserID     int
	Key        string
	TimeCreate int
	IP         string
	UA         string
}

// заполняем массив данных Collector
func dataCollectors() []Collector {
	collectors := make([]Collector, 0, 5)

	for i := 10; i < 15; i++ {
		c := Collector{
			ID:         i,
			UserID:     100 + i,
			Key:        "key" + strconv.Itoa(i),
			TimeCreate: int(time.Now().Unix()),
			IP:         "127.0.0." + strconv.Itoa(i),
			UA:         "ua " + strconv.Itoa(100+i),
		}

		collectors = append(collectors, c)
	}
	return collectors
}
