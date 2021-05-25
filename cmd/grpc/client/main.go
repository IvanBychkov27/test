package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"os"
	"test/cmd/grpc/api/grpc/api"
)

func main() {
	var addr string
	flag.StringVar(&addr, "a", "127.0.0.1:2000", "listen address")
	flag.Parse()

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("error dial, %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := api.NewAccountantClient(conn)

	data := api.ManData{
		Ip:   1,
		Name: "Test",
	}

	result, err := client.GetMan(context.Background(), &data)
	if err != nil {
		log.Printf("error getman, %v", err)
		os.Exit(1)
	}

	log.Printf("result: addres: %s, data: %s", result.Addr, result.Tel)
}
