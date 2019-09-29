package no1

// Rotate 旋转
func Rotate(s string, n int) (r string) {
	l := len(s)
	if n > l {
		panic("n is too big")
	}
	if n == l {
		return s
	}

	// 收集指定长度的字符串
	s1 := s[:n]
	s2 := s[n:]

	// 移动到后面
	r = s2 + s1

	return
}
