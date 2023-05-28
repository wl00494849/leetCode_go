package main

import (
	"fmt"
	sumoddlengthsubarrays "leetCode/sumOddLengthSubarrays"
)

func main() {
	nums := []int{10, 11, 12}
	res := sumoddlengthsubarrays.SumOddLengthSubarrays(nums)
	fmt.Println(res)
}
