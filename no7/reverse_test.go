package no7

import "testing"

func TestReverse(t *testing.T) {
	for _, cas := range []struct {
		x int
		r int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
	} {
		r := reverse(cas.x)
		if r != cas.r {
			t.Fatalf("Bad result: %d != %d\n", r, cas.r)
		}
	}
}
