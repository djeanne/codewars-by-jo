package kata

import "strings"

func FindShort(s string) int {
	words := strings.Split(s, " ")
	short := len(words[0])

	for _, word := range words {
		if len(word) < short {
			short = len(word)
		}
	}

	return short
}