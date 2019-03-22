package kata

import (
		"strconv"
		"strings"
)

func HighAndLow(in string) string {
	slice := strings.Fields(in)
	var hal = []int{}

	for _, i := range slice {
		num, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		hal = append(hal, num)
	}

	max, min := hal[0], hal[0]

	for _, a := range hal {
		if a > max {
			max = a
		}
		if a < min {
			min = a
		}
	}

	high, low := strconv.Itoa(max), strconv.Itoa(min)

	return high + " " + low
}