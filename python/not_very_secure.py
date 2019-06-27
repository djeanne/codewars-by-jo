def alphanumeric(string):
	if len(string) < 1:
		return False
	return string.isalnum()