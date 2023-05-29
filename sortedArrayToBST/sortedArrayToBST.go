package sortedarraytobst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return recur(nums, 0, len(nums)-1)
}

func recur(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = recur(nums, left, mid-1)
	root.Right = recur(nums, mid+1, right)
	return root
}
