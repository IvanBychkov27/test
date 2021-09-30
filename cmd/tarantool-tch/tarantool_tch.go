package main

import (
	"fmt"
	"github.com/hashicorp/go-uuid"
	"github.com/tarantool/go-tarantool"
	"math/rand"
	"strconv"
	"time"
)

// таблица tch содержит поля:
//-- 1 id            (string)
//-- 2 bucket_id     (unsigned)
//-- 3 created_at    (unsigned)   - дата создания объекта статы
//-- 4 time_on_page  (unsigned)   - время нахождения пользователя на странице в секундах (кол-во вызовов Ping, которые дергаются с фронта каждую секунду)
//-- 5 captcha_reaction (bool)    - true, если была реакция на капчу
//-- 6 captcha_solved   (bool)    - true, если была капча пройдена
//-- 7 mousemove        (bool)    - true, если были движения мышкой
//-- 8 stat_basic       (string)  - объект статы который пишется сразу (на основе UA, IP)
//-- 9 stat_full        (string)  - объект статы, который приходит с JS
// -- 10 has_pixel       - true, если сработал пиксель
// -- 11 has_js_pixel    - true, если сработал JS пиксель
// -- 12 stat_fingerprint - объект статы, который приходит с TCH-POSF (пишется на основе TCP/IP)
// -- 13 ip_port          - string ip:port

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	conn, _ := connectTarantool("127.0.0.1:3401")
	defer conn.Close()

	uin, _ := setUIN()
	fmt.Println("uin =", uin)

	var stat []byte
	var err error
	ipPort := "127.0.0.1:2001"

	stat = []byte("basic_2")
	err = setRecord(conn, uin, ipPort, stat)
	if err != nil {
		fmt.Println("error setRecord:", err.Error())
		return
	}

	//uin = "ad478665-1092-0fe4-3542-8a1476c33040"
	//ipPort = uin

	//stat = []byte("full_002")
	//err = updateRecordFull(conn, uin, stat)
	//if err != nil {
	//	fmt.Println("error updateRecord:", err.Error())
	//	return
	//}

	//stat = []byte("fingerprint_1")
	//err = updateRecordFingerprint(conn, ipPort, stat)
	//if err != nil {
	//	fmt.Println("error updateRecord:", err.Error())
	//	return
	//}

	//fmt.Println("В таблицу tch добавлена запись с uin =", uin)
}

// Пишет одну запись в таблицу tch (tarantool-tch)
func setRecord(conn *tarantool.Connection, uin, ipPort string, stat []byte) error {
	_, err := conn.Call("store_basic", []interface{}{uin, ipPort, stat})
	return err
}

// Обнавляем запись в таблицу tch по uin (tarantool-tch)
func updateRecordFull(conn *tarantool.Connection, uin string, stat []byte) error {
	_, err := conn.Call("store_full", []interface{}{uin, stat})
	return err
}

// Обнавляем запись в таблицу tch по uin (tarantool-tch)
func updateRecordFingerprint(conn *tarantool.Connection, uin string, stat []byte) error {
	_, err := conn.Call("store_fingerprint", []interface{}{uin, stat})
	return err
}

func connectTarantool(addr string) (conn *tarantool.Connection, err error) {
	conn, err = tarantool.Connect(addr, tarantool.Opts{
		User: "user",
		Pass: "secret",
	})
	if err != nil {
		fmt.Println("tarantool no connect")
		return
	} else {
		fmt.Println("tarantool connection... addr ", addr)
	}
	return
}

func setUIN() (uin string, err error) {
	uin, err = uuid.GenerateUUID()
	if err != nil {
		fmt.Println("error generate uin:", err.Error())
		uin = strconv.Itoa(rand.Intn(1e9))
	}
	return
}

func chk(err error) {
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
}

//
//type tarantoolResponse struct {
//	Code  int
//	Error string
//}
//
//// Пишет одну запись в таблицу tch (tarantool-tch)
//func setRecord_1(conn *tarantool.Connection, uin string, stat []byte) error {
//	resp := &tarantoolResponse{}
//
//	err := conn.CallTyped("store_basic", []interface{}{uin, stat}, resp)
//	if err != nil {
//		return fmt.Errorf("error call '%s', %w", "store_basic", err)
//	}
//	if resp.Error != "" {
//		return fmt.Errorf("response error, %s", resp.Error)
//	}
//	if resp.Code != 201 {
//		return fmt.Errorf("unexpected response code %d", resp.Code)
//	}
//	return nil
//}
