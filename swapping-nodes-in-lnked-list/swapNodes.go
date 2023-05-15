package swappingnodesinlnkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// hint Version
func swapNodes(head *ListNode, k int) *ListNode {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	swap(&arr[k-1], &arr[len(arr)-k])

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

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
