package checkifitisastraightline

func checkStraightLine(coordinates [][]int) bool {
	deltaX := coordinates[0][0] - coordinates[1][0]
	deltay := coordinates[0][1] - coordinates[1][1]
	for i := 2; i < len(coordinates); i++ {
		x := coordinates[i-1][0] - coordinates[i][0]
		y := coordinates[i-1][1] - coordinates[i][1]
		if deltaX*y != deltay*x {
			return false
		}
	}
	return true
}
