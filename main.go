package main

import (
	"fmt"
	"runtime"
)

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

// ClearScreen clear the terminal screen
func ClearScreen() {
	value, ok := Clear[runtime.GOOS]
	if ok {
		value()
	}
}
