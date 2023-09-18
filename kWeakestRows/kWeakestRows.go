package kweakestrows

func kWeakestRows(mat [][]int, k int) []int {
	var res []int
	m := map[int]bool{}

	for j := 0; j < len(mat[0]); j++ {
		for i := 0; i < len(mat); i++ {
			if _, b := m[i]; b == false && mat[i][j] == 0 {
				res = append(res, i)
				m[i] = true
			}
			if len(res) == k {
				return res
			}
		}
	}

	if k != len(res) {
		for i := 0; i < len(mat); i++ {
			if m[i] == false {
				res = append(res, i)
			}
			if len(res) == k {
				return res
			}
		}
	}

	return res
}

// fast version
func kWeakestRows1(mat [][]int, k int) []int {
	res := []int{}
	for j := 0; j < len(mat[0]); j++ {
		for i := 0; i < len(mat); i++ {
			if mat[i][j] == 0 && ((j == 0) || (mat[i][j-1] != 0)) {
				res = append(res, i)
			}
		}
	}
	for i := 0; i < len(mat); i++ {
		if mat[i][len(mat[0])-1] == 1 {
			res = append(res, i)
		}
	}
	return res[:k]
}
