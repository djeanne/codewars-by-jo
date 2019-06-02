package kata

import "strings"

func FindShort(s string) int {

	words := strings.Fields(s)

	var l int
	list := make([]int, 0)

	for _, w := range words {
		l = len(w)
		list = append(list, l)
	}

	min := list[0]

	for _, n := range list {
		if n < min {
			min = n
		}
	}

	return min
}
