package main

import (
	"context"
	"fmt"
	"time"
)

type Item struct {
	campaignID int
	balance    int64

	bttl    int64
	timeout int64

	updating     int64
	cancel       context.CancelFunc
	ctx          context.Context
	toUpdateChan chan<- *Item
}

func main() {
	i := &Item{}
	i.ctx, i.cancel = context.WithCancel(context.Background())

	if i.cancel != nil {
		fmt.Println("start i.cancel != nil")
	} else {
		fmt.Println("start i.cancel == nil")
	}

	//time.Sleep(time.Second)
	//if i.cancel != nil {
	//	i.cancel()
	//}
	//time.Sleep(time.Second)

	go func(i *Item, interval time.Duration) {
		fmt.Println("intrval =", interval)
		select {
		case <-i.ctx.Done():
			fmt.Println("ctx.Done")
			if i.cancel != nil {
				fmt.Println("func(ctx.Done): i.cancel != nil")
			} else {
				fmt.Println("func(ctx.Done): i.cancel == nil")
			}
			return
		case <-time.After(interval):
			fmt.Println("time.After")
		}

		if i.cancel != nil {
			fmt.Println("func(): i.cancel != nil")
		}
		fmt.Println("exit func")

	}(i, time.Second*5)

	if i.cancel != nil {
		fmt.Println("i.cancel != nil")
	} else {
		fmt.Println("i.cancel == nil")
	}

	//time.Sleep(time.Second * 2)
	//i.cancel()
	time.Sleep(time.Second)

	fmt.Println("Для завершения жми Enter")
	var s string
	fmt.Scanln(&s)

	i.cancel()

	if i.cancel != nil {
		fmt.Println("exit: i.cancel != nil")
	} else {
		fmt.Println("exit: i.cancel == nil")
	}

	time.Sleep(time.Second)
	fmt.Println("Done")
}

//
//func main() {
//	var ctx context.Context
//	i := &Item{}
//
//	ctx, i.cancel = context.WithCancel(context.Background())
//
//	go func(ctx context.Context, i *Item, interval time.Duration) {
//		fmt.Println("intrval =", interval)
//		select {
//		case <-ctx.Done():
//			fmt.Println("ctx.Done")
//			if i.cancel != nil {
//				fmt.Println("func(): i.cancel != nil")
//			} else {
//				fmt.Println("func(): i.cancel == nil")
//			}
//			return
//		case <-time.After(interval):
//			fmt.Println("time.After")
//		}
//
//		if i.cancel != nil {
//			fmt.Println("func(): i.cancel != nil")
//		}
//		fmt.Println("exit func")
//
//	}(ctx, i, time.Second*15)
//
//	if i.cancel != nil {
//		fmt.Println("i.cancel != nil")
//	} else {
//		fmt.Println("i.cancel == nil")
//	}
//
//	var s string
//	fmt.Scanln(&s)
//
//	i.cancel()
//
//	if i.cancel != nil {
//		fmt.Println("exit: i.cancel != nil")
//	} else {
//		fmt.Println("exit: i.cancel == nil")
//	}
//
//	time.Sleep(time.Second)
//	fmt.Println("Done")
//}
