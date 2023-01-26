package finder

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Document returns a *goquery.Document object for the HTML content passed
// as a parameter. If the document is not successfully read, it panics.
func Document(res *http.Response) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Ocorreu um erro ao tentar ler o HTML do site '%s'.", res.Request.URL)
	}

	return doc, nil
}
