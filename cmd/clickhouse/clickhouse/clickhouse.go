package clickhouse

import (
	"database/sql"
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	var exports = map[string]lua.LGFunction{
		"query": query,
	}
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)
	return 1
}

func query(L *lua.LState) int {
	resL := L.ToString(1)
	L.Push(lua.LString(resL))
	schemaAll := resL

	var ch *sql.DB
	var err error

	ch, err = sql.Open("clickhouse", params())
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

	var countAll int
	err = ch.QueryRow(schemaAll).Scan(&countAll)

	L.Push(lua.LNumber(countAll))

	if err != nil {
		L.Push(lua.LString(err.Error()))
	} else {
		L.Push(lua.LNil)
	}
	return 2
}

// обработка ошибки
func chk(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

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
func params() string {
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
