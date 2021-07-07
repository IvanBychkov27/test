package main

import (
	"fmt"
	"testing"
)

type MockData struct {
}

func (m *MockData) sum100() int {
	return -100
}

func (m *MockData) sum200() int {
	return -200
}

func Test_sum(t *testing.T) {
	m := MockData{}
	fmt.Println(m.sum100())
	fmt.Println(m.sum200())
}
