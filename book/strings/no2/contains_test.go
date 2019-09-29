package no2

import (
	"testing"
)

func TestRotate(t *testing.T) {
	for _, cas := range []struct {
		A    string
		B    string
		Want bool
	}{
		{"abcdef", "cdefa", true},
		{"abcdef", "cdefg", false},
	} {
		r := Contains(cas.A, cas.B)
		if r != cas.Want {
			t.Fatalf("In: %s, %s, %v != %v\n", cas.A, cas.B, r, cas.Want)
		}
	}
}

func BenchmarkRotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Contains("abcdef", "cdefa")
	}
}
