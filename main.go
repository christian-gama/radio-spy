package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/christian-gama/radio-spy/finder"
	"github.com/christian-gama/radio-spy/http"
	"github.com/christian-gama/radio-spy/radio"
)

// clear is a map with a function for each supported platform
var clear map[string]func()

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

// ClearScreen clear the terminal screen
func ClearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
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

// Spy is the main function of the application
func Spy() string {
	radios := []*radio.Radio{
		radio.NewRadio("Band FM", "https://hts02.kshost.com.br:12688/index.html?sid=1", `(?i)with (\d+) of .* listeners`),
		radio.NewRadio("Clube", "http://hts01.kshost.com.br:9994/index.html?sid=1", `(?i)with (\d+) of .* listeners`),
		radio.NewRadio("Up", "https://audio8.cmaudioevideo.com/proxy/radioupc", `(?i)listeners \(current\):(\d+)`),
	}

	output := fmt.Sprintf("Quantidade de ouvintes por rádio (%v)\n", time.Now().Format("02/01/2006 15:04:05"))
	var totalListeners uint32 = 0

	for _, radio := range radios {
		listeners := SetListenersFromURL(radio)
		totalListeners += listeners
	}

	for _, radio := range radios {
		var percentage float32 = float32((radio.GetListeners() * 100)) / float32(totalListeners)
		output += fmt.Sprintf("%s: %d (%.2f%%)\n", radio.GetName(), radio.GetListeners(), percentage)
	}

	return output
}

// SetListenersFromURL get the listeners from the radio url and set the listeners value
func SetListenersFromURL(radio *radio.Radio) uint32 {
	res := http.MustGet(radio.GetUrl())
	doc := finder.Document(res)
	listeners := finder.FindListeners(radio.GetListenersPattern(), doc.Text())
	radio.SetListeners(listeners)

	return listeners
}
