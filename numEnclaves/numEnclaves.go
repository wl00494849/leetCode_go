package numenclaves

func enlavesDfs(row int, col int, grid [][]int, count *int) {
	if row < 0 || col < 0 || row == len(grid) || col == len(grid[0]) {
		*count = -(1e5)
		return
	}
	if grid[row][col] == 0 {
		return
	}

	*count++
	grid[row][col] = 0

	enlavesDfs(row-1, col, grid, count)
	enlavesDfs(row, col-1, grid, count)
	enlavesDfs(row+1, col, grid, count)
	enlavesDfs(row, col+1, grid, count)
}

func numEnclaves(grid [][]int) int {
	var total int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				var count int
				enlavesDfs(i, j, grid, &count)
				total += maxI(0, count)
			}
		}
	}
	return total
}

func maxI(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
