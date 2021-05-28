package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	url := `http://127.0.0.1:2000/metrics`
	response := getMetricsData(url)
	printArray(metrics(response))
}

func getMetricsData(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	response := ""

	for true {
		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		response += string(bs[:n])
		if n == 0 || err != nil {
			break
		}
	}
	return response
}

func metrics(response string) []string {
	if response == "" {
		return nil
	}
	lines := strings.Split(response, "\n")
	result := make([]string, 0)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "requests{hour=") {
			result = append(result, line)
		}
	}

	return result
}

func printArray(data []string) {
	for _, d := range data {
		fmt.Println(d)
	}
}
