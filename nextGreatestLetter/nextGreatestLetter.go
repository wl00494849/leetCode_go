package nextgreatestletter

func nextGreatestLetter(letters []byte, target byte) byte {
	var min byte = 255
	for i := 0; i < len(letters); i++ {
		if letters[i] > target {
			if letters[i] < min {
				min = letters[i]
			}
		}
	}

	if min == 255 {
		return letters[0]
	}

	return min
}
