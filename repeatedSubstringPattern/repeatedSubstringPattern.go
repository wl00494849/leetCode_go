package repeatedsubstringpattern

import "strings"

func repeatedSubstringPattern(s string) bool {
	var len = len(s)
	for i := 1; i <= len/2; i++ {
		if len%i == 0 && strings.Repeat(s[:i], len/i) == s {
			return true
		}
	}

	return false
}
