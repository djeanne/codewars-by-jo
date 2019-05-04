def round_to_next5(n):
	a = n % 5
	if n == 0:
		b = 0
	elif a == 0:
		b = 5 * (n // 5)
	else:
		b = 5 * ((n // 5) + 1)

	return b