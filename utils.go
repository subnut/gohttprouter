package gohttprouter

import (
	"fmt"
	str "strings"
)

func normalizeURL(uri string) (string, error) {
	url := []byte(uri)
	// Trim #fragment and ?query
	for _, char := range []byte{'#', '?'} {
		if index := str.IndexByte(uri, char); index != -1 {
			uri = uri[:index]
		}
	}
	for i := str.IndexByte(uri[:], '%'); i != -1; i = str.IndexByte(uri[i:], '%') {
		if i+2 >= len(uri) {
			return "", fmt.Errorf("invalid escape code %q encountered while parsing URI %q\n", uri[i:], uri)
		}
		if i++; uri[i] >= 'a' {
			url[i] -= 'a' - 'A'
		}
		if i++; uri[i] >= 'a' {
			url[i] -= 'a' - 'A'
		}
	}
	return string(url), nil
}
