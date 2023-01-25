package finder

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Document returns a *goquery.Document object for the HTML content passed
// as a parameter. If the document is not successfully read, it panics.
func Document(res *http.Response) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Panic("Não foi possível ler o documento HTML")
	}

	return doc
}
