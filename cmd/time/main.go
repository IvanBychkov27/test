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

	dif := time.Since(t)
	fmt.Println("dif =", dif)

	var f float64
	f = float64(dif)
	fmt.Println("nano sec:", f)

	ds := []int{1, 2, 3, 4, 5}
	for _, d := range ds {
		fmt.Printf("d = %d \n", d)
	}

	query := "ALTER TABLE %s.%s DROP PARTITION '%s'"
	query = fmt.Sprintf(query, "ch.task.Database", "ch.task.Table", "partition")
	fmt.Println("query =", query)

}
