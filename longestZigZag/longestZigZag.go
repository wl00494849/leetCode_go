package longestZigZag

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	var max int
	maxZigZagDFS(root.Left, 0, 0, &max)
	maxZigZagDFS(root.Right, 1, 0, &max)
	return max
}

// 0 = left ; 1 = right
func maxZigZagDFS(root *TreeNode, direction int, count int, max *int) {
	if root == nil {
		if count > *max {
			*max = count
		}
		return
	}

	if direction == 0 {
		maxZigZagDFS(root.Right, 1, count+1, max)
		maxZigZagDFS(root.Left, 0, 0, max)
	}

	if direction == 1 {
		maxZigZagDFS(root.Left, 0, count+1, max)
		maxZigZagDFS(root.Right, 1, 0, max)
	}
}
