package spiralorder

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	col := len(matrix[0])
	row := len(matrix)
	total := col * row
	var count int
	startCol := 0
	startRow := 0
	endCol := col - 1
	endRow := row - 1

	for count < total {
		for i := startCol; i <= endCol && count < total; i++ {
			res = append(res, matrix[startRow][i])
			count++
		}
		startRow++
		for i := startRow; i <= endRow && count < total; i++ {
			res = append(res, matrix[i][endCol])
			count++
		}
		endCol--
		for i := endCol; i >= startCol && count < total; i-- {
			res = append(res, matrix[endRow][i])
			count++
		}
		endRow--
		for i := endRow; i >= startRow && count < total; i-- {
			res = append(res, matrix[i][startCol])
			count++
		}
		startCol++
	}

	return res
}
