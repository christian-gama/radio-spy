package yml

// read yml file and return a slice of radios

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// SettingsYml is the struct that represents the yml file
type SettingsYml struct {
	Radios []radioYml `yaml:"radios"`
	Opcoes optionsYml `yaml:"opcoes"`
}

// radio is the struct that represents a radioYml
type radioYml struct {
	Nome    string `yaml:"nome"`
	Url     string `yaml:"url"`
	Pattern string `yaml:"padrao"`
}

// options is the struct that represents the optionsYml
type optionsYml struct {
	ShowQuantity bool `yaml:"mostrar_quantidade"`
	Server       bool `yaml:"servidor"`
	Port         int  `yaml:"porta"`
}

// Settings is a SettingsYML struct
var Settings *SettingsYml

// ReadYML read the yml file and return a slice of radios
func ReadYML() *SettingsYml {
	var settingsYml *SettingsYml

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
