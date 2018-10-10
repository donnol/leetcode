package no3

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}

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
