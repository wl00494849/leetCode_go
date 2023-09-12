package generate

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	if numRows == 1 {
		return [][]int{{1}}
	}

	res := [][]int{{1}, {1, 1}}
	for i := 2; i < numRows; i++ {
		var temp []int
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				temp = append(temp, 1)
			} else {
				temp = append(temp, res[i-1][j]+res[i-1][j-1])
			}
		}
		res = append(res, temp)
	}
	return res
}
