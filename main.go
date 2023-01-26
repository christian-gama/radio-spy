package main

import (
	"runtime"

	"github.com/christian-gama/radio-spy/options"
	"github.com/christian-gama/radio-spy/radio"
	"github.com/christian-gama/radio-spy/yml"
)

func init() {
	yml.Settings = yml.ReadYML()
}

func main() {
	options := options.UnmarshalOptions(yml.Settings)
	radios := radio.UnmarshalRadios(yml.Settings)

	if options.Server {
		Server(radios, options)
	} else {
		Executable(radios, options)
	}

}

// ClearScreen clear the terminal screen
func ClearScreen() {
	value, ok := Clear[runtime.GOOS]
	if ok {
		value()
	}
}
