package http

import (
	"log"
	"net/http"
)

// GETRequestTo returns a response to an HTTP GET request to the url passed as
// a parameter. If the request fails, it panics.
func GETRequestTo(url string) *http.Response {
	res, err := http.Get(url)
	if err != nil {
		log.Panicf("Não foi possível se conectar ao site '%s'", url)
	}

	if res.StatusCode != http.StatusOK {
		log.Panicf("O site '%s' retornou o status code %d. Por favor, verifique se o site está disponível.", url, res.StatusCode)
	}

	return res
}
