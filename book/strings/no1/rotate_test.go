package no1

import (
	"testing"
)

func TestRotate(t *testing.T) {
	for _, cas := range []struct {
		In   string
		N    int
		Want string
	}{
		{"abcdef", 2, "cdefab"},
	} {
		r := Rotate(cas.In, cas.N)
		if r != cas.Want {
			t.Fatalf("In: %s, %s != %s\n", cas.In, r, cas.Want)
		}
	}
}

func BenchmarkRotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rotate("abcdef", 2)
	}
}
