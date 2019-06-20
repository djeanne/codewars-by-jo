package kata

import "strings"

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SpinWords(str string) string {

	var err error

	words := strings.Split(str, " ")

	for i, word := range words {
		if len(word) > 4 {
			words[i] = reverse(word)

			if err != nil {
				panic(err)
			}
		}
	}

	spinned := strings.Join(words, " ")

	return spinned
}
