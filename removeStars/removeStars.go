package removestars

func removeStars(s string) string {
	var temp string
	for i := len(s); i > 0; i-- {
		if s[i] == '*' {
			i--
			temp += string(s[i])
		} else {
			temp += string(s[i])
		}
	}
	var res string
	for i := len(temp); i > 0; i-- {
		res += string(temp[i])
	}
	return res
}
