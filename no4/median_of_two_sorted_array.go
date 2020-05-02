package no4

import (
	"sort"
)

// 中间一个或两个数的平均数
// 根据输入数组的长度，可以知道中位数的位置，只要找到该位置的一个/两个数字即可计算出中位数的值
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays2(nums1, nums2)
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	var r float64

	l1 := len(nums1)
	l2 := len(nums2)

	// 如果存在数组为空
	if l1 == 0 && l2 == 0 {
		return r
	} else if l1 == 0 {
		return getMid(nums2, 0)
	} else if l2 == 0 {
		return getMid(nums1, 0)
	}

	// 如果大小数组位置互换

	la := (l1 + l2)
	mid := (la / 2)
	isOdd := (la % 2) != 0 // 是否奇数个元素

	if nums1[l1-1] <= nums2[0] { // 首个数组的最大值不大于次数组的最小值
		if isOdd {
			if l1 > l2 {
				return float64(nums1[mid]) / 1.0
			}
			return float64(nums2[mid-l1]) / 1.0
		} else {
			if l1 == l2 {
				return float64(nums1[l1-1]+nums2[0]) / 2.0
			} else if l1 > l2 {
				return float64(nums1[mid-1]+nums1[mid]) / 2.0
			} else {
				return float64(nums2[mid-l1-1]+nums2[mid-l1]) / 2.0
			}
		}
	} else {
		// 找到公共部分
		tnums1 := make([]int, 0, len(nums1))
		for i := l1 - 1; i >= 0; i-- {
			n1 := nums1[i]
			if n1 >= nums2[0] {
				tnums1 = append(tnums1, n1)
				continue
			}
			break
		}
		tnums2 := make([]int, 0, len(nums2))
		for i := l2 - 1; i >= 0; i-- {
			n2 := nums2[i]
			if n2 <= nums1[l1-1] {
				tnums2 = append(tnums2, n2)
				continue
			}
			break
		}
		index := mid - (l1 - len(tnums1))

		tnums := append(tnums1, tnums2...)

		sort.Ints(tnums)

		return getMid(tnums, index)
		// if isOdd {
		// 	return float64(tnums[index]) / 1.0
		// } else {
		// 	fmt.Printf("%+v, %d\n", tnums, index)
		// 	return float64(tnums[index-1]+tnums[index]) / 2.0
		// }
	}
}

func getMid(nums []int, index int) float64 {
	l := len(nums)
	if l == 0 {
		return 0.0
	} else if l == 1 {
		return float64(nums[0])
	} else if l == 2 {
		return float64(nums[0]+nums[1]) / 2.0
	} else {
		mid := l / 2
		isOdd := (l % 2) != 0
		if isOdd {
			return float64(nums[mid-index]) / 1.0
		} else {
			return float64(nums[mid-1-index]+nums[mid-index]) / 2.0
		}
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)

	left := (n + m + 1) / 2
	right := (n + m + 2) / 2

	return float64(getKth(nums1, 0, n-1, nums2, 0, m-1, left)+getKth(nums1, 0, n-1, nums2, 0, m-1, right)) * 0.5
}

func getKth(nums1 []int, start1, end1 int, nums2 []int, start2, end2 int, k int) int {
	l1 := end1 - start1 + 1
	l2 := end2 - start2 + 1

	if l1 > l2 {
		return getKth(nums2, start2, end2, nums1, start1, end1, k)
	}
	if l1 == 0 {
		return nums2[start2+k-1]
	}
	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}

	i := start1 + min(l1, k/2) - 1
	j := start2 + min(l2, k/2) - 1

	if nums1[i] > nums2[j] {
		return getKth(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
	} else {
		return getKth(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
