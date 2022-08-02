package gohttprouter

import (
	"bytes"
	"net/http"
	"strings"
)

// Converts an uppercase hexadecimal character to its integer form
func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	}
	return 0
}

// Converts an ASCII character to its uppercase form
func toUpper(c byte) byte {
	if 'a' <= c && c <= 'z' {
		return c - 'a' + 'A'
	}
	return c
}

// Returns the normalized version of http.Request.RequestURI
func (r *router) getPath(request *http.Request) string {
	// NOTE: RFC 2616 § 5.1.2 "Request-URI is a Uniform Resource Identifier"
	// That means, unless a new RFC supersedes it, RequestURI is not an IRI.
	url := []byte(request.RequestURI)
	// Trim #fragment and ?query
	for _, char := range []byte{'#', '?'} {
		if index := bytes.IndexByte(url, char); index != -1 {
			url = url[:index]
		}
	}
	// Truncate empty segments
	if !r.config.EmptySegmentsAreImportant {
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
	// Normalize percent-encoded characters
	for i := 0; i < len(url); i++ {
		if url[i] != '%' {
			continue
		}
		decoded := (unhex(url[i+2]) | unhex(url[i+1])<<4)
		if strings.IndexByte(rfc3986_pchar, decoded) == -1 {
			// The decoded character isn't allowed in path segments.
			// Just ensure that the hexadecimal characters are uppercase, and move along.
			url[i+1] = toUpper(url[i+1])
			url[i+2] = toUpper(url[i+2])
			continue
		}
		url = bytes.Join([][]byte{url[:i], {decoded}, url[i+3:]}, []byte{})
	}
	return string(url)
}
