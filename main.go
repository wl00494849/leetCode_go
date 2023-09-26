package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{2850, 2950, 3050, 2880, 2755, 2710, 2890, 3130, 2940, 3325, 2920, 2880}
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	mean := sum / 12
	fmt.Println("Mean = ", mean)
	sort.Ints(nums)
	fmt.Println("Sort : ", nums)

	var nums1 = []float64{43.9, 42.0, 42.3, 38.9, 39.9, 32.9, 43.9, 39.8, 38.8, 32.7, 40.2, 41.5, 37.9, 39.2, 33.6, 42, 40.8, 41.8, 36.8, 44.2, 42.1, 38.2, 42, 37.3, 32, 43.2, 42.6, 39.7, 41.9, 42.4}
	var sum1 float64
	for i := 0; i < len(nums1); i++ {
		sum1 += nums1[i]
	}
	mean1 := sum1 / 30
	fmt.Println("Mean = ", mean1)
	sort.Float64s(nums1)
	fmt.Println("Sort1 =", nums1)
	fmt.Println("15:", nums1[14], "16:", nums1[15])

}
