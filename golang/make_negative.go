package kata

func MakeNegative(x int) int {
	var y int
	if x == 0 {
		y = 0
	} else if x < 0 {
		y = x
	} else {
		y = -x
	}

	return y
}