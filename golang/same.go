package kata

import (
	"reflect"
	"sort"
)

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil {
		return false
	} else if len(array1) != len(array2) {
		return false
	} else {
		squared := make([]int, 0)
		var s int

		for _, v := range array1 {
			s = v * v
			squared = append(squared, s)
		}

		sort.Ints(array2)
		sort.Ints(squared)

		answer := reflect.DeepEqual(array2, squared)

		return answer
	}
}
