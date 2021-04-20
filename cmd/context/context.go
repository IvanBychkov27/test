package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(3)

	go f(cancel, wg)
	go g(ctx, wg)
	go p(ctx, wg)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)
	signal.Notify(signals, syscall.SIGINT)

	select {
	case sig := <-signals:
		fmt.Println("terminate by signal", sig)
		cancel()
	case <-ctx.Done():
		fmt.Println("terminate by context")
	}

	wg.Wait()

}

func f(cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer cancel()
	defer wg.Done()

	for i := 0; i < 7; i++ {
		fmt.Println("f =", i)
		time.Sleep(time.Second)
	}
	fmt.Println("func f done")
}

func g(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 10
	tik := time.NewTicker(time.Second * 2)
	for {
		fmt.Println("g =", i)
		select {
		case <-tik.C:
		case <-ctx.Done():
			fmt.Println("func g done: <-ctx.Done()")
			return
		}
		i++
	}
}

func p(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	sig := make(chan int)
	n := make(chan int)
	t := 2
	go tikSig(sig, n)
	for {
		n <- t
		select {
		case <-sig:
			fmt.Println("func p done: sig")
			t = 1
		case <-ctx.Done():
			fmt.Println("func p done: <-ctx.Done()")
			return
		}
	}
}

func tikSig(sig, n chan int) {
	for {
		time.Sleep(time.Second * time.Duration(<-n))
		sig <- 2
	}
}
