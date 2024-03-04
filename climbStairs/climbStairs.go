package climbstairs

func climbStairs(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		c := a + b
		a, b = b, c
	}
	return b
}
