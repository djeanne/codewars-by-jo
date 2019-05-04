def high_and_low(numbers):
	numbers = numbers.split(" ")
	intnumbers = []
	for number in numbers:
		intnumber = int(number)
		intnumbers.append(intnumber)
	x = max(intnumbers)
	y = min(intnumbers)
	return (str(x) + " " + str(y))