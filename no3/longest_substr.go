package no3

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}

	return lengthOfLongestSubstring2(s)
}

func lengthOfLongestSubstring1(s string) int {
	var l int
	var m = make(map[byte]int)

finishLabel:
	for i := 0; i < len(s); i++ {
		sub := s[i]
		m[sub] = i // 记录已出现过的
		for j := i + 1; j < len(s); j++ {
			k := s[j]
			if _, ok := m[k]; ok { // 如果出现了重复的，则统计长度并跳出
				tmp := j - i
				if tmp > l {
					l = tmp
				}
				break
			} else {
				m[k] = j           // 记录已出现过的
				if j == len(s)-1 { // 如果刚好到末尾
					tmp := len(m)
					if tmp > l {
						l = tmp
					}
					break finishLabel
				}
			}
		}
		m = make(map[byte]int)
	}

	return l
}

func lengthOfLongestSubstring2(s string) int {
	var l int
	sl := len(s)

	var sm = make(map[int]byte)
	var em = make(map[byte]int)
	// 遍历字符串-拆细
	// 什么时候开始，什么时候结束
	for i := 0; i < sl; i++ {
		// 拿到字符，并记录字符位置
		sub := s[i]

		// 当出现重复字符，则上个重复字符出现的位置及之前的子串可忽略
		if oi, ok := em[sub]; ok {
			tl := len(em)
			if tl > l {
				l = tl
			}

			// 重置
			em = make(map[byte]int)
			for k, v := range sm {
				if k <= oi {
					delete(sm, k)
				} else {
					em[v] = k
				}
			}
		}

		sm[i] = sub
		em[sub] = i
	}
	tl := len(em)
	if tl > l {
		l = tl
	}

	return l
}
