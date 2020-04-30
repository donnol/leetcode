package no3

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, cas := range []struct {
		Name  string
		Input string
		Want  int
	}{
		{"No1", "abcabcbb", 3},
		{"No2", "bbbbb", 1},
		{"No3", "pwwkew", 3},
		{"No4", "dvdf", 3},
		{"No5", " ", 1},
		{"No6", "umvejcuuk", 6},
		{"No7", "bbtablud", 6},
	} {
		t.Run(cas.Name, func(t *testing.T) {
			r := lengthOfLongestSubstring(cas.Input)
			if r != cas.Want {
				t.Fatalf("bad result, %d != %d\n", r, cas.Want)
			}
		})
	}
}

func BenchmarkLengthOfLongestSubstring1(b *testing.B) {
	in := "bbtabludbbtabludbbtabludbbtabludbbtabludbbtabludbbtablud"
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring1(in)
	}

	// BenchmarkLengthOfLongestSubstring1-8   	   77547	     13269 ns/op	    7200 B/op	     100 allocs/op
}

func BenchmarkLengthOfLongestSubstring2(b *testing.B) {
	in := "bbtabludbbtabludbbtabludbbtabludbbtabludbbtabludbbtablud"
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring2(in)
	}

	// BenchmarkLengthOfLongestSubstring2-8   	  133368	      9087 ns/op	    2944 B/op	      42 allocs/op
}

func BenchmarkLengthOfLongestSubstring3(b *testing.B) {
	in := "bbtabludbbtabludbbtabludbbtabludbbtabludbbtabludbbtablud"
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring3(in)
	}

	// BenchmarkLengthOfLongestSubstring3-8   	 1269602	       901 ns/op	     280 B/op	      56 allocs/op
}
