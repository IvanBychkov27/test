// Решето́ Эратосфе́на — алгоритм нахождения всех простых чисел до некоторого целого числа n
// см. "github.com/yourbasic/bit"

package main

import (
	"fmt"
	"github.com/yourbasic/bit"
	"math"
	"math/big"
	"sort"
	"sync"
	"time"
)

type Data struct {
	primeRowNumbersMx sync.Mutex
	primeRowNumbers   map[int]struct{}
}

func main() {
	//n := 10000000 // 30 sec - 80ms (bit)
	//n := 1000000 // 1 sec
	//n := 1000
	//n := 173
	n := 50
	fmt.Println("n =", n)

	timeStart := time.Now()
	//primeNumbers := sieveEratosthenes(n)
	primeN := sieveEratosthenes_bit(n)
	timeJob := time.Now().Sub(timeStart)

	//fmt.Println("prime Numbers =", len(primeNumbers))
	fmt.Println("primeN =", primeN)
	//fmt.Println("primeNumbers =", primeNumbers)

	fmt.Println("timeJob =", timeJob)

	data := exempleCreateBit()
	if !data.Contains(0) {
		fmt.Println(data.Next(0))
	} else {
		fmt.Println(0)
	}
	b := exempleBit(data)
	fmt.Println(b)

	d := exempleCreateMap()
	m := exempleMap(d)
	fmt.Println(m)

}

func exempleCreateBit() *bit.Set {
	return bit.New().AddRange(0, 1000000)
}

func exempleBit(d *bit.Set) bool {
	//d := bit.New().AddRange(0, 1000000)
	//fmt.Println(d, d.Size(), d.Max())

	//d.Add(100)
	//fmt.Println(d, d.Size())

	//for i := 0; i != -1; i = d.Next(i) {
	//	fmt.Print(i)
	//}

	//fmt.Println()
	return d.Contains(155555)
}

func exempleCreateMap() map[int]struct{} {
	d := make(map[int]struct{})
	for i := 0; i < 1000000; i++ {
		d[i] = struct{}{}
	}
	return d
}

func exempleMap(d map[int]struct{}) bool {
	//d := make(map[int]struct{})
	//for i := 0; i < 1000000; i++ {
	//	d[i] = struct{}{}
	//}
	_, ok := d[155555]
	return ok
}

// Решето́ Эратосфе́на через битовые операции (см. пакет "github.com/yourbasic/bit")
func sieveEratosthenes_bit(n int) int {
	sieve := bit.New().AddRange(2, n)
	sqrtN := int(math.Sqrt(float64(n)))
	for p := 2; p <= sqrtN; p = sieve.Next(p) {
		for k := p * p; k < n; k += p {
			sieve.Delete(k)
		}
	}

	//res := make([]int, 0, n/2)
	//for i := 0; i < n/2; i++ {
	//	d := sieve.
	//}
	//fmt.Println(sieve)

	return sieve.Size()
}

// проверка числа на простое целочисленное число
func probablyPrime(n int64) bool {
	if big.NewInt(n).ProbablyPrime(0) {
		fmt.Println(n, "is prime")
		return true
	}
	fmt.Println(n, "is not prime")
	return false
}

// Решето́ Эратосфе́на
func sieveEratosthenes(n int) []int {
	if n < 2 {
		return nil
	}

	rowNumbers, primeNumbers := initialRowNumbers(n)
	//fmt.Println("rowNumbers =", rowNumbers)

	d := &Data{
		primeRowNumbers: primeNumbers,
	}

	wg := &sync.WaitGroup{}
	var counter int
	for i, p := range rowNumbers {
		if p == 2 {
			continue
		}
		if p*p > n {
			break
		}
		wg.Add(1)
		go delElement(wg, p, rowNumbers[i*2:], d)
		counter++
	}

	wg.Wait()

	fmt.Println("counter =", counter)
	return resultPrimeNumbers(d)
}

// удаление элемента если оно делится без остатка
func delElement(wg *sync.WaitGroup, p int, rowNumbers []int, d *Data) {
	defer wg.Done()

	for _, el := range rowNumbers {
		if el%p == 0 {
			d.primeRowNumbersMx.Lock()
			delete(d.primeRowNumbers, el)
			d.primeRowNumbersMx.Unlock()
		}
	}

}

// преобразуем мапу в отсортированный слайс
func resultPrimeNumbers(d *Data) []int {
	d.primeRowNumbersMx.Lock()
	defer d.primeRowNumbersMx.Unlock()

	result := make([]int, 0, len(d.primeRowNumbers))
	for r := range d.primeRowNumbers {
		result = append(result, r)
	}
	sort.Ints(result)
	return result
}

// формируем начальный ряд чисел (от 2ки и все нечетные)
func initialRowNumbers(n int) ([]int, map[int]struct{}) {
	rowNumbers := []int{2}
	primeNumbers := map[int]struct{}{2: {}}

	for i := 3; i < n+1; i += 2 {
		rowNumbers = append(rowNumbers, i)
		primeNumbers[i] = struct{}{}
	}
	return rowNumbers, primeNumbers
}
