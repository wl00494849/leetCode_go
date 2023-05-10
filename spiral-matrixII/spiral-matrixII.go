package spiralmatrixii

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	var startRow int
	var startCol int
	endCol := n - 1
	endRow := n - 1
	total := n * n
	count := 1

	for count <= total {
		for i := startCol; i <= endCol && count <= total; i++ {
			res[startRow][i] = count
			count++
		}
		startRow++
		for i := startRow; i <= endRow && count <= total; i++ {
			res[i][endCol] = count
			count++
		}
		endCol--
		for i := endCol; i >= startCol && count <= total; i-- {
			res[endRow][i] = count
			count++
		}
		endRow--
		for i := endRow; i >= startRow && count <= total; i-- {
			res[i][startCol] = count
			count++
		}
		startCol++
	}

	return res
}
