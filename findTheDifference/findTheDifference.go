package findthedifference

func findTheDifference(s string, t string) byte {
	var i, j int
	var res byte
	for i < len(s) || j < len(t) {
		if s[i] != t[j] {

		}
		i++
		j++
	}
	return res
}
