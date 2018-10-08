package no1

func twoSum(nums []int, target int) []int {
	return twoSum1(nums, target)
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
