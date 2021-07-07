package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sumSlice01(t *testing.T) {
	a := []int{0, 2, 4}
	b := []int{1, 3, 5}

	res := sumSlice01(a, b)

	expected := []int{0, 1, 2, 3, 4, 5}
	assert.Equal(t, expected, res)
}

func Benchmark_sumSlice01(b *testing.B) {
	x, y := xySet(10000)
	for i := 0; i < b.N; i++ {
		sumSlice01(x, y)
	}
}

func xySet(n int) ([]int, []int) {
	x := make([]int, 0, n)
	y := make([]int, 0, n)

	for i := 0; i < n; i++ {
		x = append(x, i)
		i++
		y = append(y, i)
	}
	return x, y
}

func Test_sumSlice02(t *testing.T) {
	a := []int{0, 2, 4}
	b := []int{1, 3, 5}

	res := sumSlice02(a, b)

	expected := []int{0, 1, 2, 3, 4, 5}
	assert.Equal(t, expected, res)
}

func Benchmark_sumSlice02(b *testing.B) {
	x, y := xySet(10000)
	for i := 0; i < b.N; i++ {
		sumSlice02(x, y)
	}
}
