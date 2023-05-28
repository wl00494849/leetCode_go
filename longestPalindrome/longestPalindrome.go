package longestpalindrome

func longestPalindrome(s string) int {
	var count int
	m := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		if v, b := m[s[i]]; b == false || v == 0 {
			m[s[i]] = 1
		} else {
			m[s[i]]--
			count += 2
		}
	}
	if len(s) != count {
		count++
	}

	return count
}
