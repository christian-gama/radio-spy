package main

import (
	"fmt"
	"time"

	"github.com/christian-gama/radio-spy/options"
	"github.com/christian-gama/radio-spy/radio"
)

func Executable(radios []*radio.Radio, options *options.Options) {
	fmt.Println("Buscando ouvintes, aguarde...")

	fmt.Println(Output(radios, options))
	ClearScreen()
	fmt.Print("Pressione ENTER para repetir a busca ou CTRL+C para sair...")
	fmt.Scanln()
	ClearScreen()
	Executable(radios, options)
}

func Output(radios []*radio.Radio, options *options.Options) string {
	output := fmt.Sprintf("Quantidade de ouvintes (em porcentagem) por rádio - %v\n", time.Now().Format("02/01/2006 15:04:05"))

	radios, totalListeners := radio.PopulateRadio(radios)

	for _, radio := range radios {
		var percentage float32 = float32((radio.GetListeners() * 100)) / float32(totalListeners)

		name := radio.GetName()
		listeners := radio.GetListeners()
		if listeners == 0 {
			percentage = 0
		}

		if options.ShowQuantity {
			output += fmt.Sprintf("- %s: %d (%.2f%%)\n", name, listeners, percentage)
		} else {
			output += fmt.Sprintf("- %s: %.2f%%\n", name, percentage)
		}
	}

	return output
}
