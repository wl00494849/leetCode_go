package subsets

func Subsets(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})
	backtracking(nums, &res, 1)
	return res
}

func backtracking(nums []int, res *[][]int, count int) {

	if count > len(nums) {
		return
	}

	var start int

	for i := 0; i < len(nums); i = i + count {
		var arr []int
		slide := start
		for j := 0; j < count; j++ {
			arr = append(arr, nums[slide])
			slide++
		}
		start++
		*res = append(*res, arr)
	}
	backtracking(nums, res, count+1)
}
