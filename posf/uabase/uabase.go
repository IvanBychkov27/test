package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("posf/uabase/database/combined.json")
	if err != nil {
		fmt.Println(err)
	}
	if len(data) == 0 {
		fmt.Println("Exit (data = 0)")
		return
	}

	fmt.Println("data = ", len(data))
}
