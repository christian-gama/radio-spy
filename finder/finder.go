package finder

import (
	"log"
	"regexp"
	"strconv"
)

// FindListeners will look in the document for the pattern and return the number of listeners
func FindListeners(pattern string, html string) uint32 {
	var listeners uint32

	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(html)

	if len(matches) > 1 {
		listenersStr := matches[1]
		listenersUint64, err := strconv.ParseUint(listenersStr, 10, 32)
		if err != nil {
			log.Panicf("Não foi possível converter o número de ouvintes para uint32: %s", listenersStr)
		}

		listeners = uint32(listenersUint64)
	} else {
		log.Panicf("Não foi possível encontrar o número de ouvintes no HTML: %s", html)
	}

	return listeners
}
