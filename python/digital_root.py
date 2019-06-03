def digital_root(n):

	while n >= 10:
		n = list(str(n))
		for i, x in enumerate(n):
			n[i] = int(n[i])
		n = sum(n)

	return n