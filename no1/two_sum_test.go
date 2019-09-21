package no1

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	var want = []int{0, 1}
	r := twoSum(nums, target)
	if !reflect.DeepEqual(r, want) {
		t.Fatalf("bad result, r is %v, want is %v\n", r, want)
	}
}

func BenchmarkTwoSum(b *testing.B) {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}

func BenchmarkTwoSum1(b *testing.B) {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	for i := 0; i < b.N; i++ {
		twoSum1(nums, target)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	for i := 0; i < b.N; i++ {
		twoSum2(nums, target)
	}
}
