package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/christian-gama/radio-spy/options"
	"github.com/christian-gama/radio-spy/radio"
)

type RadioData struct {
	Name       string
	Listeners  uint32
	Percentage string
}

func Home(radios []*radio.Radio, options *options.Options) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			log.Panicf("Não foi possível ler o arquivo index.gohtml")
		}

		var radioData []*RadioData
		radios, totalListeners := radio.PopulateRadio(radios)

		for _, radio := range radios {
			var percentage float32 = float32((radio.GetListeners() * 100)) / float32(totalListeners)

			name := radio.GetName()
			listeners := radio.GetListeners()
			if listeners == 0 {
				percentage = 0
			}

			if !options.ShowQuantity {
				listeners = 0
			}

			radioData = append(radioData, &RadioData{
				Name:       name,
				Listeners:  listeners,
				Percentage: fmt.Sprintf("%.2f%%", percentage),
			})
		}
		data := struct {
			Radios         []*RadioData
			Now            string
			TotalListeners uint32
		}{
			Radios: radioData,
			Now:    time.Now().Format("02/01/2006 15:04:05"),
		}

		t.Execute(w, data)
	}
}

func Server(radios []*radio.Radio, options *options.Options) {
	http.Handle("/", Home(radios, options))

	ClearScreen()
	log.Printf("Servidor iniciado na porta %d", options.Port)
	log.Printf("Acesse http://localhost:%d para ver o resultado\n", options.Port)
	log.Println("Pressione CTRL+C para encerrar o servidor")

	err := http.ListenAndServe(fmt.Sprintf(":%d", options.Port), nil)
	if err != nil {
		log.Panicf("Não foi possível iniciar o servidor na porta %d", options.Port)
	}
}
