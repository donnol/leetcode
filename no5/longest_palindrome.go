package no5

func longestPalindrome(s string) string {
	return longestPalindrome1(s)
}

func longestPalindrome1(s string) string {
	l := len(s)
	if l == 0 || l == 1 {
		return s
	}

	r := string(s[0])
	for i := 0; i < l; i++ { // 每遍历一个字符，查看前后的字符是否相同
		sub := s[i]
		substr := string(sub)

		// 重叠回文

		str := substr
		for j := 1; ; j++ {
			var left, right byte
			if i-j >= 0 {
				left = s[i-j]
			}
			if i+j <= l-1 {
				right = s[i+j]
			}
			if left == 0 && right == 0 {
				break
			} else if left == 0 {
				if sub != right {
					break
				}
				str = substr + string(right)
			} else if right == 0 {
				if left != sub {
					break
				}
				str = str + substr
			} else {
				if left != right {
					if sub != right {
						break
					}
					str = substr + string(right)
					break
				} else {
					str = string(left) + str + string(right)
				}
			}
		}
		if len(str) > 1 && len(str) >= len(r) {
			r = str
		}
	}

	return r
}
