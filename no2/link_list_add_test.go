package no2

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := NewListNode(1, 8, 9, 9)
	l2 := NewListNode(0, 1, 1, 0)
	want := NewListNode(1, 9, 0, 0, 1)
	r := addTwoNumbers(l1, l2)
	reflect.DeepEqual(r, want)
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	l1 := NewListNode(1, 8)
	l2 := NewListNode(0)
	for i := 0; i < b.N; i++ {
		addTwoNumbers(l1, l2)
	}
}
