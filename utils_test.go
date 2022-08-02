package gohttprouter

import (
	"net/http"
	"testing"
)

func Test_unhex(t *testing.T) {
	tests := []struct{ in, out byte }{
		{'0', 0}, {'1', 1}, {'2', 2}, {'3', 3}, {'4', 4}, {'5', 5}, {'6', 6}, {'7', 7}, {'8', 8}, {'9', 9},
		{'a', 10}, {'b', 11}, {'c', 12}, {'d', 13}, {'e', 14}, {'f', 15}, {'g', 0},
		{'A', 10}, {'B', 11}, {'C', 12}, {'D', 13}, {'E', 14}, {'F', 15}, {'G', 0},
		{},
	}
	for _, test := range tests {
		if out := unhex(test.in); out != test.out {
			t.Errorf("getPath(%v)\nwant: %v\n got: %v\n", test.in, test.out, out)
		}
	}
}

func Test_toUpper(t *testing.T) {
	low := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_!=+?&~@ #")
	up := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_!=+?&~@ #")
	for i, c := range low {
		if out := toUpper(c); out != up[i] {
			t.Errorf("toUpper(%c)\nwant: %c\n got: %c\n", c, up[i], out)
		}
	}
}

func TestRouter_getPath(t *testing.T) {
	router := New()
	tests := []struct{ in, out string }{
		// NOTE: all URLs must be valid
		{"/", "/"},
		{"//", "/"},
		{"///", "/"},
		{"////", "/"},
		{"/%20", "/%20"},
		{"/%2f", "/%2F"},
		{"/%3Fa", "/%3Fa"},
		{"/hi///there", "/hi/there"},
		{"/hi///there/", "/hi/there/"},
		{"/hi///there////", "/hi/there/"},
		{"/hi///there//%20//", "/hi/there/%20/"},
		{"/hi#abc?xyz#123", "/hi"},
		{"/hi?abc#xyz?123", "/hi"},
		{"/hi?abc#123?xyz#456", "/hi"},
		{"/hi#abc?123#xyz?456", "/hi"},
		{"/%48%65%6C%6C%6F%2C%20%28%55%73%65%72%29%2E%20%4E%69%63%65%20%74%6F%20%6D%65%65%74%20%79%6F%75%21", "/Hello,%20(User).%20Nice%20to%20meet%20you!"},
	}
	for _, test := range tests {
		if out := router.getPath(&http.Request{RequestURI: test.in}); out != test.out {
			t.Errorf("getPath(%v)\nwant: %v\n got: %v\n", test.in, test.out, out)
		}
	}
}
func TestRouter_getPath_EmptySegmentsAreImportant(t *testing.T) {
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
		if out := router.getPath(&http.Request{RequestURI: url}); out != url {
			t.Errorf("getPath(%v)\nwant: %v\n got: %v\n", url, url, out)
		}
	}
}
