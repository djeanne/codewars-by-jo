package kata

import (
	"strconv"
	"strings"
)

func DigitalRoot(n int) int {

	for n > 9 {
		str := strconv.Itoa(n)
		arr := strings.Split(str, "")

		intarr := make([]int, 0)

		for _, x := range arr {
			y, err := strconv.Atoi(x)
			intarr = append(intarr, y)

			if err != nil {
				panic(err)
			}
		}

		sum := 0

		for _, el := range intarr {
			sum += el
		}

		n = sum
	}

	return n
}
