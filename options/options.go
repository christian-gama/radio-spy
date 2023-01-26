package options

import (
	"log"

	"github.com/christian-gama/radio-spy/yml"
)

// Options is the struct that contains the options of the application
type Options struct {
	ShowQuantity bool
	Server       bool
	Port         int
}

func (o *Options) validate() {
	if o.Port < 1024 || o.Port > 65535 {
		log.Panicf("A porta deve estar entre 1024 e 65535, mas foi informada a porta %d", o.Port)
	}
}

// NewOptions returns a new Options struct
func NewOptions(showQuantity bool, server bool, port int) *Options {
	options := &Options{
		ShowQuantity: showQuantity,
		Server:       server,
		Port:         port,
	}

	options.validate()

	return options
}

// UnmarshalOptions returns a new Options struct
func UnmarshalOptions(settingsYml *yml.SettingsYml) *Options {
	return NewOptions(settingsYml.Opcoes.ShowQuantity, settingsYml.Opcoes.Server, settingsYml.Opcoes.Port)
}
