package finder

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/url"
	"time"
)

// Get retrieves the content from the given URL and returns it as a byte slice.
// If the request fails, it returns an error.
func Get(urlStr string) ([]byte, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("Erro ao analisar a URL '%s': %v", urlStr, err)
	}

	// Determine host and port
	host := parsedURL.Hostname()
	port := parsedURL.Port()
	if port == "" {
		if parsedURL.Scheme == "https" {
			port = "443"
		} else {
			port = "80"
		}
	}

	// Create a connection
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 30*time.Second)
	if err != nil {
		return nil, fmt.Errorf("Erro ao conectar com '%s': %v", urlStr, err)
	}
	defer conn.Close()

	path := parsedURL.Path
	if path == "" {
		path = "/"
	}
	if parsedURL.RawQuery != "" {
		path += "?" + parsedURL.RawQuery
	}
	fmt.Fprintf(
		conn,
		"GET %s HTTP/1.0\r\nHost: %s\r\nAccept-Encoding: identity\r\n\r\n",
		path,
		host,
	)

	var buf bytes.Buffer
	_, err = io.Copy(&buf, conn)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("Erro ao ler a resposta de '%s': %v", urlStr, err)
	}

	response := buf.Bytes()
	parts := bytes.SplitN(response, []byte("\r\n\r\n"), 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("Resposta inválida de '%s'", urlStr)
	}

	return parts[1], nil
}
