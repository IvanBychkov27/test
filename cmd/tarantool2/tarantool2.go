package main

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	"math/rand"
	"strconv"
	"time"
)

type DataImp struct {
	Id        string
	Bucket_id int
	Ttl       int
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var err error
	var bucket_id int

	nAddr := "1" //  1 - router; 2, 3, 4 - storage
	addr := "127.0.0.1:330" + nAddr
	conn, err := tarantool.Connect(addr, tarantool.Opts{
		User: "user",
		Pass: "secret",
	})

	if err != nil {
		fmt.Println("tarantool no connect")
		return
	} else {
		fmt.Println("tarantool connection... addr ", addr)
	}
	defer conn.Close()

	// таблица mem_cap содержит поля:
	// -- 1 id          (- sessions_id)
	// -- 2 bucket_id   (- поле записавается автоматически как код из id)
	// -- 3 ttl

	id := "id"

	// включить/отключить запись данных в таблицу
	//rec_ok := false // запись отключена
	rec_ok := true // запись включена

	if rec_ok {
		fmt.Println("Запись данных включена")

		n := rand.Intn(100)
		id = id + strconv.Itoa(n)
		bucket_id, err = strconv.Atoi(time.Now().Format("150405"))
		ttl := 120 // время хранения записи в базе тарантул в секундах

		data := DataImp{id, bucket_id, ttl}

		err = setImpRecord(conn, nAddr, data)
		if err != nil {
			fmt.Println("ошибка записи: ", err.Error())
			return
		}
		fmt.Println(" в таблицу  link_sessions_imp (тарантула) записана запись с id: ", id)
		fmt.Println()
	}

	//id = "id98"
	resp, err := getImpRecord(conn, nAddr, id)
	if err != nil {
		fmt.Println(id, err.Error())
	}
	if len(resp) != 0 {
		fmt.Println("getImpRecord", id, ":", resp)
	} else {
		fmt.Println("getImpRecord", id, " : no found")
	}
}

// пишет одну запись с инфо (id, bucket_id, ttl) в таблицу link_sessions_imp (tarantool)
func setImpRecord(conn *tarantool.Connection, nAddr string, d DataImp) error {
	var err error
	if nAddr == "1" {
		_, err = conn.Call("set_link_imp", []interface{}{d.Id, d.Ttl})
	} else {
		_, err = conn.Call("_set_link_imp", []interface{}{d.Id, d.Bucket_id, d.Ttl})
	}

	if err != nil {
		return err
	}
	return nil
}

// читает одну запись с id из таблицы link_sessions_imp (tarantool)
func getImpRecord(conn *tarantool.Connection, nAddr, id string) (map[string]struct{}, error) {
	var impRecords []DataImp
	var err error

	if nAddr == "1" {
		err = conn.CallTyped("get_link_imp", []interface{}{id}, &impRecords)
	} else {
		err = conn.CallTyped("_get_link_imp", []interface{}{id}, &impRecords)
	}
	if err != nil {
		return nil, err
	}

	impData := make(map[string]struct{})
	for _, c := range impRecords {
		if c.Id != "" {
			impData[c.Id] = struct{}{}
		}
	}

	return impData, nil
}
