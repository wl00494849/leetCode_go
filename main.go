package main

import (
	"fmt"
	"leetCode/subsets"
)

func main() {
	nums := []int{1, 2, 3}
	res := subsets.Subsets(nums)
	fmt.Println(res)
}
