package diagonalsum

func diagonalSum(mat [][]int) int {
	var sum int
	col := len(mat[0])
	var j int
	for i := 0; i < col; i++ {
		sum += mat[j][i] + mat[j][col-1-i]
		j++
	}
	if col%2 == 0 {
		sum -= mat[col/2][col/2]
	}

	return sum
}
