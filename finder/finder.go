package finder

import (
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// FindListeners will look in the document for the pattern and return the number of listeners
func FindListeners(pattern string, doc *goquery.Document) uint32 {
	var listeners uint32

	doc.Find(pattern).Each(func(i int, s *goquery.Selection) {
		listenersUint64, err := strconv.ParseUint(s.Text(), 10, 32)
		if err != nil {
			log.Panicf("Não foi possível converter '%s' para uint32", s.Text())
		}

		listeners = uint32(listenersUint64)

	})

	return listeners
}
