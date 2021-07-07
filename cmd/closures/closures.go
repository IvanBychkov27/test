// замыкания
package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// замыкания
func closures() {
	pos, neg := adder(), adder()
	for i := 1; i < 11; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func main() {
	Perm([]rune("()()"), func(a []rune) { fmt.Println(string(a)) })

	//closures()
}

// Perm вызвает f с каждой пермутацией a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Пермутируем значения в индексе i на len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// Мы используем типы rune для обработки и срезов, и строк. Runes являются кодовыми точками из Unicode, а значит могут парсить строки и срезы одинаково.
