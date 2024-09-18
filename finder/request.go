package finder

import (
	"bufio"
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
	parsedURL, err := parseURL(urlStr)
	if err != nil {
		return nil, err
	}

	host, port := determineHostAndPort(parsedURL)

	conn, err := connectToServer(host, port, urlStr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	sendHTTPRequest(conn, parsedURL, host)

	return readResponse(conn)
}

// parseURL converts a string URL into a structured URL object.
// This is necessary to easily extract components like hostname and path.
func parseURL(urlStr string) (*url.URL, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL '%s': %v", urlStr, err)
	}
	return parsedURL, nil
}

// determineHostAndPort extracts the host and port from the URL.
// If no port is specified, it defaults to 80 for HTTP.
func determineHostAndPort(parsedURL *url.URL) (string, string) {
	host := parsedURL.Hostname()
	port := parsedURL.Port()
	if port == "" {
		port = "80"
	}
	return host, port
}

// connectToServer establishes a TCP connection to the target server.
// We use a timeout to prevent hanging on unresponsive servers.
func connectToServer(host, port, urlStr string) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 30*time.Second)
	if err != nil {
		return nil, fmt.Errorf("error connecting to '%s': %v", urlStr, err)
	}
	return conn, nil
}

// sendHTTPRequest writes the HTTP GET request to the connection.
// This function formats the request according to HTTP/1.0 specifications.
func sendHTTPRequest(conn net.Conn, parsedURL *url.URL, host string) {
	fmt.Fprintf(conn, "GET %s HTTP/1.0\r\nHost: %s\r\n\r\n", parsedURL.RequestURI(), host)
}

// readResponse handles reading the entire HTTP response from the server.
func readResponse(conn net.Conn) ([]byte, error) {
	reader := bufio.NewReader(conn)

	if err := discardHeaders(reader); err != nil {
		return nil, err
	}

	return readBody(reader)
}

// discardHeaders reads and discards the HTTP headers from the response.
// This is necessary because we're only interested in the response body,
// and headers end with a blank line (CRLF).
func discardHeaders(reader *bufio.Reader) error {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("error reading headers: %v", err)
		}
		if line == "\r\n" {
			break
		}
	}
	return nil
}

// readBody reads the HTTP response body into a byte slice.
// We use a bytes.Buffer for efficient appending of the incoming data.
func readBody(reader *bufio.Reader) ([]byte, error) {
	var body bytes.Buffer
	_, err := io.Copy(&body, reader)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("error reading body: %v", err)
	}
	return body.Bytes(), nil
}
