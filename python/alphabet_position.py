import string

def alphabet_position(text):
	result = []
	for letter in text:
		if letter.isalpha():
			result.append(str(string.ascii_lowercase.index(letter.lower())+1))

	abc = " ".join(result)
	abc.rstrip()
	return abc