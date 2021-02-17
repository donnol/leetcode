package no5

import "testing"

func TestLongestPalindrome(t *testing.T) {
	for _, cas := range []struct {
		name string
		in   string
		want string
	}{
		{"1", "babad", "bab"},
		{"2", "cbbd", "bb"},
		{"3", "a", "a"},
		{"4", "ac", "a"},
		{"5", "abcda", "a"},
		{"6", "bb", "bb"},
		{"7", "ccd", "cc"},
		{"8", "aaaa", "aaaa"},
		{"9", "abcba", "abcba"},
		{"10", "aaabaaa", "aaabaaa"},
		{"11", "aaabaaaa", "aaabaaa"},
		{"12", "tattarrattat", "tattarrattat"},
		{"13", "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth", "ranynar"},
	} {
		r := longestPalindrome(cas.in)
		if r != cas.want {
			t.Fatalf("No.%s: bad result, %v != %v\n", cas.name, r, cas.want)
		}
	}
}

func BenchmarkLongestPalindrome(b *testing.B) {
	in := "abcddfdfdfdfdfdfdfddfd"

	for i := 0; i < b.N; i++ {
		longestPalindrome(in)
	}

	// BenchmarkLongestPalindrome-8   	 1606592	       713 ns/op	     352 B/op	       1 allocs/op
}
