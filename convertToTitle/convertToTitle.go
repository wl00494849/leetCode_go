package converttotitle

// ASCII A = 65
func convertToTitle(columnNumber int) string {
	var result = make([]byte, 7)
	var left = 6
	for columnNumber > 0 {
		columnNumber--
		result[left] = 65 + byte(columnNumber%26)
		columnNumber = columnNumber / 26
		left--
	}

	return string(result[left+1:])
}
