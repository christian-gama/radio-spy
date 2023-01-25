package radio

import (
	"log"
	"regexp"
)

// Radio is a struct that represents a radio and its metadata
type Radio struct {
	name             string
	url              string
	listeners        uint32
	listenersPattern string
}

// GetUrl returns the URL of the radio
func (r *Radio) GetUrl() string {
	return r.url
}

// SetUrl sets the URL of the radio
func (r *Radio) SetUrl(url string) {
	r.url = url
	r.validate()
}

// GetName returns the name of the radio
func (r *Radio) GetName() string {
	return r.name
}

// GetListeners returns the number of listeners of the radio
func (r *Radio) GetListeners() uint32 {
	return r.listeners
}

// GetListenersPattern returns the pattern used to extract the number of listeners of the radio
func (r *Radio) GetListenersPattern() string {
	return r.listenersPattern
}

func (r *Radio) SetListeners(listeners uint32) {
	r.listeners = listeners
}

// validate validates that the radio is correctly configured
func (r *Radio) validate() {
	if r.name == "" {
		log.Panic("O nome da rádio não pode ser vazio")
	}

	if len(r.name) > 50 {
		log.Panic("O nome da rádio não pode ter mais de 50 caracteres")
	}

	if r.url == "" {
		log.Panic("A URL da rádio não pode ser vazia")
	}

	urlRegex := regexp.MustCompile(`^(http|https)://[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`)
	if !urlRegex.MatchString(r.url) {
		log.Panicf("A URL da rádio é inválida. Espera-se o formato http://www.exemplo.com ou https://www.exemplo.com, mas foi usado %s", r.url)
	}

	if r.listenersPattern == "" {
		log.Panic("O padrão de extração de ouvintes da rádio não pode ser vazio")
	}

	if _, err := regexp.Compile(r.listenersPattern); err != nil {
		log.Panicf("O padrão de extração de ouvintes da rádio é inválido: %s", r.listenersPattern)
	}
}

// NewRadio creates a new radio with the given name and URL
func NewRadio(name string, url string, listenersPattern string) *Radio {
	radio := &Radio{
		name:             name,
		url:              url,
		listenersPattern: listenersPattern,
		listeners:        0,
	}

	radio.validate()

	return radio
}
