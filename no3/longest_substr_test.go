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
