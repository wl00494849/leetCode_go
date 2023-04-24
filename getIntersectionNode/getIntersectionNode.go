package getintersectionnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ptrA := headA
	ptrB := headB

	for ptrA != nil && ptrB != nil {
		ptrA = ptrA.Next
		ptrB = ptrB.Next
	}

	for ptrA != nil {
		headA = headA.Next
		ptrA = ptrA.Next
	}

	for ptrB != nil {
		headB = headB.Next
		ptrB = ptrB.Next
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
