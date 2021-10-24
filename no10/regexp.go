package no10

func isMatch(s string, p string) bool { //nolint
	return isMatchByLeetcode(s, p)
}

// 使用动态规划
func isMatchByLeetcode(s string, p string) bool {
	m, n := len(s), len(p)

	// 判断函数：输入s和p的下标，判断对应字符是否一样
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	// 状态二维列表
	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	// 状态初始值
	f[0][0] = true
	// 状态转移过程：从左到右
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' { // 如果前一个字符为*
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}

func isMatchByMyself(s string, p string) bool { //nolint
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
