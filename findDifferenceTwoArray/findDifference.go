package finddifferencetwoarray

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := map[int]bool{}
	m2 := map[int]bool{}
	res := [][]int{{}, {}}
	for _, v := range nums1 {
		m1[v] = true
	}
	for _, v := range nums2 {
		m2[v] = true
	}
	for _, v := range nums2 {
		if !m1[v] {
			res[1] = append(res[1], v)
			m1[v] = true
		}
	}
	for _, v := range nums1 {
		if !m2[v] {
			res[0] = append(res[0], v)
			m2[v] = true
		}
	}
	return res
}
