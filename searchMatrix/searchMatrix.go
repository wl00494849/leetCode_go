package searchmatrix

// BST
func searchMatrix(nums [][]int, target int) bool {
	var row int
	col := len(nums[0]) - 1
	for row < len(nums) && col > -1 {
		val := nums[row][col]
		if val == target {
			return true
		}
		if target > val {
			row++
		} else {
			col--
		}
	}

	return false
}
