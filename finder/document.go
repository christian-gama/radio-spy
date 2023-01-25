package finder

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// GetRequestTo returns a response to an HTTP GET request to the url passed as
// a parameter. If the request fails, it panics.
func GetRequestTo(url string) *http.Response {
	res, err := http.Get(url)
	if err != nil {
		log.Panicf("Não foi possível se conectar ao site '%s'", url)
	}

	return res
}

// Document returns a *goquery.Document object for the HTML content passed
// as a parameter. If the document is not successfully read, it panics.
func Document(res *http.Response) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Panic("Não foi possível ler o documento HTML")
	}

	return doc
}
