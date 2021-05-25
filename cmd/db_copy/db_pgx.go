package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"strconv"
	"time"
)

// массовая загрузка (запись) данных в PostgreSQL
func main() {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, params())
	if err != nil {
		fmt.Println("error connect db: ", err.Error())
		return
	}
	err = db.Ping(ctx)
	if err != nil {
		fmt.Println("error ping postgres db: ", err.Error())
	}
	fmt.Println("connecting to postgres...")

	//writeOk := false
	writeOk := true
	if writeOk {
		dataColl := dataCollectors()
		rows := inputRows(dataColl)

		tableName := pgx.Identifier{"collector01"}
		columnName := []string{"user_id", "key", "time_create", "ip", "ua"}

		dataRows := pgx.CopyFromRows(rows)

		var nRows int64
		nRows, err = db.CopyFrom(ctx, tableName, columnName, dataRows) // массовая загрузка данных в PostgreSQL
		if err != nil {
			fmt.Println("error copy form: ", err.Error())
		}
		fmt.Println("записано строк ", nRows)
	}

	//readOk := false
	readOk := true
	if readOk {
		var data []Collector
		query := `SELECT * FROM collector01`
		data, err = readDB(db, query) // читаем базу данных
		chk(err)
		printCar(data) // печать информации полученной из БД
	}
}

type Collector struct {
	ID         int
	UserID     int
	Key        string
	TimeCreate int
	IP         string
	UA         string
}

// заполняем массив строк
func inputRows(colls []Collector) [][]interface{} {
	rows := make([][]interface{}, 0, len(colls))

	for _, c := range colls {
		input := []interface{}{c.UserID, c.Key, c.TimeCreate, c.IP, c.UA}
		rows = append(rows, input)
	}

	return rows
}

// заполняем массив данных Collector
func dataCollectors() []Collector {
	collectors := make([]Collector, 0, 5)

	for i := 30; i < 35; i++ {
		c := Collector{
			ID:         0,
			UserID:     i,
			Key:        "collectors pgx key",
			TimeCreate: int(time.Now().Unix()),
			IP:         "127.0.0." + strconv.Itoa(i),
			UA:         "collectors pgx ua" + strconv.Itoa(i),
		}

		collectors = append(collectors, c)
	}
	return collectors
}

// создает новую запись и возвращает кол-во затронутых строк и ошибку
func insert(db *pgx.Conn, c Collector) error {
	_, err := db.Exec(
		context.Background(),
		"INSERT INTO collector01 VALUES (default, $1, $2, $3, $4, $5)",
		c.UserID, c.Key, c.TimeCreate, c.IP, c.UA)

	return err
}

// чтение данных из БД
func readDB(db *pgx.Conn, query string) ([]Collector, error) {
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	//defer rows.Close()
	collectors := make([]Collector, 0)

	for rows.Next() {
		c := Collector{}
		err = rows.Scan(&c.ID, &c.UserID, &c.Key, &c.TimeCreate, &c.IP, &c.UA)
		if err != nil {
			fmt.Println("error read rows", err)
			continue
		}
		collectors = append(collectors, c)
	}
	return collectors, nil
}

// печать данных
func printCar(collectors []Collector) {
	fmt.Printf("Table Products  - %d records\n", len(collectors))
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Printf("%4s  %10s %12s %27s %5s %10s \n", "ID", "UserID", "Key", "TimeCreate", "IP", "UA")
	fmt.Println("--------------------------------------------------------------------------------------")
	for _, c := range collectors {
		fmt.Printf("%4d  %8d  %25s %15d %-10s %-10s \n", c.ID, c.UserID, c.Key, c.TimeCreate, c.IP, c.UA)
	}
	fmt.Println("--------------------------------------------------------------------------------------")
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
