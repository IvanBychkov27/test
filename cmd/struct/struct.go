// Оптимизация структур в Golang для эффективного распределения памяти
// https://nuancesprog.ru/p/11674/
// https://nuancesprog.ru/tag/golang/  - статьи

package main

import (
	"fmt"
	"unsafe"
)

type FooNil struct {
}

type Foo struct {
	aaa bool   // 2 байта
	bbb int32  // 4 байта
	ccc bool   // 2 байта
	s   string // 20 байт
}

type Foo2 struct {
	aaa bool   // 2 байта
	ccc bool   // 2 байта
	bbb int32  // 4 байта
	s   string // 16 байт
}

func main() {
	//f01()
	f02()

}

func f02() {
	x := Foo{}
	y := Foo2{}

	fmt.Print(unsafe.Sizeof(x), "  ") // 12
	fmt.Println(unsafe.Sizeof(y))     // 8
}

func f01() {
	x := &FooNil{}
	y := FooNil{}

	fmt.Print(unsafe.Sizeof(x), "  ") // 8 -  указатель
	fmt.Println(unsafe.Sizeof(y))     // 0
}
