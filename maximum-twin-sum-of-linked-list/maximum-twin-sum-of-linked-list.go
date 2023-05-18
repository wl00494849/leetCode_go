package maximumtwinsumoflinkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

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
		fmt.Println(copyhead.Val)
		fmt.Println(prev.Val)
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
	return 0
}
