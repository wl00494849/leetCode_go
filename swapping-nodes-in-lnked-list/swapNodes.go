package swappingnodesinlnkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// swap Val
func swapNodes1(head *ListNode, k int) *ListNode {
	left := head
	count := 1
	for count < k {
		count++
		left = left.Next
	}
	right := head
	for node := left; node.Next != nil; node = node.Next {
		right = right.Next
	}

	left.Val, right.Val = right.Val, left.Val

	return head
}

// hint Version
func swapNodes(head *ListNode, k int) *ListNode {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	swap(&arr[k-1], &arr[len(arr)-k])

	newHead := ListNode{}
	linkList := &newHead

	for i := 0; i < len(arr); i++ {
		linkList.Next = &ListNode{
			Val: arr[i],
		}
		linkList = linkList.Next
	}

	return newHead.Next
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
