package kata

import "sort"

func TwoOldestAges(ages []int) [2]int {
	var a [2]int
	sort.Sort(sort.Reverse(sort.IntSlice(ages)))
	a[0], a[1] = ages[1], ages[0]
	return a
}
