package issubsequence

func isSubsequence(s string, t string) bool {
	var i, j int
	for i < len(t) && j < len(s) {
		if s[j] == t[i] {
			j++
		}
		i++
	}

	if j == len(s) {
		return true
	}

	return false
}
