package subsets

func subsets(nums []int) [][]int {
	res := [][]int{}
	backtracking(nums, &res, 0, []int{})
	return res
}

func backtracking(nums []int, res *[][]int, index int, subset []int) {

	if index == len(nums) {
		*res = append(*res, subset)
		return
	}

	for i := index; i < len(nums); i++ {
		subset = append(subset, nums[i])
		backtracking(nums, res, index+1, subset)
		subset = subset[:len(subset)-1]
	}
}
