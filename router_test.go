package gohttprouter

import (
	"net/http"
	"testing"
)

func TestRouter_getRequestPath(t *testing.T) {
	router := New()
	tests := []struct{ in, out string }{
		// NOTE: all URLs must be valid
		{"/", "/"},
		{"//", "/"},
		{"///", "/"},
		{"////", "/"},
		{"/%20", "/%20"},
		{"/%2a", "/%2A"},
		{"/hi///there", "/hi/there"},
		{"/hi///there/", "/hi/there/"},
		{"/hi///there////", "/hi/there/"},
		{"/hi#abc?xyz#123", "/hi"},
		{"/hi?abc#xyz?123", "/hi"},
		{"/hi?abc#123?xyz#456", "/hi"},
		{"/hi#abc?123#xyz?456", "/hi"},
	}
	for _, test := range tests {
		if out := router.getRequestPath(&http.Request{RequestURI: test.in}); out != test.out {
			t.Errorf("getRequestPath(%v)\nwant: %v\n got: %v\n", test.in, test.out, out)
		}
	}
}
func TestRouter_getRequestPath_EmptySegmentsAreImportant(t *testing.T) {
	router := New()
	router.config.EmptySegmentsAreImportant = true
	urls := []string{
		"/",
		"//",
		"///",
		"////",
		"/hi///there",
		"/hi///there/",
		"/hi///there////",
	}
	for _, url := range urls {
		if out := router.getRequestPath(&http.Request{RequestURI: url}); out != url {
			t.Errorf("getRequestPath(%v)\nwant: %v\n got: %v\n", url, url, out)
		}
	}
}
