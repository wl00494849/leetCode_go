package mergealternately

func mergeAlternately(word1 string, word2 string) string {
	var i int
	var j int
	var res string
	lenWord1 := len(word1)
	lenWord2 := len(word2)
	for i < lenWord1 || j < lenWord2 {
		if i < lenWord1 {
			res += string(word1[i])
			i++
		}
		if j < lenWord2 {
			res += string(word2[j])
			j++
		}
	}
	return res
}
