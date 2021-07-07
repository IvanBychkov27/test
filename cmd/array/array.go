package main

import (
	"fmt"
	"github.com/yourbasic/bit"
)

func main() {
	a := []int{0, 2, 4, 6, 8, 10, 12, 14}
	b := []int{1, 3, 5, 7, 9, 11, 13, 15}

	printSlice(sumSlice01(a, b))

	printSlice(sumSlice02(a, b))

}

func sumSlice01(x []int, y []int) []int {

	s := make([]int, len(x)+len(y))

	var kx, ky int

	if len(x) == 0 {
		s = s[:0]
		s = append(s, y[:]...)
		return s
	}

	if len(y) == 0 {
		s = s[:0]
		s = append(s, x[:]...)
		return s
	}

	for i := 0; i < len(s); i++ {
		if x[kx] <= y[ky] {
			s[i] = x[kx]
			if kx == len(x)-1 {
				s = s[:i+1]
				s = append(s, y[ky:]...)
				break
			}
			kx++
		} else {
			s[i] = y[ky]
			if ky == len(y)-1 {
				s = s[:i+1]
				s = append(s, x[kx:]...)
				break
			}
			ky++
		}
	}
	return s
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sumSlice02(x []int, y []int) []int {
	lenX := len(x)
	lenY := len(y)

	if lenX == 0 {
		return y
	}

	if lenY == 0 {
		return x
	}

	var r bit.Set

	j := 0
	for _, a := range x {
		r.Add(a)
		if j < lenY {
			r.Add(y[j])
			j++
		}
	}

	res := make([]int, 0, r.Size())

	for e := 0; e != -1; e = r.Next(e) {
		res = append(res, e)
	}
	return res
}

// находим первый элемент массива, если массив пуст, то -1
func firstElement(x bit.Set) int {
	if x.Empty() {
		return -1
	}
	i := 0
	for {
		if x.Contains(i) {
			break
		}
		i++
		continue
	}
	return i
}
