package log

import (
	lua "github.com/yuin/gopher-lua"
	"go.uber.org/zap"
	"log"
	"os"
)

func Loader(L *lua.LState) int {
	var exports = map[string]lua.LGFunction{
		"info": info,
	}
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)
	return 1
}

func info(L *lua.LState) int {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Printf("error create logger, %v", err)
		os.Exit(1)
	}

	s := L.ToString(1)
	logger.Info("logger", zap.String("result", s))
	return 0
}
