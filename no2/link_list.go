package no2

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode 新建
func NewListNode(nums ...int) *ListNode {
	var r = new(ListNode)

	var tmpR = r
	for i, v := range nums {
		if i > 0 {
			tmpNode := new(ListNode)
			tmpR.Next = tmpNode
			tmpR = tmpNode
		}
		tmpR.Val = v
	}

	return r
}
