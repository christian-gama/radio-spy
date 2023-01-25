package http

import (
	nhttp "net/http"
	"testing"
)

func TestGETRequestTo(t *testing.T) {
	res := GETRequestTo("https://www.google.com")
	if res.StatusCode != nhttp.StatusOK {
		t.Errorf("Expected status code %d, got %d", nhttp.StatusOK, res.StatusCode)
	}
}

func TestGETRequestToPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	GETRequestTo("https://www.google.com/404")
}

func TestGETRequestPanicsWithInvalidUrl(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	GETRequestTo("invalid url")
}
