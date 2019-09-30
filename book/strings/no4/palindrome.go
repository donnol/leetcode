package no4

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
