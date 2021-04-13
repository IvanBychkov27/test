// https://habr.com/ru/post/344312/
// https://github.com/yuin/gopher-lua
package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"math/rand"
	"test/cmd/lua/mymodule"
	"time"
)

const (
	strLua1 = `
function summ(a,b)
	return a + b
end
print("strLua summ =", summ(5,7))
print(mult(20, 5))
`
)

func Mult(L *lua.LState) int {
	lv := L.ToInt(1) // get argument
	lv2 := L.ToInt(2)
	L.Push(lua.LNumber(lv * lv2)) // push result
	return 1                      // number of results
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var err error
	L := lua.NewState()
	defer L.Close()

	L.PreloadModule("mymodule", mymodule.Loader)
	err = L.DoFile("cmd/lua/main.lua")
	chk(err)

	//fmt.Println("Res = ", mymodule.Res)

	//L.SetGlobal("mult", L.NewFunction(Mult))
	//err = L.DoString(strLua1)
	//chk(err)

	//err = L.DoFile("cmd/lua/file.lua")
	//chk(err)

}

func chk(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
