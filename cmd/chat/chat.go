// Запускаем программу и далее запускаем программу клиента netcad из папки chatnet

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:2121")
	if err != nil {
		log.Println("error listener", err)
		os.Exit(1)
	}
	log.Println("listen... 127.0.0.1:2121")

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error conn", err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // канал исходящих сообщений
var (
	enterning = make(chan client)
	leaving   = make(chan client)
	messages  = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool) // все подключенные клиенты
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-enterning:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You" + who
	messages <- who + " connection"
	enterning <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <- ch
	messages <- who + " on connection"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
