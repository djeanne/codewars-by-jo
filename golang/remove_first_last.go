package kata

func RemoveChar(word string) string {
	switch {
	case len(word) < 2:
		return word
	default:
		return word[1:len(word)-1]
	}
}