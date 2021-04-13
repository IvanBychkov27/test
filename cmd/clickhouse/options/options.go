package options

import (
	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	var exports = map[string]lua.LGFunction{
		"get": get,
	}
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)
	return 1
}

func get(L *lua.LState) int {
	data := 10
	L.Push(lua.LNumber(data))
	return 1
}
