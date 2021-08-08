package no8

func myAtoi(s string) int {
	if s == "" {
		return 0
	}

	// 移除前面的空格
	i := 0
	for _, ch := range []byte(s) {
		if ch == ' ' {
			i++
			continue
		}
		break
	}
	if i == len(s) {
		return 0
	}
	if i != 0 {
		s = s[i:]
	}

	var s0 byte
	if s[0] == '-' || s[0] == '+' {
		s0 = s[0]
		s = s[1:]
		if len(s) < 1 {
			return 0
		}
	}

	n := 0
	for _, ch := range []byte(s) {
		// 确认是数字
		ch -= '0'
		if ch > 9 {
			break
		}

		// 已有数字乘10，留出个位给新的数字
		// 由于go可以直接通过类型转换，将byte转为int，非常方便
		n = n*10 + int(ch)

		// 溢出校验
		if n > 2<<32 {
			break
		}
	}

	// 根据符号，对溢出情况做补偿
	switch s0 {
	case '-':
		if n > (2 << 30) {
			n = 2 << 30
		}
		n = -n
	default:
		if n > 2<<30-1 {
			n = 2<<30 - 1
		}
	}

	return n
}
