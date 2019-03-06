def fake_bin(x):
	fbin = ""
	for char in x:
		intchar = int(char)
		if intchar < 5:
			strchar = "0"
		else:
			strchar = "1"
		fbin += strchar
	return fbin