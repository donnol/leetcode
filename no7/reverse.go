package no7

func reverse(x int) int {
	r := 0
	for x != 0 {
		r = r*10 + x%10 // 已有数字往左移后加上个位数
		x = x / 10      // 剩下的部分
	}
	if r < -(1<<31) || r > 1<<31-1 {
		return 0
	}
	return r
}
