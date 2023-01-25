package finder

import (
	"testing"

	"github.com/christian-gama/radio-spy/http"
)

func TestDocument(t *testing.T) {
	res := http.MustGet("https://www.google.com")
	doc := Document(res)

	if doc == nil {
		t.Error("Expected a non-nil *goquery.Document, got nil")
	}
}

func TestDocumentPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	res := http.MustGet("https://www.google.com/404")
	Document(res)
}
