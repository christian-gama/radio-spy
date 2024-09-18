package finder

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// FindListeners will look in the document for the pattern and return the number of listeners
func FindListeners(pattern string, html string) (uint32, error) {
	var listeners uint32

	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(html)

	if len(matches) > 1 {
		listenersStr := matches[1]
		listenersUint64, err := strconv.ParseUint(listenersStr, 10, 32)
		if err != nil {
			return 0, fmt.Errorf("Não foi possível converter o número de ouvintes para uint32")
		}

		listeners = uint32(listenersUint64)
	} else {
		log.Printf("Não foi possível encontrar o padrão '%s' no HTML do site: %s", pattern, html)
	}

	return listeners, nil
}
