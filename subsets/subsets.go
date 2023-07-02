package subsets

func subsets(nums []int) [][]int {
	res := [][]int{}
	for i := 0; i <= len(nums); i++ {
		backtrack(nums, 0, i, []int{}, &res)
	}
	return res
}

func backtrack(nums []int, index int, k int, subset []int, res *[][]int) {
	if len(subset) == k {
		temp := make([]int, k)
		copy(temp, subset)
		*res = append(*res, temp)
		return
	}
	for i := index; i < len(nums); i++ {
		subset = append(subset, nums[i])
		backtrack(nums, index+1, k, subset, res)
		subset = subset[:len(subset)-1]
	}
}
