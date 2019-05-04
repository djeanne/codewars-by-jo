def abbrevName(name):
	abbr = name.split()
	initials = abbr[0][0].upper() + "." + abbr[1][0].upper()
	return initials