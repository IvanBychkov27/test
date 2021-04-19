package main

import (
	crypto "crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/fatih/color"
	"math"
	"math/big"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

type myData struct {
	data map[int]int
}

func main() {

	//workIncome()

	work19()
}

func work19() {
	// feature/cors-access-control-allow-headers
	// rw.Header().Add("Access-Control-Allow-Headers", "*")

	m := make(map[string]struct{})
	m["a"] = struct{}{}
	fmt.Println(m)

	if len(m) > 0 {
		_, ok := m["a"]
		fmt.Println("len(m) > 0  ok =", ok)
	}

}

// =======CAP добавление и удаление ================================
type Cap struct {
	memCap  map[string]int
	capData map[string]int
}

func newCap() *Cap {
	c := &Cap{
		memCap:  make(map[string]int),
		capData: make(map[string]int),
	}
	return c
}

func work18() {
	c := newCap()
	c.capData["1"] = int(time.Now().Unix()) + 3
	c.capData["2"] = int(time.Now().Unix()) + 6
	c.capData["3"] = int(time.Now().Unix()) + 9

	go c.getCap(5)
	go c.delCap(1)
	go c.setCap(2)

	var s string
	fmt.Scanln(&s)
}

func (c *Cap) setCap(interval int) {
	key := 3
	tik := time.NewTicker(time.Second * time.Duration(interval))
	for range tik.C {
		key++
		c.capData[strconv.Itoa(key)] = int(time.Now().Unix()) + 7
	}
}

func (c *Cap) delCap(interval int) {
	tik := time.NewTicker(time.Second * time.Duration(interval))
	for range tik.C {
		for key, val := range c.memCap {
			if val > int(time.Now().Unix()) {
				continue
			}
			if val != 0 {
				delete(c.memCap, key)
				delete(c.capData, key)
				fmt.Println("del:", key)
			}
		}
		fmt.Println("memCap: ", c.memCap)
	}
}

func (c *Cap) getCap(intervalUpdateCap int) {
	for k, v := range c.capData {
		c.memCap[k] = v
	}
	t := time.Now().Unix()
	tik := time.NewTicker(time.Second * time.Duration(intervalUpdateCap))
	for range tik.C {
		for k, v := range c.capData {
			if v > int(t) {
				c.memCap[k] = v
			}
		}
		t = time.Now().Unix()
	}
}

//=== Доходность ======================
type compains struct {
	id         int
	imp        int
	clicks     int
	priceClick float64
	gain       float64
}

func workIncome() {
	comp := []compains{
		{id: 1, imp: 1000, clicks: 90, priceClick: 0.011},
		{id: 2, imp: 1000, clicks: 83, priceClick: 0.012},
		{id: 3, imp: 1000, clicks: 77, priceClick: 0.013},
		{id: 4, imp: 1000, clicks: 60, priceClick: 0.014},
		{id: 5, imp: 1000, clicks: 50, priceClick: 0.015},
		{id: 6, imp: 1000, clicks: 10, priceClick: 0.055},
	}

	var maxGain, maxPriceClick, gainMaxPriceClick float64
	fmt.Println(" ID  Imp   Clicks  Price   Доход ")
	line := "---------------------------------"
	fmt.Println(line)
	flag := true
	for i, c := range comp {
		if c.imp/2 > c.clicks && c.imp > 0 {
			kImp := float64(c.imp) / float64(1000)
			gain := float64(c.clicks) * c.priceClick / kImp
			comp[i].gain = gain
			if maxGain < gain {
				maxGain = gain
			}
			if maxPriceClick < c.priceClick {
				maxPriceClick = c.priceClick
				gainMaxPriceClick = gain
			}
			if c.imp < 1000 {
				flag = false
			}
			//fmt.Printf(" %d %6d %6d   %.3f   %.3f \n", c.id, c.imp, c.clicks, c.priceClick, gain)
		}
	}

	if flag {
		sort.Slice(comp, func(i, j int) bool {
			return comp[i].gain > comp[j].gain // сортировка по убыванию доходности
		})
	} else {
		sort.Slice(comp, func(i, j int) bool {
			return comp[i].priceClick > comp[j].priceClick // сортировка по убыванию стоимости клика
		})
	}

	for _, c := range comp {
		fmt.Printf(" %d %6d %6d   %.3f   %.3f \n", c.id, c.imp, c.clicks, c.priceClick, c.gain)
	}

	fmt.Println(line)
	if flag {
		fmt.Printf("доп.доход = %.4f \n", maxGain-gainMaxPriceClick)
	} else {
		fmt.Println("сотрировка по Price")
	}
}

func work17() {

	subscriptionDate := 1617150000
	now := int(time.Now().Unix())
	res := (now - subscriptionDate) / 86400 // кол-во дней подписки

	//fmt.Println(now)
	fmt.Println("дней: ", res)

}

// ------------------------------------------
type dan struct {
	a int
	b int
}

func work16() {
	d := &dan{}

	fmt.Println("Start...")

	go f01(d) //  01 &{5 10}

	// горутины запускаются не сразу!!!
	// нужно быть внимательным при изменении данных для горутин
	// НЕ ИСПОЛЬЗУЙ ОДНИ И ТЕЖЕ ДАННЫЕ ДЛЯ РАЗНЫХ ГОРУТИН!!!

	d.a = 5
	d.b = 10

	go f02(d) //  02 &{0 0}

	time.Sleep(time.Millisecond * 3000)
	fmt.Println("Done...")
}

func f01(d *dan) {
	time.Sleep(time.Millisecond * 800)
	fmt.Println("01", d)
	d.a = 0
	d.b = 0
}

func f02(d *dan) {
	time.Sleep(time.Millisecond * 1200)
	fmt.Println("02", d)
}

// -------------------------------------

func work15() {
	const (
		percent_n = 0.20
		percent_k = 0.92
	)
	var (
		result float64 = 0
		res    float64 = 15
		t      float64 = 100
	)

	if res/t < percent_n {
		result = (-1) * ((res/t - percent_n) / percent_n)
	}
	if res/t > percent_k {
		result = (res/t - percent_k) / (1 - percent_k)
	}
	//result[v.key] = ((v.cnt / t) - percent)/(1 - percent)
	fmt.Println("result =", result)
}

//==========================================
// sort
type cad struct {
	//a int
	//b string
	n int
	m string
}

func work14() {
	c := []cad{{n: 2, m: "d"}, {n: 2, m: "c"}, {n: 3, m: "d"}, {n: 2, m: "b"}, {n: 2, m: "a"}, {n: 3, m: "c"}, {n: 3, m: "b"}, {n: 3, m: "a"}}
	fmt.Println(c)
	sort.Slice(c, func(i, j int) bool {
		return c[i].n < c[j].n
	})
	fmt.Println(c)

	res := sortStruct(c)
	fmt.Println(res)

	//n := 4
	//rand.Shuffle(n, func(i, j int) {
	//	i = i + n
	//	j = j + n
	//	res[i], res[j] = res[j], res[i]
	//})
	//fmt.Println(res)

}

func sortStruct(c []cad) []cad {
	res := make([]cad, 0)
	temp := make([]cad, 0)
	count := 0
	ln := len(c) - 1
	for idx := 0; idx < ln; idx++ {
		flag := c[idx].n == c[idx+1].n
		if flag {
			count++
			if idx != ln-1 {
				continue
			}
			flag = false
			idx++
		}
		if !flag {
			if count > 0 {
				temp = append(temp, c[idx-count:idx+1]...)
				sort.Slice(temp, func(i, j int) bool {
					return temp[i].m < temp[j].m
				})
				res = append(res, temp...)
				temp = nil
				count = 0
			} else {
				res = append(res, c[idx])
			}
			if idx+1 == ln {
				res = append(res, c[idx+1])
			}
		}
	}
	return res
}

// ========================================
func work13() {
	var t = []byte{123, 34, 105, 100, 34, 58, 34, 84, 69, 83, 84, 34, 44, 34, 104, 111, 117, 114, 34, 58, 55, 55, 55, 44, 34, 99, 111, 117, 110, 116, 114, 121, 34, 58, 34, 82, 117, 115, 34, 44, 34, 114, 95, 99, 108, 105, 99, 107, 34, 58, 102, 97, 108, 115, 101, 44, 34, 114, 95, 105, 109, 112, 34, 58, 102, 97, 108, 115, 101, 44, 34, 114, 95, 118, 105, 101, 119, 34, 58, 102, 97, 108, 115, 101, 44, 34, 97, 100, 118, 101, 114, 116, 105, 115, 101, 114, 95, 99, 97, 115, 104, 98, 97, 99, 107, 34, 58, 48, 125}
	fmt.Println(string(t))

	text := "123 34 105 100 34 58 34 84 69 83 84 34 44 34 104 111 117 114 34 58 55 55 55 44 34 99 111 117 110 116 114 121 34 58 34 82 117 115 34 44 34 114 95 99 108 105 99 107 34 58 102 97 108 115 101 44 34 114 95 105 109 112 34 58 102 97 108 115 101 44 34 114 95 118 105 101 119 34 58 102 97 108 115 101 44 34 97 100 118 101 114 116 105 115 101 114 95 99 97 115 104 98 97 99 107 34 58 48 125"
	ss := strings.Split(text, " ")
	res := ""
	for _, s := range ss {
		i, _ := strconv.Atoi(s)
		res += string(i)
	}
	fmt.Println(res)
}

//======================================
func work12() {
	tStart := time.Now()

	var x, y uint64
	x = 1
	y = ^x

	fmt.Println("x =", x)
	fmt.Println("y =", y)
	s := strconv.FormatUint(y, 10)
	r := ""
	k := 0
	for i := len(s) - 1; i >= 0; i-- {
		if k == 3 {
			r = string(s[i]) + " " + r
			k = 0
		} else {
			r = string(s[i]) + r
		}
		k++
	}
	fmt.Println("Y =", r)
	tEnd := time.Since(tStart)
	fmt.Println("time1 =", tEnd)

	tStart = time.Now()
	k = 0
	for {
		k++
		if k == 1000000000 {
			break
		}
	}
	tEnd = time.Since(tStart)
	fmt.Println("time2 =", tEnd)

}

func work11() {
	var f, res float64
	f = 2.9999
	n := int(f)
	fmt.Println("n = ", n)

	res = math.Round(f)
	fmt.Printf("res = %.2f \n", res)

}

//==== Случайное число Crypto=================
func work10() {
	var randNum int
	for i := 0; i < 3; i++ {
		criptoNum := newCryptoRand()
		randNum = rand.Intn(1000)
		fmt.Printf("randNum = %d, criptoNum = %d \n", randNum, criptoNum)
	}
	randNum = 100 + rand.Intn(900) // от 100 до 1000
	fmt.Println("randNum =", randNum)

	c := 'a' + rune(rand.Intn('z'-'a'+1)) // 'a' ≤ c ≤ 'z'
	fmt.Println("c =", string(c))

}
func newCryptoRand() int64 {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(1000000000))
	if err != nil {
		panic(err)
	}
	return safeNum.Int64()
}

// === Изменение времени тикера ============================
type chTest struct {
	ch chan int
}

func work09() {
	w := chTest{ch: make(chan int)}
	//ch := make(chan int)

	go tik(w.ch)
	go setCh(w.ch)

	var s string
	fmt.Scanln(&s)
}
func tik(ch chan int) {
	var interval int
	interval = <-ch
	t := time.NewTicker(time.Second * time.Duration(interval))
	for {
		select {
		case interval = <-ch:
			t.Stop()
			t = time.NewTicker(time.Second * time.Duration(interval))
			fmt.Println()
		case <-t.C:
			fmt.Print(time.Now().Format("05"), " ")
		}
	}
}
func setCh(ch chan int) {
	var t int
	t = 1
	for {
		ch <- t
		t++
		time.Sleep(time.Second * 5)
		ch <- t
		t--
		time.Sleep(time.Second * 5)
	}
}

//==================================================
func work08() {
	s := "tEsT"
	t := "TeSt"

	b := strings.EqualFold(s, t)

	fmt.Println(b) // result = true
}

//===================================================
func work07() {
	m := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	for i := range m {
		fmt.Printf("%d", m[i])
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("\b\b")
	}
}

//===================================================
func work06() {
	m := map[int]int{1: 10, 2: 20, 3: 30, 4: 40, 5: 50, 6: 60}
	m[1]++
	fmt.Println(m)

	for i := range m {
		fmt.Printf(" m[%d] = %d\n", i, m[i])
	}

}

//==================================================
func work05() {
	fmt.Println(work05_func(1, 1))
}
func work05_func(a, b int) bool {
	return a == b
}

//===================================================
func work04() {
	var data float64
	data = 3.1415926535
	s := strconv.FormatFloat(data, 'f', -1, 64)
	fmt.Println("s =", s)
}

//==================================================
//b64 "encoding/base64" - кодировка и раскодировка
func work03() {
	data := []byte("Hello!")
	fmt.Println(string(data))

	sEnc := base64.StdEncoding.EncodeToString(data)
	fmt.Println(sEnc)
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	data = []byte("127.0.0.1")
	fmt.Println(string(data))
	uEnc := base64.URLEncoding.EncodeToString(data)
	fmt.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}

//================================================
func work02() {
	//rand.Seed(time.Now().UTC().UnixNano())
	data := &myData{make(map[int]int)}
	newData := &myData{make(map[int]int)}

	n := 10

	data.writeData(n, 1)
	data.printData("data1")

	newData.writeData(n, 2)
	newData.printData("data2")

	data.updateInfo(newData)
	data.printData("data1")
}

func (d *myData) updateInfo(newData *myData) {
	resultData := &myData{data: make(map[int]int)}
	for key, value := range newData.data {
		if v, ok := d.data[key]; ok {
			value = v
		}
		resultData.data[key] = value
	}
	d.data = resultData.data
}

func (d *myData) printData(msg string) {
	fmt.Printf("%s = %v len = %d\n", msg, d.data, len(d.data))
}

func (d *myData) writeData(randN int, value int) {
	for i := 0; i < randN; i++ {
		key := rand.Intn(randN)
		d.data[key] = value
	}
}

//========================================

func work01() {
	color.Set(color.FgBlue)
	color.Red("Start...")
	color.Set(color.FgBlue)

	n := 10
	data := make(map[int]int)
	newData := make(map[int]int)
	resultData := make(map[int]int)

	setData(n, 1, data)
	fmt.Printf("data = %v\nlen  = %d\n", data, len(data))

	setData(n, 2, newData)
	fmt.Printf("newData = %v\nlen  = %d\n", newData, len(newData))

	updateData(data, newData, resultData)
	data = nil
	data = resultData
	fmt.Printf("data = %v\nlen  = %d\n", data, len(data))

	color.Unset()
	color.Red("Done...")
}

func updateData(data, newData, resultData map[int]int) {
	for key, value := range newData {
		if v, ok := data[key]; ok {
			value = v
		}
		resultData[key] = value
	}
}

func setData(n int, value int, data map[int]int) {
	for i := 0; i < n; i++ {
		key := rand.Intn(n)
		data[key] = value
	}
}

//==========================================================

// func greatComDiv(masN []int) int { // наибольший общий делитель
//	if len(masN) < 2 {
//		return 1
//	}
//
//	allMasFactors := make([][]int, 0, len(masN))
//	for _, n := range masN {
//		masFactors := factors(n)
//		allMasFactors = append(allMasFactors, masFactors)
//	}
//	nod := 1
//	for i := 0; i < len(allMasFactors[0]); i++ {
//		n := allMasFactors[0][i]
//		k := 1
//		for m := 1; m < len(allMasFactors); m++ {
//			for j := 0; j < len(allMasFactors[m]); j++ {
//				if n == allMasFactors[m][j] {
//					allMasFactors[m][j] = 0
//					k++
//					break
//				}
//			}
//			if k != m+1 {
//				break
//			}
//		}
//		if k == len(allMasFactors) {
//			nod = nod * n
//		}
//	}
//	return nod
// }
//
// func factors(n int) []int { // разложение числа на простые множители
//	if n < 1 {
//		return nil
//	}
//	if n == 1 {
//		return []int{1}
//	}
//	div := 2
//	result := make([]int, 0, 10)
//	for n > 1 {
//		for n%div == 0 {
//			result = append(result, div)
//			n = n / div
//
//		}
//		if div == 2 {
//			div++
//		} else {
//			div += 2
//		}
//	}
//	return result
// }
