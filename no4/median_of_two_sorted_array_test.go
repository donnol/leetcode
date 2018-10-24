package no4

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3}
	want := 2.0
	r := findMedianSortedArrays(nums1, nums2)
	if r != want {
		t.Fatalf("bad result, %v != %v\n", r, want)
	}
}
