package no3

import (
	"strconv"
	"testing"
)

func TestAtoi(t *testing.T) {
	for _, cas := range []struct {
		In   string
		Want int
	}{
		{"0", 0},
		{"-0", 0},
		{"123", 123},
		{"0123", 123},
		{"1234", 1234},
		{"12345", 12345},
		{"120345", 120345},
		{"-120345", -120345},
		{"+120345", 120345},
	} {
		r := atoi(cas.In)
		if r != cas.Want {
			t.Fatalf("%v != %v", r, cas.Want)
		}
	}
}

func BenchmarkAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atoi("1234")
	}
}

func BenchmarkAtoiStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Atoi("1234")
	}
}
