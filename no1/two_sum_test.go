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
