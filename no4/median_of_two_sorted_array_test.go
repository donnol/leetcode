package no4

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	for _, cas := range []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "1",
			nums1: []int{1, 2},
			nums2: []int{3},
			want:  2.0,
		},
		{
			name:  "2",
			nums1: []int{1},
			nums2: []int{2, 3},
			want:  2.0,
		},
		{
			name:  "3",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
		{
			name:  "4",
			nums1: []int{1, 2, 3},
			nums2: []int{4},
			want:  2.5,
		},
		{
			name:  "5",
			nums1: []int{1},
			nums2: []int{2, 3, 4},
			want:  2.5,
		},
		{
			name:  "6",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			name:  "7",
			nums1: []int{},
			nums2: []int{1, 2, 3, 4},
			want:  2.5,
		},
		{
			name:  "8",
			nums1: []int{},
			nums2: []int{1, 2, 3, 4, 5},
			want:  3.0,
		},
		{
			name:  "9",
			nums1: []int{1, 2},
			nums2: []int{-1, 3},
			want:  1.5,
		},
	} {
		r := findMedianSortedArrays(cas.nums1, cas.nums2)
		if r != cas.want {
			t.Fatalf("No.%s: bad result, %v != %v\n", cas.name, r, cas.want)
		}
	}
}

func BenchmarkFindMedianSortedArrays(b *testing.B) {
	nums1 := []int{1, 2}
	nums2 := []int{3}

	for i := 0; i < b.N; i++ {
		findMedianSortedArrays(nums1, nums2)
	}
}
