package no3

func atoi(s string) (r int) {
	l := len(s)

	for i, t := range []byte(s) {
		if i == 0 &&
			(t == '0' || t == '-' || t == '+') {
			continue
		}

		if t >= '0' && t <= '9' {
			r += pow(l-i) * (int(t) - 48)
		} else {
			panic("非法字符")
		}
	}

	// 符号
	if s[0] == '-' {
		r = -r
	}

	return
}

func pow(n int) int {
	r := 1
	base := 10
	for i := 0; i < n-1; i++ {
		r *= base
	}
	return r
}
