def first_non_repeating_letter(string):
	new = string.lower()
	sort_s = []
	count = {}
	for letter in new:
		if letter in count:
			count[letter] += 1
		else:
			count[letter] = 1 
		sort_s.append(letter)


	for letter in sort_s:
		if count[letter] == 1:
			for i, (a, b) in enumerate(zip(string, new)):
				if new[i] == letter:
					if string[i] == new[i]:
						return letter
					else:
						return (letter.upper())
	return ""