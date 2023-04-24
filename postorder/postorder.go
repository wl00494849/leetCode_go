package postorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorder(root *TreeNode, sli *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, sli)
	postorder(root.Right, sli)
	*sli = append(*sli, root.Val)
}

func main() {

}
