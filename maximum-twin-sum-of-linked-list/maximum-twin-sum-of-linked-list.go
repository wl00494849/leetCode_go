package maximumtwinsumoflinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// first version
func pairSum(head *ListNode) int {
	var prev *ListNode
	var count int
	var result int
	copyList := &ListNode{}
	copyhead := copyList
	cur := head

	for cur != nil {
		copyList.Next = &ListNode{
			Val: cur.Val,
		}
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
		copyList = copyList.Next
		count++
	}
	copyhead = copyhead.Next
	for i := 0; i < count/2; i++ {
		result = max(copyhead.Val+prev.Val, result)
		copyhead = copyhead.Next
		prev = prev.Next
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// optimal version
func pairSum1(head *ListNode) int {
	var prev *ListNode
	var result int
	slow := head
	fast := head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	for slow != nil {
		temp := slow.Next
		slow.Next = prev
		prev = slow
		slow = temp
	}

	for prev != nil {
		result = max(head.Val+prev.Val, result)
		head = head.Next
		prev = prev.Next
	}

	return result
}
