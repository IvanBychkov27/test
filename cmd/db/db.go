// см новый пакет работы с PostgreSQL https://github.com/jackc/pgx

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

const (
	pageForm = `
<html>
 <head>
    <title>DataBase</title>
    <style>
            td{
                width: 100px;
                height: 32px;
                border: solid 1px silver;
                text-align: center;
            }
        </style>
 </head>
 <body>
    <br>
    <form action="/add" method="POST">
        <input type="submit" value="add" />
	</form>
	<br>
	<form action="/" method="POST">
		<label for="value"><B>SELECT * FROM Products</B></label><br>
		<input type="text" autofocus id="value" name="value" size=100 />
		<input type="submit" value="Send" />
	</form>
`
	endPage = `
 </body>
</html>
`
)

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	queryUser := ""
	err := req.ParseForm()
	if err != nil {
		log.Println(err.Error())
		return
	}
	if len(req.Form) != 0 {
		queryUser = req.Form.Get("value")
	}

	query := "SELECT * FROM Products " + queryUser
	res := `<h4 style="color:blue;"> ` + query + `</h4>`

	log.Println(query)

	if queryUser != "" {
		res += `<h3 style="color:green;">`
		res += db_update(query)
		res += "</h3>"
	}

	_, err = w.Write([]byte(pageForm + res + endPage))
	if err != nil {
		log.Println(err.Error())
	}
}

// Добавление новой записи в базу данных
func addNewRecord(w http.ResponseWriter, req *http.Request) {
	startPage := `
<html>
 <head>
    <title>Add a new record to the database</title>
		<style>
            td{
                width: 100px;
                height: 32px;
                border: solid 1px silver;
                text-align: center;
            }
        </style>
 </head>
<body>
	<br>
	<form action="/" method="POST">
        <input type="submit" value="main" />
	</form>
	<br>
`
	res := "<table><thead> <tr> <th> ID </th> <th> Model </th> <th> Company </th> <th> Price,т.р.</th> </tr></thead>"
	res += `<tbody>
			<form action="/add" method="POST">
			<tr>
            <td>avto</td>
            <td><input autofocus type="text" name="formModel" size=10 /></td>
			<td><input type="text" name="formCompany" size=10 /></td>
			<td ><input type="text" name="formPrice" size=10 /></td></tr>
			    <input type="submit" value="save" />
			</form>
`
	res += "</tbody></table>"

	var formModel, formCompany, formPrice string
	err := req.ParseForm()
	if err != nil {
		log.Println(err.Error())
		return
	}
	if len(req.Form) != 0 {
		formModel = req.Form.Get("formModel")
		formCompany = req.Form.Get("formCompany")
		formPrice = req.Form.Get("formPrice")
	}

	if formModel != "" && formCompany != "" && formPrice != "" {
		price, err := strconv.Atoi(formPrice)
		if err != nil {
			res += "<br> полe Price внесено не корректно "
		} else {
			data := []car{{0, formModel, formCompany, price}}
			addAllRecords(data)
			res += "<br> запись внесена в БД "
		}
	}

	_, err = w.Write([]byte(startPage + res + endPage))
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	fmt.Println("server is listening... 127.0.0.1:8181")
	http.HandleFunc("/", ServeHTTP)
	http.HandleFunc("/add", addNewRecord)
	err := http.ListenAndServe("localhost:8181", nil)
	chk(err)
}

// печать данных на Web страницу - создание строки
func webPrintCar(dataAll []car) string {
	line := "----------------------------------------------<br>"
	res := line
	res += fmt.Sprintf("Table Products  - %d records <br>", len(dataAll)) + line
	res += "<table>"
	res += "<thead> <tr> <th> ID </th> <th> Model </th> <th> Company </th> <th> Price,т.р.</th> </tr></thead>"
	res += "<tbody>"
	for _, data := range dataAll {
		res += "<tr>"
		res += fmt.Sprintf(`<td style="width:40px;">%d</td>  <td>%s</td> <td>%s</td> <td style="text-align:right;">%d</td>`, data.id, data.model, data.company, data.price)
		res += "</tr>"
	}
	res += "</tbody></table>" + line

	return res
}

type config struct {
	PgLogin        string
	PgPassword     string
	PgAddress      string
	PgPort         int
	PgDatabaseName string
	PgSSLMode      string
	PgSSLCertPath  string
}

type car struct {
	id      int
	model   string
	company string
	price   int
}

// основная функция обращения к БД
func db_update(q string) string {
	db, err := sql.Open("postgres", params()) // подключение к базе данных
	if err == nil {
		err = db.Ping() // отправляем запрос к базе данных - проверка подключения к базе данных
		//fmt.Println("connecting to postgres...")
	}
	chk(err)
	defer func() {
		err = db.Close() // закрытие базы данных
		chk(err)
	}()

	//createTable(db) // создание таблицы
	//addRecord(db)    // добавление одной записи в таблицу
	//addAllRecordsInTable(db) // добавление всех записей в таблицу
	//delRecordID(db, 48)

	//var avto = car{id: 49, model: "Kaptur", company: "Renault", price: 1140}
	//updateRecodID(db, avto.id, avto)

	//query := `SELECT * FROM Products WHERE company IN ('Renault', 'Ford') ORDER BY company, price`
	//query := `SELECT * FROM Products WHERE company LIKE 'Ren%' ORDER BY price DESC`
	//query := `SELECT * FROM Products ORDER BY id`
	//query := `SELECT * FROM Products ORDER BY price DESC`
	//query := "SELECT * FROM Products WHERE price > 1000 AND price < 1500"
	//query := "SELECT * FROM Products WHERE id = 5"
	//query := "SELECT * FROM Products WHERE NOT company ='Renault'" // ' - это апостров см. где буква 'э'
	//query := "SELECT * FROM Products WHERE company ='Ford' AND price < 1000"

	query := q
	data, err := readDB(db, query)
	chk(err)
	printCar(data) // печать информации полученной из БД
	return webPrintCar(data)
}

// изменение записи по id
func updateRecodID(db *sql.DB, ID int, newData car) {
	_, err := db.Exec("UPDATE Products SET model = $1, company = $2, price = $3 WHERE id = $4", newData.model, newData.company, newData.price, ID)
	if err != nil {
		fmt.Printf("ошибка обнавления записи с ID = %d \n", ID)
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("обнавлена запись с ID = %d \n", ID)
}

// удаление записи по id
func delRecordID(db *sql.DB, ID int) {
	_, err := db.Exec("DELETE FROM Products WHERE id = $1", ID)
	if err != nil {
		fmt.Printf("ошибка удаления записи с ID = %d \n", ID)
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("из таблицы Products удалена запись с ID = %d \n", ID)
}

// query := `SELECT traffic_type, script, name_script, weight FROM tch_rules`
// создание таблицы
func createTable(db *sql.DB) {
	// query := `SELECT traffic_type, script, name_script, weight FROM tch_rules`

	var schema = `CREATE TABLE widgets (
    id    integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    user_id integer NOT NULL,
    widget_id integer NOT NULL,
    endpoint_id integer NOT NULL,
    token varchar(30) NOT NULL
);`

	_, err := db.Exec(schema) // создание таблицы
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Создана таблица widgets")
}

// добавление одной записи в таблицу
func addRecord(db *sql.DB) {
	res, err := insert(db, "", "", 0)
	chk(err)
	fmt.Printf("Добавлена %d запись в таблицу Products\n", res)
}

//func addAllRecordsInTable(db *sql.DB) {
//	avtoAll := make([]car, 0)
//	avtoAll = []car{
//		//{model: "test", company: "test", price: 100},
//		//{model: "111", company: "111", price: 111},
//		//{model: "", company: "", price: },
//	}
//	addAllRecords(db, avtoAll)
//}

// добавление всех записей из слайса в таблицу
func addAllRecords(data []car) {
	db, err := sql.Open("postgres", params()) // подключение к базе данных
	if err == nil {
		err = db.Ping() // отправляем запрос к базе данных - проверка подключения к базе данных
		//fmt.Println("connecting to postgres...")
	}
	chk(err)
	defer func() {
		err = db.Close() // закрытие базы данных
		chk(err)
	}()

	if len(data) == 0 {
		fmt.Println("новые данные для внесения в таблицу отсутствуют")
		return
	}
	var recordsCount int64
	for _, d := range data {
		res, err := insert(db, d.model, d.company, d.price)
		chk(err)
		recordsCount += res
	}
	fmt.Printf("Добавлено %d записей в таблицу Products\n", recordsCount)
}

// создает новую запись и возвращает кол-во затронутых строк и ошибку
func insert(db *sql.DB, model, company string, price int) (int64, error) {
	res, err := db.Exec("INSERT INTO Products VALUES (default, $1, $2, $3)", model, company, price)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// чтение данных из БД
func readDB(db *sql.DB, query string) ([]car, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	avto := make([]car, 0)

	for rows.Next() {
		p := car{}
		err = rows.Scan(&p.id, &p.model, &p.company, &p.price)
		if err != nil {
			fmt.Println("error read rows", err)
			continue
		}
		avto = append(avto, p)
	}
	return avto, nil
}

// печать данных
func printCar(avtoAll []car) {
	fmt.Printf("Table Products  - %d records\n", len(avtoAll))
	fmt.Println("----------------------------------")
	fmt.Printf("%2s  %-10s %-10s %s \n", "ID", "Model", "Company", "Price")
	fmt.Println("----------------------------------")
	for _, avto := range avtoAll {
		fmt.Printf("%2d  %-10s %-10s %4d \n", avto.id, avto.model, avto.company, avto.price)
	}
	fmt.Println("----------------------------------")
}

// параметры подключения к базе данных
func params() string {
	cfg := config{
		PgLogin:        "dbuser",
		PgPassword:     "secret",
		PgAddress:      "127.0.0.1",
		PgPort:         5432,
		PgDatabaseName: "test",
		PgSSLMode:      "disable",
		PgSSLCertPath:  "",
	}

	pgConnString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&sslrootcert=%s",
		cfg.PgLogin,
		cfg.PgPassword,
		cfg.PgAddress,
		cfg.PgPort,
		cfg.PgDatabaseName,
		cfg.PgSSLMode,
		cfg.PgSSLCertPath,
	)

	return pgConnString
}

// обработка ошибки
func chk(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
