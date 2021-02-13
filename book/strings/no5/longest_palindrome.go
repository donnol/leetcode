package no5

func longestPalindrome(s string) (r string) {
	// 滑动窗口

	l := len(s)
	for i := 0; i < l; i++ {
		for j := i; j <= l; j++ {
			e := s[i:j] // j用在了切片后下标，因此需要取到s的长度
			if !isPalindrome(e) {
				continue
			}
			if len(e) > len(r) {
				r = e
			}
		}
	}
	return
}

func isPalindrome(s string) bool {
	rs := []rune(s)
	l := len(rs)
	isOdd := l%2 != 0

	if isOdd {
		l = (l - 1) / 2
		s1 := rs[:l]
		s2 := rs[l+1:]
		s2 = reverse(s2)
		if string(s1) == string(s2) {
			return true
		}
	} else {
		l = l / 2
		s1 := rs[:l]
		s2 := rs[l:]
		s2 = reverse(s2)
		if string(s1) == string(s2) {
			return true
		}
	}

	return false
}

func reverse(s []rune) (r []rune) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}
