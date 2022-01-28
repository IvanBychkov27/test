package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	d := Data{
		ID:   1,
		Name: "Ivan",
	}

	res, err := json.Marshal(&d)
	if err != nil {
		fmt.Println("error json marshal", err.Error())
	}

	fmt.Println("res =", string(res))
}
