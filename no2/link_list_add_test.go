package no2

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := NewListNode(1, 8)
	l2 := NewListNode(0)
	want := NewListNode(1, 8)
	r := addTwoNumbers(l1, l2)
	reflect.DeepEqual(r, want)
}
