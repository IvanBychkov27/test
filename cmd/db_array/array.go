package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"strings"
)

type workLog struct {
	ID        int
	UserID    int
	ModelName string
	ModelID   int
}

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

	query := `
SELECT id, user_id, model_name, model_id
FROM work_log 
WHERE model_name = ANY($1)
`

	dataStr := "user,token"

	array := strings.Split(dataStr, ",")

	//array := []string{"user", "token"}

	fmt.Println("array:", array)

	ws, err := readFinance(db, query, array)
	if err != nil {
		fmt.Println("error:", err.Error())
	}

	fmt.Println("res len =", len(ws))

}

// чтение данных из БД
func readFinance(db *pgx.Conn, query string, ar []string) ([]workLog, error) {
	rows, err := db.Query(context.Background(), query, ar)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]workLog, 0)
	for rows.Next() {
		w := workLog{}
		err = rows.Scan(
			&w.ID,
			&w.UserID,
			&w.ModelName,
			&w.ModelID,
		)
		if err != nil {
			fmt.Println("error read rows", err)
			continue
		}
		data = append(data, w)
	}
	return data, nil
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
		PgDatabaseName: "finance",
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
