// Запуск генерации в папке с файлом api.proto
// protoc -I=. api.proto --go_out=plugins=grpc:.

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"test/cmd/grpc/api/grpc/api"
)

func main() {
	var addr string
	flag.StringVar(&addr, "a", "127.0.0.1:2000", "listen address")
	flag.Parse()

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("error listen, %v", err)
		os.Exit(1)
	}
	log.Printf("listen %s", addr)

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	api.RegisterAccountantServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	err = grpcServer.Serve(ln)
	if err != nil {
		log.Printf("error grpc, %v", err)
		os.Exit(1)
	}

	log.Printf("done")
}

type server struct{}

type User struct {
	Id      int
	Name    string
	Telefon string
	Place   Address
}

type Address struct {
	City      string
	Street    string
	House     int
	Apartment int
}

func (s *server) GetMan(context.Context, *api.ManData) (*api.ManRequest, error) {

	a := []Address{
		{"Bryansk", "22 sezd", 14, 93},
		{"Bryansk", "sov", 4, 82},
		{"Bryansk", "sov", 4, 82},
	}

	users := []User{
		{2, "Andrey", "9208647777", a[0]},
		{3, "Tanya", "9102379977", a[1]},
		{4, "Olya", "9107432216", a[2]},
	}

	json_data, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	result := &api.ManRequest{
		Addr: "Bryansk",
		Tel:  string(json_data),
	}
	return result, nil
}

func (s *server) mustEmbedUnimplementedAccountantServer() {
}
