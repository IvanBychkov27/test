package main

import (
	"fmt"
	"time"
)

func main() {
	namefile := "file.buffer"
	t := time.Now()
	str := t.Format("20060102-150405-")
	str += namefile
	fmt.Println(str)

	fmt.Println(time.Now().Format("20060102-150405-"))

	fmt.Println()
	str = fmt.Sprintf("%s%s", time.Now().Format("20060102-150405-"), namefile)
	fmt.Println(str)

	date := time.Now().AddDate(0, 0, -2).Round(time.Hour * 24)
	ageDays := (int(time.Now().Sub(date).Hours()) / 24) + 1
	fmt.Println("ageDays =", ageDays)

	ageDays = (int(time.Since(date).Hours()) / 24) + 1
	fmt.Println("ageDays =", ageDays)

}
