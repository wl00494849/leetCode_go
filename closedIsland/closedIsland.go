package closedisland

func land_Dfs(row int, col int, grid [][]int, isIsland *bool) {
	if row < 0 || col < 0 || row == len(grid) || col == len(grid[0]) {
		*isIsland = false
		return
	}
	if grid[row][col] == 1 {
		return
	}
	grid[row][col] = 1

	land_Dfs(row+1, col, grid, isIsland)
	land_Dfs(row-1, col, grid, isIsland)
	land_Dfs(row, col+1, grid, isIsland)
	land_Dfs(row, col-1, grid, isIsland)

}

func closedIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	var result int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				isIsland := true
				land_Dfs(i, j, grid, &isIsland)
				if isIsland == true {
					result++
				}
			}
		}
	}
	return result
}
