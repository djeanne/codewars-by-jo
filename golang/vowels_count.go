package kata

import "regexp"

func GetCount(str string) (count int) {
	vowels := regexp.MustCompile("a|e|i|o|u")
	amount := vowels.FindAllStringIndex(str, -1)
	count = len(amount) 
	return count
}