package kata

import "strings"

func bandNameGenerator(word string) string {
	switch {
	case word[0] == word[len(word)-1]: {
		return (strings.ToUpper(string(word[0])) + string(word[1:len(word)]) + string(word[1:len(word)]))
	}
	case strings.Contains(word, "-"): {
		h := strings.Index(word, "-")
		i := h + 1
		j := h + 2
		return "The " + (strings.ToUpper(string(word[0]))) + string(word[1:i]) + (strings.ToUpper(string(word[i])) + string(word[j:len(word)]))
	}
	default: {
		return "The " + strings.ToUpper(string(word[0])) + string(word[1:len(word)])
	}
	}
}