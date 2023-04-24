package widthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BFS struct {
	level int
	index int
	root  *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	var res int
	queue := []BFS{}
	var level int
	preNum := 1
	queue = append(queue, BFS{
		level: level,
		root:  root,
		index: preNum,
	})
	for i := 0; i < len(queue); i++ {
		if level < queue[i].level {
			level = queue[i].level
			preNum = queue[i].index
		}

		res = maxI(res, preNum-queue[i].index+1)

		if queue[i].root.Left != nil {
			queue = append(queue, BFS{
				level: level + 1,
				index: 2 * queue[i].index,
				root:  queue[i].root.Left,
			})
		}
		if queue[i].root.Right != nil {
			queue = append(queue, BFS{
				level: level + 1,
				index: 2*queue[i].index + 1,
				root:  queue[i].root.Right,
			})
		}
	}

	return res
}

func maxI(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
