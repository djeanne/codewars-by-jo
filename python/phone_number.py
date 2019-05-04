def create_phone_number(n):
	num = []
	for i, c in enumerate(n):
		x = str(n[i])
		num.append(x)

	part_1 = "".join(num[0:3])
	part_2 = "".join(num[3:6])
	part_3 = "".join(num[6:])
	phone_number = ("(" + part_1 + ") " + part_2 + "-" + part_3)
	return phone_number