package adddigits

func AddDigits(num int) int {
	var sum int
	for num > 0 {
		sum += num % 10
		num = num / 10
	}
	if sum < 10 {
		return sum
	}
	return AddDigits(sum)
}
