package partitionstring

func partitionString(s string) int {
	var partition int
	mp := map[byte]bool{}
	for i := 0; i < len(s); i++ {
		if mp[s[i]] {
			mp = make(map[byte]bool)
			partition++
		}
		mp[s[i]] = true
	}
	return partition
}
