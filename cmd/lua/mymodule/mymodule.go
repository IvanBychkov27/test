package mymodule

import (
	"github.com/yuin/gopher-lua"
	"math/rand"
	"time"
)

var exports = map[string]lua.LGFunction{
	"myfunc": myfunc,
}

func myfunc(L *lua.LState) int {
	return 0
}

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.SetField(mod, "nowTime", lua.LString(nowTime()))
	L.SetField(mod, "randNumber", lua.LNumber(randomNumber()))
	L.SetField(mod, "rn1", lua.LNumber(rN1()))
	L.SetField(mod, "rn2", lua.LNumber(rN2()))
	L.Push(mod)
	L.SetGlobal("mult", L.NewFunction(Mult))
	L.SetGlobal("getInLua", L.NewFunction(GetInLua))
	L.SetGlobal("data", L.NewFunction(InLuaData))
	return 1
}

func Mult(L *lua.LState) int {
	lv := L.ToInt(1) // get argument
	lv2 := L.ToInt(2)
	res := lv * lv2
	L.Push(lua.LNumber(res)) // push result
	Res = lv
	return 1 // number of results
}

var Res int

func InLuaData(L *lua.LState) int {
	data := 1000
	L.Push(lua.LNumber(data))
	return 1
}

func GetInLua(L *lua.LState) int {
	res := L.ToInt(1)
	L.Push(lua.LNumber(res))
	Res = res
	return res
}

func nowTime() string {
	t := time.Now()
	strNowTime := "Текущее время: " + t.Format("15:04:05")
	return strNowTime
}

func randomNumber() float64 {
	randNum := 100 + rand.Intn(900)
	return float64(randNum)
}

func rN1() float64 {
	return float64(rand.Intn(1000))
}

func rN2() float64 {
	return float64(rand.Intn(1000))
}
