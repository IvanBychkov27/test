package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Start...")
	fmt.Println("для завершения жми: Ctrl+C")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	_, cansel := context.WithTimeout(context.Background(), time.Second)

	defer exitMSG(cansel)

	fmt.Println()
	fmt.Println("Done")
}

func exitMSG(cansel context.CancelFunc) {
	defer cansel()

	fmt.Println()
	fmt.Print("exit: ")

	i := 5
	for {
		select {
		case <-time.After(time.Second):
			fmt.Print(i, " ")
			i--

		}
		if i == 0 {
			fmt.Println()
			break
		}
	}
}
