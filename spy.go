package main

import (
	"fmt"
	"time"

	"github.com/christian-gama/radio-spy/finder"
	"github.com/christian-gama/radio-spy/http"
	"github.com/christian-gama/radio-spy/radio"
)

// Spy is the main function of the application
func Spy() string {
	radios := CreateRadios(Settings)
	options := CreateOptions(Settings)

	output := fmt.Sprintf("Quantidade de ouvintes (em porcentagem) por rádio - %v\n", time.Now().Format("02/01/2006 15:04:05"))
	var totalListeners uint32 = 0

	for _, radio := range radios {
		listeners := SetListenersFromURL(radio)
		totalListeners += listeners
	}

	for _, radio := range radios {
		var percentage float32 = float32((radio.GetListeners() * 100)) / float32(totalListeners)
		if options.ShowQuantity {
			output += fmt.Sprintf("- %s: %d (%.2f%%)\n", radio.GetName(), radio.GetListeners(), percentage)
		} else {
			output += fmt.Sprintf("- %s: %.2f%%\n", radio.GetName(), percentage)
		}
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
