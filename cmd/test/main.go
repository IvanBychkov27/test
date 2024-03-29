package main

import (
	"bytes"
	crypto "crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/beevik/etree"
	"github.com/fatih/color"
	"math"
	"math/big"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type myData struct {
	data map[int]int
}

func main() {
	//workIncome()
	//n := 90
	//b := []int{5, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	//w1 := work22(b, n)
	//w2 := work22_my(b, n)
	//fmt.Println(w1)
	//fmt.Println(w2)

	//data := []byte("GET /?uin=999AAABBBDDD HTTP/1.1\r\nHost: 127.0.0.1:2999\r\nUser-Agent: Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:90.0) Gecko/20100101 Firefox/90.0\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\r\nAccept-Language: ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3\r\nAccept-Encoding: gzip, deflate\r\nConnection: keep-alive\r\nCookie: i18n_redirected=en\r\nUpgrade-Insecure-Requests: 1\r\nSec-Fetch-Dest: document\r\nSec-Fetch-Mode: navigate\r\nSec-Fetch-Site: none\r\nSec-Fetch-User: ?1\r\n\r\n")
	//
	//fmt.Println(work30(string(data)))
	//fmt.Println(work31(data))

	work32()
}

func work32() {

	SYN := false
	PSH := false

	if SYN || PSH {
		fmt.Println("1 SYN =", SYN, "PSH =", PSH)
	}

	if !SYN && !PSH {
		fmt.Println("2 SYN =", !SYN, "PSH =", !PSH)
	} else {
		fmt.Println("2 Ok")
	}

	fmt.Println("Done...")
}

func work31(data []byte) string {
	var uin []byte
	searchStr := []byte("uin=")
	//data := []byte("GET /?uin=999AAABBBDDD HTTP/1.1\r\nHost: 127.0.0.1:2999\r\nUser-Agent: Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:90.0) Gecko/20100101 Firefox/90.0\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\r\nAccept-Language: ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3\r\nAccept-Encoding: gzip, deflate\r\nConnection: keep-alive\r\nCookie: i18n_redirected=en\r\nUpgrade-Insecure-Requests: 1\r\nSec-Fetch-Dest: document\r\nSec-Fetch-Mode: navigate\r\nSec-Fetch-Site: none\r\nSec-Fetch-User: ?1\r\n\r\n")

	if !bytes.Contains(data, searchStr) {
		return ""
	}
	idx := bytes.Index(data, searchStr) + len(searchStr)
	for i := idx; i < len(data); i++ {
		d := data[i]
		if d == ' ' {
			break
		}
		uin = append(uin, d)
	}

	//fmt.Println("UIN =", string(uin))
	return string(uin)
}

func work30(data string) string {
	// получим uin из данных Payload

	var uin string
	separator := " "
	searchStr := "uin="

	//data := "GET /?uin=999AAABBBDDD HTTP/1.1\r\nHost: 127.0.0.1:2999\r\nUser-Agent: Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:90.0) Gecko/20100101 Firefox/90.0\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\r\nAccept-Language: ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3\r\nAccept-Encoding: gzip, deflate\r\nConnection: keep-alive\r\nCookie: i18n_redirected=en\r\nUpgrade-Insecure-Requests: 1\r\nSec-Fetch-Dest: document\r\nSec-Fetch-Mode: navigate\r\nSec-Fetch-Site: none\r\nSec-Fetch-User: ?1\r\n\r\n"

	if !strings.Contains(data, "uin=") {
		return ""
	}

	var idx int
	dataStr := strings.Split(data, separator)
	for _, str := range dataStr {
		if !strings.Contains(str, searchStr) {
			continue
		}

		idx = strings.Index(str, searchStr) + len(searchStr)
		if idx != -1 {
			uin = str[idx:]
			break
		}
	}

	//fmt.Println("UIN =", uin)
	return uin
}

func work29() {
	srcs := make([][]byte, 0)
	src := []byte("1")
	srcs = append(srcs, src)
	srcs = append(srcs, []byte("2"))

	for i := range srcs {
		fmt.Println(i, "=", string(srcs[i]))
	}
}

func work28() {
	user := "service_user"
	password := "QY2cvhcpCKPBnUwtPeNJUpkC"
	expectation := "c2VydmljZV91c2VyOlFZMmN2aGNwQ0tQQm5Vd3RQZU5KVXBrQw=="

	data := user + ":" + password
	fmt.Println(data)

	res := base64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println(res)

	if res == expectation {
		fmt.Println("ok!")
	} else {
		fmt.Println("not equal")
	}

}

var countGo int64

// подсчет запущенных горутин
func work27() {
	var i int
	var wg sync.WaitGroup
	t := time.NewTicker(time.Second)
	for {
		fmt.Print(atomic.LoadInt64(&countGo), " ")

		wg.Add(1)
		go g(&wg)

		i++
		if i > 10 {
			break
		}
		<-t.C
	}
	wg.Wait()
	fmt.Print("Done... ")
}

func g(wg *sync.WaitGroup) {
	defer wg.Done()
	var i int
	atomic.AddInt64(&countGo, 1)
	for {
		i++
		time.Sleep(time.Second * 2)
		if i > 10 {
			atomic.AddInt64(&countGo, -1)
			break
		}
	}
	fmt.Print("g_end ")
}

// округление до любого кратного числа
func work26() {
	k := 2
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 100}
	for _, d := range data {
		fmt.Print(d, " ")
	}

	fmt.Println()

	for _, d := range data {
		d += k / 2
		div := d % k
		res := d - div
		fmt.Print(res, " ")
	}

}

//--------------------------------------------
func work25() {
	wrapLinkCDNDomain := `https://cdnspace.io,https://v2.cdnspace.net`
	wrapLinkCDNPrefix := `https://adplat.sfo2.cdn.digitaloceanspaces.com,https://apimg.fra1.cdn.digitaloceanspaces.com`

	//link := `111`
	link := `https://adplat.sfo2.cdn.digitaloceanspaces.com/123/test`
	//link := `https://apimg.fra1.cdn.digitaloceanspaces.com/456/new`

	//domain := getWrapLinkCDNs(wrapLinkCDNDomain)
	//prefix := getWrapLinkCDNs(wrapLinkCDNPrefix)

	link = wrapImageLinkCDN(link, getWrapLinkCDNs(wrapLinkCDNDomain), getWrapLinkCDNs(wrapLinkCDNPrefix))

	fmt.Println(link)

}

// см.: binder/pkg/selfserve/get_bids
func wrapImageLinkCDN(link string, domain, prefix []string) string {
	if len(domain) == 0 || len(prefix) == 0 || len(domain) != len(prefix) {
		return link
	}

	for i, pref := range prefix {
		if strings.HasPrefix(link, pref) {
			link = domain[i] + link[len(pref):]
			return link
		}
	}

	return link
}

func getWrapLinkCDNs(link string) []string {
	linkCDNs := []string{}
	if link == "" {
		return nil
	}
	for _, linkCDN := range strings.Split(link, ",") {
		linkCDN = strings.TrimSpace(linkCDN)
		if linkCDN == "" {
			continue
		}
		linkCDNs = append(linkCDNs, linkCDN)
	}
	return linkCDNs
}

func work24() {

	uniqueCapIDData := []byte("109.194.11.58")
	uniqueCapID := fmt.Sprintf("%x", sha1.Sum(uniqueCapIDData))

	fmt.Println(uniqueCapID)
}

func work23() {
	//resp := `
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<result>
	//	<link bid="0.000004" url="http://xml.hotmaracas.com/click?i=e5L6KneAjGk_0" pixel="http://xml.hotmaracas.com/pixel?i=e5L6KneAjGk_0" caption="title of push"/>
	//	</result>
	//	`
	resp := `
        <link bid="0.000006" url="http://xml.hotmaracas.com/click?i=e5L6KneAjGk_6" pixel="http://xml.hotmaracas.com/pixel?i=e5L6KneAjGk_6" caption="title of push6"/>
		<link bid="0.000004" url="http://xml.hotmaracas.com/click?i=e5L6KneAjGk_0" pixel="http://xml.hotmaracas.com/pixel?i=e5L6KneAjGk_0" caption="title of push"/>
		<link bid="0.000005" url="http://xml.hotmaracas.com/click?i=e5L6KneAjGk_5" pixel="http://xml.hotmaracas.com/pixel?i=e5L6KneAjGk_5" caption="title of push5"/>
		`
	body := []byte(resp)

	document := etree.NewDocument()
	err := document.ReadFromBytes(body)
	if err != nil {
		fmt.Println("error etree: ", err.Error())
	}

	var elements []*etree.Element
	var root *etree.Element

	root = document.Root()

	elements = append(elements, root)

	for _, item := range elements {
		fmt.Println("bid     = ", item.SelectAttr("bid").Value)
		//elem := item.SelectElement("bid")
		//fmt.Println(elem.Text())
	}

	fmt.Println()
	//fmt.Println(elements)

	fmt.Println("url     = ", root.SelectAttr("url").Value)
	fmt.Println("bid     = ", root.SelectAttr("bid").Value)
	fmt.Println("pixel   = ", root.SelectAttr("pixel").Value)
	fmt.Println("caption = ", root.SelectAttr("caption").Value)

}

//======================================================
// бинарный поиск - test
func work22(b []int, n int) int {
	return sort.SearchInts(b, n)
}

func work22_my(b []int, n int) int {
	res := 0
	for i, v := range b {
		if v >= n {
			res = i
			break
		}
	}
	return res
}

func work21() {
	var count int = -3
	address := []string{"127.0.0.1:2100", "1.2.3.4:3000", "10.20.30.40:4000", "100.200.230.240:5000"}
	lenAddr := len(address)
	tic := time.NewTicker(time.Second * 1)
	for {
		dif := count % lenAddr
		if dif < 0 {
			dif *= -1
		}
		fmt.Println(address[dif])
		fmt.Println(dif)
		count++
		<-tic.C
	}
}

// бинарный поиск
func work20() {
	//b := []int{5, 80, 20, 50, 40, 30, 60, 70, 10}
	//sort.Ints(b)
	//fmt.Println(sort.SearchInts(b, 50))

	text := `0@a
             2@b
             3@c
             4@dd
             5@ee
             6@ff
             7@gg
             8@hh
             9@ii
            10@jj`

	weights, croppedLines := getMass(text)
	fmt.Println(weights, " len =", len(weights))
	fmt.Println(croppedLines, " len =", len(croppedLines))

	// random
	n := 12

	fmt.Println("nn =", weights[len(weights)-1])

	idx := sort.SearchInts(weights, n)
	fmt.Println("ind =", idx)
	fmt.Println("res =", croppedLines[idx])

}

func getMass(text string) ([]int, []string) {
	var summ int

	lines := strings.Split(text, "\n")

	weights := make([]int, 0, len(lines))
	croppedLines := make([]string, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		p := strings.Index(line, "@")
		// если @ не найден, либо находится в первой или последней позиции строки - ошибка
		if p == -1 || p == 0 || p == len(line) {
			continue
		}
		weight, err := strconv.Atoi(line[0:p])
		if err != nil || weight == 0 {
			continue
		}
		summ += weight
		weights = append(weights, summ)
		croppedLines = append(croppedLines, line[p+1:])
	}

	return weights, croppedLines
}

func work19() {
	d := map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1}
	n := map[int]int{1: 2, 3: 2, 5: 2, 7: 2}
	up, nov, del := 0, 0, 0

	fmt.Println("d =", d)
	fmt.Println("n =", n)

	for k, v := range n {
		if _, ok := d[k]; !ok {
			d[k] = v
			nov++
			continue
		}
		d[k] = v
		up++
	}
	fmt.Println("----------------------------------------")
	fmt.Println("d =", d)

	for k := range d {
		if _, ok := n[k]; !ok {
			delete(d, k)
			del++
		}
	}
	fmt.Println("----------------------------------------")
	fmt.Println("d =", d)
	fmt.Println("----------------------------------------")
	fmt.Println("Обнавлено:", up, " Добавлено новых: ", nov, " Удалено:", del, " Всего:", up+nov)

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
	//var t = []byte{123, 34, 105, 100, 34, 58, 34, 84, 69, 83, 84, 34, 44, 34, 104, 111, 117, 114, 34, 58, 55, 55, 55, 44, 34, 99, 111, 117, 110, 116, 114, 121, 34, 58, 34, 82, 117, 115, 34, 44, 34, 114, 95, 99, 108, 105, 99, 107, 34, 58, 102, 97, 108, 115, 101, 44, 34, 114, 95, 105, 109, 112, 34, 58, 102, 97, 108, 115, 101, 44, 34, 114, 95, 118, 105, 101, 119, 34, 58, 102, 97, 108, 115, 101, 44, 34, 97, 100, 118, 101, 114, 116, 105, 115, 101, 114, 95, 99, 97, 115, 104, 98, 97, 99, 107, 34, 58, 48, 125}
	//fmt.Println(string(t))
	//
	//text := "123 34 105 100 34 58 34 84 69 83 84 34 44 34 104 111 117 114 34 58 55 55 55 44 34 99 111 117 110 116 114 121 34 58 34 82 117 115 34 44 34 114 95 99 108 105 99 107 34 58 102 97 108 115 101 44 34 114 95 105 109 112 34 58 102 97 108 115 101 44 34 114 95 118 105 101 119 34 58 102 97 108 115 101 44 34 97 100 118 101 114 116 105 115 101 114 95 99 97 115 104 98 97 99 107 34 58 48 125"
	//text1 := "123 34 100 97 116 101 34 58 34 84 69 83 84 34 44 34 119 105 100 103 101 116 95 105 100 34 58 49 44 34 99 111 117 110 116 114 121 34 58 34 82 117 115 34 125"
	//fmt.Println(encodeText(text1))
	//
	//text2 := "123 34 100 97 116 101 34 58 34 84 69 83 84 34 44 34 119 105 100 103 101 116 95 105 100 34 58 49 44 34 99 111 117 110 116 114 121 34 58 34 82 117 115 34 44 34 105 115 95 112 117 115 104 95 115 117 98 115 99 114 105 112 116 105 111 110 34 58 102 97 108 115 101 44 34 105 115 95 112 117 115 104 95 117 110 115 117 98 115 99 114 105 112 116 105 111 110 34 58 102 97 108 115 101 125"
	//fmt.Println(encodeText(text2))

}

//func encodeText(text string) string {
//	ss := strings.Split(text, " ")
//	res := ""
//	for _, s := range ss {
//		i, _ := strconv.Atoi(s)
//		res += string(i)
//	}
//	return res
//}

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
