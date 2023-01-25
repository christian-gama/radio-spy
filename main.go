package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

var Settings SettingsYML

func init() {
	Settings = ReadYML()

}

func main() {
	fmt.Println("Buscando ouvintes, aguarde...")
	result := Spy()
	ClearScreen()

	fmt.Println(result)
	fmt.Print("Pressione ENTER para repetir a busca ou CTRL+C para sair...")
	fmt.Scanln()
	ClearScreen()
	main()
}

// clear is a map with a function for each supported platform
var clear map[string]func()

// ClearScreen clear the terminal screen
func ClearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	}
}
