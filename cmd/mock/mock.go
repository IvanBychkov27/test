package main

import (
	"fmt"
)

type DataMock interface {
	sum100() int
	sum200() int
}

type Data struct {
}

func main() {
	d := &Data{}
	fmt.Println(d.sum100())
	fmt.Println(d.sum200())
}

func (d *Data) sum100() int {
	return 100
}

func (d *Data) sum200() int {
	return 200
}
