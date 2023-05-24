package main

import (
	"fmt"
	topkfrequent "leetCode/topKFrequent"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	res := topkfrequent.TopKFrequent(nums, 2)
	fmt.Println(res)
}
