package diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var result int
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, result)
	right := dfs(root.Right, result)
	*result = max(*result, left+right)
	return 1 + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
