package main

// read yml file and return a slice of radios

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/christian-gama/radio-spy/options"
	"github.com/christian-gama/radio-spy/radio"
	"gopkg.in/yaml.v3"
)

// SettingsYML is the struct that represents the yml file
type SettingsYML struct {
	Radios []Radio `yaml:"radios"`
	Opcoes Options `yaml:"opcoes"`
}

// Radio is the struct that represents a radio
type Radio struct {
	Nome    string `yaml:"nome"`
	Url     string `yaml:"url"`
	Pattern string `yaml:"padrao"`
}

// Options is the struct that represents the options
type Options struct {
	ShowQuantity bool `yaml:"mostrar_quantidade"`
}

// Settings is a SettingsYML struct
var Settings *SettingsYML

func init() {
	Settings = ReadYML()

}

// ReadYML read the yml file and return a slice of radios
func ReadYML() *SettingsYML {
	var settingsYml *SettingsYML

	reader, err := os.Open("settings.yml")
	if err != nil {
		log.Fatal(err)
	}

	yamlFile, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Panicln("Não foi possível ler o arquivo settings.yml. Verifique se o arquivo existe no mesmo diretório do executável.")
	}

	err = yaml.Unmarshal(yamlFile, &settingsYml)
	if err != nil {
		log.Panicln("Não foi possível ler o arquivo settings.yml. Verifique se o arquivo existe e se está no formato correto.")
	}

	return settingsYml
}

// CreateRadios create a slice of radios
func CreateRadios(settingsYml *SettingsYML) []*radio.Radio {
	var radios []*radio.Radio

	for _, r := range settingsYml.Radios {
		radios = append(radios, radio.NewRadio(r.Nome, r.Url, r.Pattern))
	}

	return radios
}

// CreateOptions create a slice of radios
func CreateOptions(settingsYml *SettingsYML) *options.Options {
	return options.NewOptions(settingsYml.Opcoes.ShowQuantity)
}
