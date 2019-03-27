def kebabize(string):
	
	k = []
	for i, char in enumerate(string):
		if string[i].isalpha():
			if string[i] == string[i].upper():
				if i != 0:
					k.extend(("-", string[i].lower()))
				else:
					k.append(string[i].lower())
			else:
				k.append(string[i])

	if len(k) != 0:
		if not k[0].isalpha():
			k.pop(0)

	kebab = "".join(k)
	return kebab

