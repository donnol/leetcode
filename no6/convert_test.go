package no6

import "testing"

func TestConvert(t *testing.T) {
	for _, cas := range []struct {
		s       string
		numRows int
		r       string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	} {
		r := convert(cas.s, cas.numRows)
		if r != cas.r {
			t.Fatalf("Bad result: %s != %s\n", r, cas.r)
		}
	}
}
