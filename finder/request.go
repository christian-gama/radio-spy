package finder

import (
	"fmt"
	"net/http"
)

// Get returns a response to an HTTP GET request to the url passed as
// a parameter. If the request fails, it panics.
func Get(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Ocorreu um erro ao tentar fazer uma requisição GET para o site '%s'. Por favor, verifique se o site está disponível.", url)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Ocorreu um erro ao tentar fazer uma requisição GET para o site '%s'. Por favor, verifique se o site está disponível.", url)
	}

	return res, nil
}
