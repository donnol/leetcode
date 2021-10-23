package no10

func isMatch(s string, p string) bool { //nolint
	lofs := len(s)
	lofp := len(p)

	if lofs == 0 {
		return false
	}
	if lofp == 0 {
		return false
	}

	// 先遍历一次p，获取.和*的位置信息，记录到map中
	mdot := make(map[int]struct{})
	mstar := make(map[int]struct{})
	for i := 0; i < lofp; i++ {
		if p[i] == '.' {
			mdot[i] = struct{}{}
		}
		if p[i] == '*' {
			mstar[i] = struct{}{}
		}
	}

	if len(mdot) == 0 && len(mstar) == 0 && lofs != lofp {
		return false
	}
	if len(mdot) == 0 && len(mstar) == 0 && lofs == lofp && s != p {
		return false
	}
	if len(mdot) == 0 && len(mstar) == 0 && lofs == lofp && s == p {
		return true
	}

	// 如果是匹配的，那么s在对应位置必须含有对应字符
	ignoreIndex := -1
	ignoreNumber := 0
	lastIndex := -1
	lastIsSame := false
	for j := 0; j < lofp; j++ {
		_, isdot := mdot[j]
		_, isstar := mstar[j]

		lastIndex = j - ignoreNumber

		switch {
		case isdot:
			if lofs < j+1 {
				return false
			}
		case isstar:
			b := p[j-1]
			if b == '.' {
				continue
			}
			if ignoreIndex != -1 && ignoreIndex == j-1 {
				ignoreIndex = -1
				continue
			}
			if b != s[j-ignoreNumber] {
				if lastIsSame {
					ignoreNumber--
					continue
				}

				return false
			}
		default:
			// if j-ignoreNumber >= lofs {
			// 	println(j, ignoreNumber, lofs)
			// 	return false
			// }
			if j-ignoreNumber >= lofs || p[j] != s[j-ignoreNumber] {
				// 还要检查后面一个字符是否是*，如果是的话，也可以不相等
				if j+1 < lofp && p[j+1] == '*' {
					ignoreIndex = j
					ignoreNumber += 2
					continue
				}
				return false
			} else if p[j] == s[j-ignoreNumber] {
				lastIsSame = true
			}
		}
	}

	return lastIndex == lofs-1
}
