package finder

import (
	"testing"
)

func TestDocument(t *testing.T) {
	res, _ := Get("https://www.google.com")
	doc, _ := Document(res)

	if doc == nil {
		t.Error("Expected a non-nil *goquery.Document, got nil")
	}
}
