package isvalid

func isValid(s string) bool {
	stack := []rune{}
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		}
		if v == ')' {
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-2]
			} else {
				return false
			}
		}
		if v == '}' {
			if stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-2]
			} else {
				return false
			}
		}
		if v == ']' {
			if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-2]
			} else {
				return false
			}
		}
	}
	return true
}
