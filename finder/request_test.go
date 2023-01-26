package finder

import (
	nhttp "net/http"
	"testing"
)

func TestGETRequestTo(t *testing.T) {
	res, _ := Get("https://www.google.com")
	if res.StatusCode != nhttp.StatusOK {
		t.Errorf("Expected status code %d, got %d", nhttp.StatusOK, res.StatusCode)
	}
}

func TestGETRequestToPanics(t *testing.T) {
	_, err := Get("https://www.google.com/404")
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestGETRequestPanicsWithInvalidUrl(t *testing.T) {
	_, err := Get("invalid url")
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

}
