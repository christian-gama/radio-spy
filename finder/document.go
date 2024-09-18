package finder

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// Document returns a *goquery.Document object for the HTML content passed
// as a parameter. If the document is not successfully read, it panics.
func Document(res []byte) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(res))
	if err != nil {
		return nil, fmt.Errorf(
			"Ocorreu um erro ao tentar ler o HTML do site '%s'.",
			err,
		)
	}

	return doc, nil
}
