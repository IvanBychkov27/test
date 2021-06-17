package main

import (
	"testing"
)

func Benchmark_exempleCreateBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exempleCreateBit()
	}
}

func Benchmark_exempleBit(b *testing.B) {
	data := exempleCreateBit()
	for i := 0; i < b.N; i++ {
		exempleBit(data)
	}
}

func Benchmark_exempleCreateMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exempleCreateMap()
	}
}

func Benchmark_exempleMap(b *testing.B) {
	d := exempleCreateMap()
	for i := 0; i < b.N; i++ {
		exempleMap(d)
	}
}
