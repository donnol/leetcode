package no5

func longestPalindrome(s string) (r string) {
	return longestPalindromeDP(s)
}

// 对于字符串 \textrm{``ababa''}“ababa”，如果我们已经知道 \textrm{``bab''}“bab” 是回文串，那么 \textrm{``ababa''}“ababa” 一定是回文串，这是因为它的首尾两个字母都是 \textrm{``a''}“a”。
//
// 根据这样的思路，我们就可以用动态规划的方法解决本题
//
// 1. 状态定义：最后一个、子问题
// 2. 转移方程：f(i, j) = f(i+1, j-1) && s[i] == s[j]，即子串为回文且元素相等
// 3. 初始条件、边界情况：
// 4. 计算顺序：由小到大
func longestPalindromeDP(s string) (r string) {
	n := len(s)
	ans := ""

	// n*n slice
	// 为什么会想到要用一个n*n的切片呢？
	// 什么时候赋值，赋什么值呢？
	// 什么时候取值，值得用途是什么呢？
	// 数组是用来记录子问题的解的，后面的问题会依赖到之前的子问题
	dp := make([][]int, n) // 一个二维切片，大小是n*n；值代表两个下标的元素是否相等
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for l := 0; l < n; l++ { // 两个字符的间隔
		for i := 0; i+l < n; i++ { //

			// 为什么j的值是i+l呢？
			// 以l位置作为分界线，i是前面部分，j是后面部分
			j := i + l

			if l == 0 { // 首次遍历，先将对角线上的值设为1
				dp[i][j] = 1
			} else if l == 1 { // 第二次遍历，如果前后元素相等，值设为1
				if s[i] == s[j] { // 相邻元素
					dp[i][j] = 1
				}
			} else { // 其它情况，如果前后元素相等，值设为dp[i+1][j-1]
				if s[i] == s[j] { // 此时间隔大于1，是否为回文的条件是中间部分是否是回文
					dp[i][j] = dp[i+1][j-1]
				}
			}

			// 如果值大于0，并且l+1大于已有子串长度，则替换子串为s[i : i+l+1]
			if dp[i][j] > 0 && l+1 > len(ans) {
				ans = s[i : i+l+1]
			}
		}
	}
	return ans
}

// n!
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
