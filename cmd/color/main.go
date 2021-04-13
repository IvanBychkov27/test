package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

func main() {

	fmt.Println("Hello!")
	color.Red("Red HELLO!")

	color.Blue("Blue HELLO!")
	color.Green("Green HELLO!")
	color.Yellow("Yellow HELLO!")

	//minion := color.New(color.FgBlack).Add(color.BgYellow).Add(color.Bold)
	//minion.Println("Minion says: banana!!!!!!")

	//m := minion.PrintlnFunc()
	//m("I want another banana!!!!!")

	//slantedRed := color.New(color.FgRed, color.BgWhite, color.Italic).SprintFunc()
	//fmt.Println("I’ve made a huge", slantedRed("mistake"))

	fmt.Print("New text")
	color.Red(" RED")

	color.Set(color.FgBlue)
	defer color.Unset()
	fmt.Println("New text01")
	fmt.Println("New text02")
	log.Print("Log Hello!")

	c := color.New(color.FgCyan)
	c.Println("Prints cyan text")

	c.DisableColor()
	c.Println("This is printed without any color")

	c.EnableColor()
	c.Println("This prints again cyan...")

}
