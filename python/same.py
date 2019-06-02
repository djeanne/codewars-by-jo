def comp(array1, array2):
	if array1 == None or array2 == None:
		return False
	elif len(array1) != len(array2):
		return False
	else:
		squares = []
		for x in array1:
			y = x*x
			squares.append(y)

		squares.sort()
		array2.sort()

		if squares == array2:
			return True

	return False