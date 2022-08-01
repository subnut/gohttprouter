package gohttprouter

import (
	"bytes"
	"net/http"
)

type Router struct {
	config config
}

type config struct {
	EmptySegmentsAreImportant   bool
	TrailingSlashesAreImportant bool
}

func New() *Router {
	// Useful for setting default values that aren't the "nil" values
	return &Router{}
}

// NOTE: RFC 2616 §5.1.2 "Request-URI is a Uniform Resource Identifier"
// That means, unless a new RFC supersedes it, RequestURI is not an IRI (Internationalized Resource Identifier)
func (router *Router) getRequestPath(request *http.Request) string {
	url := []byte(request.RequestURI)
	// Trim #fragment and ?query
	for _, char := range []byte{'#', '?'} {
		if index := bytes.IndexByte(url, char); index != -1 {
			url = url[:index]
		}
	}
	// Ensure all percent-encodings have uppercase hexadecimal characters
	for i := bytes.IndexByte(url[:], '%'); i != -1; i = bytes.IndexByte(url[i:], '%') {
		if i++; url[i] >= 'a' {
			url[i] -= 'a' - 'A'
		}
		if i++; url[i] >= 'a' {
			url[i] -= 'a' - 'A'
		}
	}
	// Truncate empty segments
	if !router.config.EmptySegmentsAreImportant {
		var segments [][]byte
		// Leading forward slash (if any)
		if url[0] == '/' {
			segments = append(segments, []byte{})
		}
		// All non-empty segments
		for _, segment := range bytes.Split(url, []byte{'/'}) {
			if len(segment) != 0 {
				segments = append(segments, segment)
			}
		}
		// Trailing forward slash (if any)
		if url[len(url)-1] == '/' {
			segments = append(segments, []byte{})
		}
		url = bytes.Join(segments, []byte{'/'})
	}
	return string(url)
}

func (*Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
}
