package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	ff := FFPair{
		Name: "name",
		ID:   100,
	}
	fdata := []FFPair{
		{
			Name: "name1",
			ID:   100,
		},
		{
			Name: "name2",
			ID:   200,
		},
	}

	m := CreateDataFFPair()

	m[ff] = fdata

	fmt.Println("m:", m)

	json_data, err := json.Marshal(m)
	chk(err)
	fmt.Println(string(json_data))

}

type FFPair struct {
	Name string
	ID   int
}

func NewFFPair() FFPair {
	return FFPair{}
}

func CreateDataFFPair() map[FFPair][]FFPair {
	return make(map[FFPair][]FFPair)
}

func chk(err error) {
	if err != nil {
		fmt.Println("error: ", err)
	}
}
