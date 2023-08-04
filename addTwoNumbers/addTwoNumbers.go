package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1 []int
	var stack2 []int

	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	var carry int
	var result *ListNode

	for len(stack1) != 0 || len(stack2) != 0 {
		var sum int
		if len(stack1) != 0 {
			sum += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) != 0 {
			sum += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}

		node := &ListNode{}
		sum = sum + carry
		node.Val = sum % 10
		carry = sum / 10
		node.Next = result
		result = node
	}

	if carry != 0 {
		node := &ListNode{Val: 1}
		node.Next = result
		result = node
	}

	return result
}
