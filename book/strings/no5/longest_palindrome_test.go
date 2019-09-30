package no5

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	for i, cas := range []struct {
		In   string
		Want string
	}{
		{"aaaaaamadam", "madam"},
	} {
		r := longestPalindrome(cas.In)
		if r != cas.Want {
			t.Fatalf("No.%d: %v != %v\n", i+1, r, cas.Want)
		}
	}
}
