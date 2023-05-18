package swapnodesinpairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyNode := &ListNode{}
	preNode := dummyNode
	curNode := head

	for curNode != nil && curNode.Next != nil {
		preNode.Next = curNode.Next
		curNode.Next = preNode.Next.Next
		preNode.Next.Next = curNode

		curNode = curNode.Next
		preNode = curNode.Next
	}

	return dummyNode.Next
}
