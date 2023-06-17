package getminimumdifferencebst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(root *TreeNode, prev, min *int) {
	if root == nil {
		return
	}
	inOrder(root.Left, prev, min)
	if v := abs(*prev, root.Val); *min > v {
		*min = v
	}
	*prev = root.Val
	inOrder(root.Right, prev, min)
}

func getMinimumDifference(root *TreeNode) int {
	var min = int(1e5)
	var prev = int(1e5)
	inOrder(root, &prev, &min)
	return min
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}
