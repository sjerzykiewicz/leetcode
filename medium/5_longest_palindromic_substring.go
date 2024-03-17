func expand(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	longest := ""
	for i := 0; i < len(s); i++ {
		oddPalindrome := expand(s, i, i)
		print(oddPalindrome)
		evenPalindrome := expand(s, i, i+1)
		print(evenPalindrome)

		if len(oddPalindrome) > len(longest) {
			longest = oddPalindrome
		}
		if len(evenPalindrome) > len(longest) {
			longest = evenPalindrome
		}
	}

	return longest
}