package radio

import (
	"log"
	"regexp"
)

// Radio is a struct that represents a radio and its metadata
type Radio struct {
	name      string
	url       string
	listeners uint32
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

// SetName sets the name of the radio
func (r *Radio) SetName(name string) {
	r.name = name
	r.validate()
}

// GetListeners returns the number of listeners of the radio
func (r *Radio) GetListeners() uint32 {
	return r.listeners
}

// SetListeners sets the number of listeners of the radio
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

	urlRegex, err := regexp.Compile(`^(http|https)://[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`)
	if err != nil {
		log.Panic("Erro ao compilar a expressão regular que identifica uma URL")
	}

	if !urlRegex.MatchString(r.url) {
		log.Panicf("A URL da rádio é inválida. Espera-se o formato http://www.exemplo.com ou https://www.exemplo.com, mas foi usado %s", r.url)
	}
}

// NewRadio creates a new radio with the given name and URL
func NewRadio(name string, url string) *Radio {
	radio := &Radio{
		name:      name,
		url:       url,
		listeners: 0,
	}

	radio.validate()

	return radio
}
