package average

func average(salary []int) float64 {
	var max int
	min := 100000
	var sum int
	for _, v := range salary {
		sum += v
		if v > max {
			max = v
		}
		if v <= min {
			min = v
		}
	}
	sum -= max + min
	return float64(sum) / float64(len(salary)-2)
}
