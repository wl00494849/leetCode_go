package validatestacksequences

func validateStackSequences(pushed []int, popped []int) bool {
	var k int
	j := len(popped)
	stack := []int{}
	for i := 0; i < len(pushed); i++ {
		if pushed[i] != popped[k] {
			stack = append(stack, pushed[i])
		} else {
			k++
			j--
			for len(stack) > 0 && stack[len(stack)-1] == popped[k] {
				k++
				j--
				stack = stack[:len(stack)-1]
			}
		}
	}
	if j != 0 {
		return false
	}

	return true
}
