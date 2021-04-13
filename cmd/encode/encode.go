package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	fmt.Println(md5Encode("777", ""))
	fmt.Println(md5Encode("777", "1"))
	fmt.Println(md5Encode("777", "iv"))

	t1 := md5Encode("777", "iv")
	t2 := md5Encode("777", "iv")

	if t1 == t2 {
		fmt.Println("ok")
	} else {
		fmt.Println("bad...")
	}
}

func md5Encode(data, salt string) string {
	data += salt
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}
