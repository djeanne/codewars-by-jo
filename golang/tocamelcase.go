package kata

import (
	"regexp"
	"strings"
	"unicode"
)

func ToCamelCase(s string) string {

	if len(s) != 0 {
		delims := regexp.MustCompile("-|_")
		tobe := delims.ReplaceAllLiteralString(s, " ")
		var firstWord string
		var newWord string

		words := strings.Split(tobe, " ")
		camel := make([]string, 0)

		if unicode.IsUpper(rune(words[0][0])) {
			firstWord = string(words[0][0]) + strings.ToLower(words[0][1:])
		} else {
			firstWord = strings.ToLower(words[0])
		}

		camel = append(camel, firstWord)

		for _, word := range words[1:] {
			newWord = strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:len(word)]))
			camel = append(camel, newWord)
		}

		camelCase := strings.Join(camel, "")

		return camelCase
	} else {
		return ""
	}
}
