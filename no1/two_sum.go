package no1

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	return twoSum3(nums, target)
}

func twoSum1(nums []int, target int) []int {
	var r = make([]int, 0)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				r = append(r, i, j)
				break
			}
		}
	}
	return r
}

func twoSum2(nums []int, target int) []int {
	var r = make([]int, 0)

	var m = make(map[int][]int) // 值 -> 下标数组
	for i, s := range nums {
		m[s] = append(m[s], i)
	}

	sort.Ints(nums)

	l := len(nums)
	for i := 0; i < l; i++ {
		x := nums[i]
		for j := i + 1; j < l; j++ {
			y := nums[j]
			if x+y == target {
				is := m[x]
				js := m[y]
				ai := is[0]
				aj := js[0]
				if aj == ai {
					aj = js[1]
				}
				r = append(r, ai, aj)
				break
			} else if x+y > target {
				break
			}
		}
	}
	return r
}

func twoSum3(nums []int, target int) []int {
	var r = make([]int, 2)

	var m = make(map[int]int) // target-当前值 -> 下标

	for i, s := range nums {
		d := target - s
		if v, ok := m[d]; ok {
			r[0] = v
			r[1] = i
			return r
		}
		m[s] = i
	}

	return r
}
