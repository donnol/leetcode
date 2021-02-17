package no5

func longestPalindrome(s string) (r string) {
	return longestPalindromeDP(s)
}

// 对于字符串 \textrm{``ababa''}“ababa”，如果我们已经知道 \textrm{``bab''}“bab” 是回文串，那么 \textrm{``ababa''}“ababa” 一定是回文串，这是因为它的首尾两个字母都是 \textrm{``a''}“a”。
//
// 根据这样的思路，我们就可以用动态规划的方法解决本题
func longestPalindromeDP(s string) (r string) {
	n := len(s)
	ans := ""

	// n*n slice
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for l := 0; l < n; l++ { // 遍历字符串
		for i := 0; i+l < n; i++ { // l向前进的同时，i的上限值变小
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l+1 > len(ans) {
				ans = s[i : i+l+1]
			}
		}
	}
	return ans
}

func longestPalindromeByAll(s string) (r string) {
	l := len(s)
	for i := 0; i < l; i++ {
		for j := i; j <= l; j++ {
			// 滑动窗口
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
