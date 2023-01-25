package main

import (
	"fmt"
	"time"

	"github.com/christian-gama/radio-spy/finder"
	"github.com/christian-gama/radio-spy/http"
	"github.com/christian-gama/radio-spy/radio"
)

func main() {
	fmt.Println("Buscando ouvintes, aguarde...")
	result := Spy()
	// clear the screen
	fmt.Print("\033[H\033[2J")
	fmt.Println(result)
}

func Spy() string {
	radios := []*radio.Radio{
		radio.NewRadio("Band FM", "https://hts02.kshost.com.br:12688/index.html?sid=1", `(?i)with (\d+) of .* listeners`),
		radio.NewRadio("Clube", "http://hts01.kshost.com.br:9994/index.html?sid=1", `(?i)with (\d+) of .* listeners`),
		radio.NewRadio("Up", "https://audio8.cmaudioevideo.com/proxy/radioupc", `(?i)listeners \(current\):(\d+)`),
	}

	output := fmt.Sprintf("Quantidade de ouvintes por rádio (%v)\n", time.Now().Format("02/01/2006 15:04:05"))
	var totalListeners uint32 = 0
	radiosMap := make(map[string]uint32)

	for _, radio := range radios {
		res := http.GETRequestTo(radio.GetUrl())
		doc := finder.Document(res)
		listeners := finder.FindListeners(radio.GetListenersPattern(), doc.Text())
		totalListeners += listeners

		radiosMap[radio.GetName()] = listeners
	}

	for radio, listeners := range radiosMap {
		var percentage float32 = float32((listeners * 100) / totalListeners)
		output += fmt.Sprintf("%s: %d (%.2f%%)\n", radio, listeners, percentage)
	}

	return output
}
