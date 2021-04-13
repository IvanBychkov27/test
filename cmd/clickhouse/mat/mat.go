package mat

import (
	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	var exports = map[string]lua.LGFunction{
		"sum": sum,
	}
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)
	return 1
}

func sum(L *lua.LState) int {
	a := L.ToInt(1)
	b := L.ToInt(2)

	res := a + b

	L.Push(lua.LNumber(res))
	return 1
}
