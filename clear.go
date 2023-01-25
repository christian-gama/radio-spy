package main

import (
	"os"
	"os/exec"
)

// Clear is a map with a function for each supported platform
var Clear map[string]func()

func init() {
	Clear = make(map[string]func())
	Clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	Clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
