package utils

import "testing"

func TestURItoURL(t *testing.T) {
	tests := []struct{ err int; in, out string }{
		{ 0, "/",	"/"	},
		{ 0, "/%20",	"/%20"	},
		{ 0, "/%2a",	"/%2A"	},
		{ 1, "/%2",	""	},
	}
	for _, test := range tests {
		out, err := URItoURL(test.in)
		if err != nil && test.err == 0 {
			t.Errorf("URItoURL(%v) -- UNEXPECTED ERROR -- %v", test.in, err)
			continue
		}
		if err == nil && test.err != 0 {
			t.Errorf("URItoURL(%v) -- NO ERROR RETURNED", test.in)
		}
		if out != test.out {
			t.Errorf("URItoURL(%v)\nwant: %v\n got: %v\n", test.in, test.out, out)
		}
	}
}
