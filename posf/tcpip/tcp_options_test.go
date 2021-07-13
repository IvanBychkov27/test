package main

import (
	"fmt"
	"github.com/google/gopacket/layers"
	"testing"
)

func Test_decodeTCPOptions(t *testing.T) {
	options := []layers.TCPOption{
		{
			OptionType: TCP_OPT_TIMESTAMP,
			OptionData: []byte("123456789"),
		},
	}

	fmt.Println(decodeTCPOptions(options))
}
