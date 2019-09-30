package no4

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	for i, cas := range []struct {
		In   string
		Want bool
	}{
		{"madam", true},
		{"madae", false},
		{"abccba", true},
		{"我是我", true},
		{"我我是我我", true},
		{"我我是妮妮", false},
	} {
		r := isPalindrome(cas.In)
		if r != cas.Want {
			t.Fatalf("No.%d: %v != %v\n", i+1, r, cas.Want)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome("madam")
	}
}
