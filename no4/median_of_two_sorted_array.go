package no4

// 中间一个或两个数的平均数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var r float64
	l1 := len(nums1)
	l2 := len(nums2)
	var isOdd = ((l1 + l2) % 2) != 0 // 是否奇数个元素
	if nums1[l1-1] < nums2[0] {
		if isOdd {
			if l1 > l2 {
				return float64(nums1[l1-1]) / 1.0
			}
			return float64(nums2[0]) / 1.0
		} else {
			if l1 == l2 {
				return float64(nums1[l1-1]+nums2[0]) / 2.0
			} else if l1 > l2 {
				return float64(nums1[l1-1]+nums1[l1-2]) / 2.0
			} else {
				return float64(nums2[0]+nums2[1]) / 2.0
			}
		}
	} else {
		// TODO
	}

	return r
}
