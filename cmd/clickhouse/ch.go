// https://habr.com/ru/post/344312/
// https://github.com/yuin/gopher-lua

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	lua "github.com/yuin/gopher-lua"
	"test/cmd/clickhouse/clickhouse"
	"test/cmd/clickhouse/log"
	"test/cmd/clickhouse/mat"
	"test/cmd/clickhouse/options"
)

type config struct {
	ChLogin        string
	ChPassword     string
	ChAddress      string
	ChPort         int
	ChDatabaseName string
	ChSSLMode      string
	ChSSLCertPath  string
}

// параметры подключения к базе данных
func Params() string {
	cfg := config{
		ChLogin:        "default",
		ChPassword:     "",
		ChAddress:      "127.0.0.1",
		ChPort:         9000,
		ChDatabaseName: "schemas",
		ChSSLMode:      "disable",
		ChSSLCertPath:  "",
	}

	chConnString := fmt.Sprintf("clickhouse://%s:%s@%s:%d/%s?sslmode=%s&sslrootcert=%s",
		cfg.ChLogin,
		cfg.ChPassword,
		cfg.ChAddress,
		cfg.ChPort,
		cfg.ChDatabaseName,
		cfg.ChSSLMode,
		cfg.ChSSLCertPath,
	)

	return chConnString
}

// обработка ошибки
func chk(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

var (
	schemaAll = `
SELECT  count(*)
FROM platform_raw.tch
WHERE link_id = 10 AND
      date BETWEEN '2021-01-01' AND '2021-03-01'
`

	schemaTimelog = `
SELECT count(*)
FROM platform_raw.tch
WHERE link_id = 10 AND
      date BETWEEN '2021-01-01' AND '2021-03-01' AND
      minus(time_interval_2, time_interval_1) >= 10
`

	schemaIOSMacos = `
SELECT  count(*)
FROM platform_raw.tch
WHERE link_id = 10 AND
      date BETWEEN '2021-01-01' AND '2021-03-01' AND
      platform_name IN ('iPadOS', 'macOS', 'iOS', 'Mac OS X')
`
	schemaDesktop = `
SELECT  count(*)
FROM platform_raw.tch
WHERE link_id = 10 AND
      date BETWEEN '2021-01-01' AND '2021-03-01' AND
      is_mobile = 0
`

	schemaHeadless = `
SELECT  count(*)
FROM platform_raw.tch
WHERE link_id = 10 AND
      date BETWEEN '2021-01-01' AND '2021-03-01' AND
      is_headless = 1
`
	schemaIframe = `
SELECT  count(*)
FROM platform_raw.tch
WHERE link_id = 10 AND
      date BETWEEN '2021-01-01' AND '2021-03-01' AND
      is_iframe = 1
`

	schemaCaptchaReaction = `
SELECT  count(*)
FROM platform_raw.tch
WHERE link_id = 10 AND
      date BETWEEN '2021-01-01' AND '2021-03-01' AND
      captcha_reaction = 1 AND captcha_solved = 0
`

	schema = `
SELECT time_interval_1,
       time_interval_2,
       minus(time_interval_2, time_interval_1) AS m
    FROM platform_raw.tch
    WHERE link_id = 10
`
)

func main() {

	startLua()

	//mainClickHouse()
}

func startLua() {
	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("clickhouse", clickhouse.Loader)
	L.PreloadModule("log", log.Loader)
	L.PreloadModule("options", options.Loader)
	L.PreloadModule("mat", mat.Loader)
	err := L.DoFile("cmd/clickhouse/rules.lua")
	chk(err)
}

func mainClickHouse() {
	var ch *sql.DB
	var err error

	ch, err = sql.Open("clickhouse", Params())
	if err == nil {
		err = ch.Ping() // отправляем запрос к базе данных - проверка подключения к базе данных
		if err == nil {
			fmt.Println("connecting to clickhouse...")
		} else {
			fmt.Println("no ping to clickhouse...")
		}
	}
	chk(err)
	defer func() {
		err = ch.Close() // закрытие базы данных
		chk(err)
	}()

	//data, err := readDB(ch, schema)
	//chk(err)
	//printData(data)
	//fmt.Println("rows =", len(data))

	var countTimelog, countAll int
	err = ch.QueryRow(schemaAll).Scan(&countAll)
	chk(err)
	fmt.Println("countAll =", countAll)

	err = ch.QueryRow(schemaTimelog).Scan(&countTimelog)
	chk(err)
	res := statusTraffRules(countTimelog, countAll, 20.0)
	fmt.Printf("countTimelog = %d   resTimelog  = %.2f\n", countTimelog, res)

	var countIOSMacos, countDesktop int
	err = ch.QueryRow(schemaIOSMacos).Scan(&countIOSMacos)
	chk(err)
	resIOSMacos := statusTraffProc(countIOSMacos, countAll)
	fmt.Printf("countIOSMacos = %d   resIOSMacos = %.2f\n", countIOSMacos, resIOSMacos)

	err = ch.QueryRow(schemaDesktop).Scan(&countDesktop)
	chk(err)
	res = statusTraffRules(countDesktop, countAll, 60.0)
	fmt.Printf("countDesktop = %d   resDesktop  = %.2f\n", countDesktop, res)

	var countHeadless, countIframe int
	err = ch.QueryRow(schemaHeadless).Scan(&countHeadless)
	chk(err)
	res = statusTraffRulesProc(countIOSMacos, countAll, 5.0)
	fmt.Printf("countHeadless = %d   resHeadless = %.2f\n", countHeadless, res)

	err = ch.QueryRow(schemaIframe).Scan(&countIframe)
	chk(err)
	res = statusTraffRulesProc(countIframe, countAll, 1.0)
	fmt.Printf("countIframe   = %d   resIframe   = %.2f\n", countIframe, res)

	var countCaptchaReaction int
	err = ch.QueryRow(schemaCaptchaReaction).Scan(&countCaptchaReaction)
	chk(err)
	res = statusTraffProc(countCaptchaReaction, countAll)
	fmt.Printf("countCaptchaReaction = %d   resCaptchaReaction = %.2f\n", countCaptchaReaction, res)
}

// Условие ограниченное четко процентом
func statusTraffRulesProc(data, all int, condition float64) float64 {
	const proc = 100
	res := float64(data) * proc / float64(all)
	if res < condition {
		return 0
	}
	return res
}

// Расчет процента
func statusTraffProc(data, all int) float64 {
	const proc = 100
	return float64(data) * proc / float64(all)
}

// Rules
func statusTraffRules(data, all int, condition float64) float64 {
	const proc = 100
	bed := float64(data) * proc / float64(all)
	if bed < condition {
		return 0
	}
	return (bed - condition) * proc / (proc - condition)
}

type dataTimeInterval struct {
	timeInterval1 int
	timeInterval2 int
	m             int
}

// чтение данных из БД
func readDB(ch *sql.DB, query string) ([]dataTimeInterval, error) {
	rows, err := ch.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := make([]dataTimeInterval, 0)

	for rows.Next() {
		d := dataTimeInterval{}
		err = rows.Scan(&d.timeInterval1, &d.timeInterval2, &d.m)
		if err != nil {
			fmt.Println("error read rows", err)
			continue
		}
		data = append(data, d)
	}
	return data, nil
}

// печать данных
func printData(data []dataTimeInterval) {
	line := "--------------------"
	fmt.Println(line)
	fmt.Printf("%2s  %s   %s \n", "Time1", "Time2", " m ")
	fmt.Println(line)
	for _, d := range data {
		fmt.Printf("%4d   %4d  %4d \n", d.timeInterval1, d.timeInterval2, d.m)
	}
	fmt.Println(line)
}
