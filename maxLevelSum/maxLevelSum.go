package maxlevelsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func maxLevelSum(root *TreeNode) int {
	var queue = []*TreeNode{}
	var maxLevel int
	var maxVal = int(-1e5)
	var level int
	queue = append(queue, root)

	for len(queue) != 0 {
		k := len(queue)
		var sum int
		level++
		for i := 0; i < k; i++ {
			sum += queue[0].Val
			//pop
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		if sum > maxVal {
			maxVal = sum
			maxLevel = level
		}
	}

	return maxLevel
}
