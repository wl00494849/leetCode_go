package findduplicate

import "sort"

// hash map 35.8% low runtime high memory
func findDuplicate(nums []int) int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if v := m[nums[i]]; v == 2 {
			return nums[i]
		}
	}
	return 0
}

// sort 6.85% high runtime low memory
// quick sort
func findDuplicate1(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return 0
}
