package gohttprouter

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct{ err int; in, out string }{
		{ 0, "/",	"/"	},
		{ 0, "/%20",	"/%20"	},
		{ 0, "/%2a",	"/%2A"	},
		{ 1, "/%2",	""	},
	}
	for _, test := range tests {
		out, err := normalizeURL(test.in)
		if err != nil && test.err == 0 {
			t.Errorf("normalizeURL(%v) -- UNEXPECTED ERROR -- %v", test.in, err)
			continue
		}
		if err == nil && test.err != 0 {
			t.Errorf("normalizeURL(%v) -- NO ERROR RETURNED", test.in)
			continue
		}
		if out != test.out {
			t.Errorf("normalizeURL(%v)\nwant: %v\n got: %v\n", test.in, test.out, out)
		}
	}
}
