package kata

import "strings"

func duplicate_count(s1 string) int {
	symbols := strings.Split(strings.ToLower(s1), "")
	duplicates := make(map[string]int)
	found := make([]string, 0)

	for _, x := range symbols {
		duplicates[x] = 0
	}

	for _, y := range symbols {
		_, ok := duplicates[y]
		if ok {
			duplicates[y] += 1
		} else {
			continue
		}
	}

	for k, v := range duplicates {
		if v > 1 {
			found = append(found, k)
		}
	}

	amount := len(found)

	return amount
}
