package no2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry bool // 是否有进位
	var tmpVal int
	var n1, n2 = l1, l2
	var count int
	var nums = make([]int, 0)
	for {
		var n1Val, n2Val int
		if count > 0 {
			if n1 != nil {
				n1 = n1.Next
			}
			if n2 != nil {
				n2 = n2.Next
			}
		}
		count++
		if n1 == nil && n2 == nil { // 最后一位相加后有进位
			if carry {
				nums = append(nums, 1)
			}
			break
		}
		if n1 != nil {
			n1Val = n1.Val
		}
		if n2 != nil {
			n2Val = n2.Val
		}
		tmpVal = n1Val + n2Val
		if carry {
			tmpVal++
		}
		if tmpVal >= 10 {
			carry = true
			// 取个位
			tmpVal -= 10
		} else {
			carry = false
		}
		nums = append(nums, tmpVal)

		if !carry && (n1 != nil && n1.Next == nil && n2 != nil && n2.Next == nil) {
			break
		}
	}

	return NewListNode(nums...)
}
