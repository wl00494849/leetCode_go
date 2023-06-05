package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	return recure(root)
}

func recure(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right = recure(root.Right)
	root.Left = recure(root.Left)
	root.Left, root.Right = root.Right, root.Left
	return root
}
