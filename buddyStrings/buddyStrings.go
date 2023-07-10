package buddystrings

func buddyStrings(s, goal string) bool {
	if len(s) != len(goal) || len(s) < 2 {
		return false
	}

	if s == goal {
		m := map[byte]int{}
		for i := 0; i < len(s); i++ {
			m[s[i]]++
		}
		return len(m) != len(s)
	}

	var idx []int
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			idx = append(idx, i)
		}
	}

	if len(idx) == 2 {
		if s[idx[0]] == goal[idx[1]] && s[idx[1]] == goal[idx[0]] {
			return true
		}
	}

	return false
}
