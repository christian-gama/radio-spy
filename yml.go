package main

// read yml file and return a slice of radios

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/christian-gama/radio-spy/radio"
	"gopkg.in/yaml.v3"
)

// RadioYML is the struct that represents the yml file
type RadioYML struct {
	Radios []Radio `yaml:"radios"`
}

type Radio struct {
	Nome    string `yaml:"nome"`
	Url     string `yaml:"url"`
	Pattern string `yaml:"padrao"`
}

// ReadYML read the yml file and return a slice of radios
func ReadYML() []*radio.Radio {
	var radiosYml RadioYML

	reader, err := os.Open("radios.yml")
	if err != nil {
		log.Fatal(err)
	}

	yamlFile, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Panicln("Não foi possível ler o arquivo radios.yml. Verifique se o arquivo existe no mesmo diretório do executável.")
	}

	err = yaml.Unmarshal(yamlFile, &radiosYml)
	if err != nil {
		log.Panicln("Não foi possível ler o arquivo radios.yml. Verifique se o arquivo existe e se está no formato correto.")
	}

	radios := make([]*radio.Radio, 0)
	for _, r := range radiosYml.Radios {
		radio := radio.NewRadio(r.Nome, r.Url, r.Pattern)
		radios = append(radios, radio)
	}

	return radios
}
