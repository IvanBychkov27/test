// Для запуска программы необходимо запустить tarantool
// Запуск tarantool: projects/tarantool-cartridge/tnt-main$ cartridge start
// Содержимое таблицы в tarantool можно проверить запустив терминал тарантула
// Запуск терминала тарантул: tarantoolctl connect user:secret@127.0.0.1:3302 (3303, 3304)
// далее: box.space.mem_cap:select{}
// вызываемые функции находятся в проекте projects/tarantool-cartridge/tnt-main/app/roles
// api.lua - это роутер; storage.lua - это storage (реплики)
// запросы приходят на роутер, который перераспределяет запрос на соответствующий storage (по bucket_id)
// при любых изменениях, чтобы они вступили в действие, необходимо перезагрузить тарантул

package main

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var err error
	var bucket_id, count int

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

	id := "id"

	// таблица mem_cap содержит поля:
	// -- 1 id          (- IP+SSCampID)
	// -- 2 bucket_id   (- поле записавается автоматически как код из id)
	// -- 3 ttl
	// -- 4 time_create (время создания записи)

	nRecords := 100 // сколько сделать записей в таблицу
	ttl := 300      // время хранения записи в базе тарантул в секундах

	// включить/отключить запись данных в таблицу
	rec_ok := false // запись отключена
	//rec_ok := true // запись включена

	timeStart := time.Now()
	if rec_ok {
		fmt.Println("Запись данных включена")
		fmt.Println("Время хранения записей ", ttl, " секунд")
		timeStart = time.Now()
		for i := 0; i < nRecords; i++ {
			id = "id"
			n := rand.Intn(1000)
			id = id + strconv.Itoa(n)
			bucket_id, err = strconv.Atoi(time.Now().Format("150405"))

			data := DataCap{id, bucket_id, ttl, 0}

			err = setCapRecord(conn, nAddr, data)
			if err != nil {
				fmt.Println("ошибка записи: ", err.Error())
				break
			}
			count++
			if count > 999 && count >= (nRecords/10) && (count%(nRecords/10)) == 0 {
				fmt.Println("записано:", count*100/nRecords, "%")
			}
		}
		fmt.Println("Время записи данных: ", time.Now().Sub(timeStart))
	} else {
		fmt.Println("запись данных в таблицу отключена")
	}

	fmt.Println()
	fmt.Println("Поиск данных:")

	id = "id96"

	resp, err := getCapRecord(conn, nAddr, id)
	if err != nil {
		fmt.Println(id, err.Error())
	}
	if len(resp) != 0 {
		fmt.Println("getCapRecord", id, ":", resp)
	} else {
		fmt.Println("getCapRecord", id, " : no found")
	}

	timeStart = time.Now()
	result := getCapAllRecords(conn, nAddr)
	timeInterval_getCapAllRecords := time.Now().Sub(timeStart)
	if len(result) != 0 {
		if len(result) < 20 {
			fmt.Println("getCapAllRecords  :", result)
		}
	} else {
		fmt.Println("getCapAllRecords  : records no found")
	}

	ids := []string{"id0", "id1", "id2", "id3", "id4", "id5"}
	timeStart = time.Now()
	_, err = getCapRecords(conn, nAddr, ids)
	timeInterval_getCapRecords := time.Now().Sub(timeStart)
	if err != nil {
		fmt.Println(id, err.Error())
	}
	//if len(resp) != 0 {
	//	fmt.Println("getCapRecords", ids, ":", resp)
	//} else {
	//	fmt.Println("getCapRecords", ids, ": records no found")
	//}

	timeStart = time.Now()
	t := int(time.Now().Unix()) - 30
	fmt.Println("t = ", t)
	resp, err = getCapRecordsTimeCreate(conn, t)
	timeInterval_getCapRecordsTimeCreate := time.Now().Sub(timeStart)
	if err != nil {
		fmt.Println(id, err.Error())
	}
	if len(resp) != 0 {
		fmt.Println("getCapRecordsTimeCreate", t, ":", len(resp))
	} else {
		fmt.Println("getCapRecordsTimeCreate", t, ": records no found")
	}

	fmt.Println()
	fmt.Println("all records:", len(result))
	fmt.Println("timeInterval_getCapAllRecords:", timeInterval_getCapAllRecords)
	fmt.Println("timeInterval_getCapRecords:", timeInterval_getCapRecords, " (", len(ids), ") records")
	fmt.Println("timeInterval_getCapRecordsTimeCreate:", timeInterval_getCapRecordsTimeCreate, " (", len(resp), ") records")

}

type DataCap struct {
	Id         string
	Bucket_id  int
	Ttl        int
	TimeCreate int
}

// пишет одну запись с инфо (id, bucket_id, ttl) в таблицу mem_cap (tarantool)
func setCapRecord(conn *tarantool.Connection, nAddr string, d DataCap) error {
	var err error
	if nAddr == "1" {
		_, err = conn.Call("set_cap", []interface{}{d.Id, d.Ttl})
	} else {
		_, err = conn.Call("_set_cap", []interface{}{d.Id, d.Ttl})
	}

	if err != nil {
		return err
	}
	return nil
}

// читает одну запись с id из таблицы mem_cap (tarantool)
func getCapRecord(conn *tarantool.Connection, nAddr, id string) (map[string]struct{}, error) {
	var capRecords []DataCap
	var err error

	if nAddr == "1" {
		err = conn.CallTyped("get_cap", []interface{}{id}, &capRecords)
	} else {
		err = conn.CallTyped("_get_cap", []interface{}{id}, &capRecords)
	}
	if err != nil {
		return nil, err
	}

	capData := make(map[string]struct{})
	for _, c := range capRecords {
		if c.Id != "" {
			capData[c.Id] = struct{}{}
		}
	}

	return capData, nil
}

// читает несколько записей с []id из таблицы mem_cap (tarantool)
func getCapRecords(conn *tarantool.Connection, nAddr string, ids []string) (map[string]struct{}, error) {
	var capRecords []DataCap
	var err error

	if nAddr == "1" {
		err = conn.CallTyped("getList_cap", []interface{}{ids}, &capRecords)
	} else {
		err = conn.CallTyped("_getList_cap", []interface{}{ids}, &capRecords)
	}
	if err != nil {
		return nil, err
	}

	capData := make(map[string]struct{})
	for _, c := range capRecords {
		if c.Id != "" {
			capData[c.Id] = struct{}{}
		}
	}

	return capData, nil
}

// читает все записи из таблицы mem_cap (tarantool)
func getCapAllRecords(conn *tarantool.Connection, nAddr string) map[string]struct{} {
	capData := make(map[string]struct{})
	var capRecords []DataCap

	var err error
	if nAddr == "1" {
		err = conn.CallTyped("getAll_cap", []interface{}{}, &capRecords)
	} else {
		err = conn.CallTyped("_getAll_cap", []interface{}{}, &capRecords)
	}
	if err != nil {
		fmt.Println("error: ", err.Error())
		return nil
	}
	if capRecords == nil {
		return nil
	}

	for _, c := range capRecords {
		if c.Id != "" {
			capData[c.Id] = struct{}{}
		}
	}

	return capData
}

// читает несколько записей из таблицы mem_cap начиная с времени timeCreate до текущего времени
func getCapRecordsTimeCreate(conn *tarantool.Connection, t int) (map[string]struct{}, error) {
	var capRecords []DataCap
	var err error

	err = conn.CallTyped("getListTimeCreate_cap", []interface{}{t}, &capRecords)
	if err != nil {
		return nil, err
	}

	capData := make(map[string]struct{})
	for _, c := range capRecords {
		if c.Id != "" {
			capData[c.Id] = struct{}{}
		}
	}

	return capData, nil
}
