def up_array(arr):
	if len(arr) != 0:
		for i, char in enumerate(arr):
			if arr[i] >= 10:
				return None
			else:
				arr[i] = str(arr[i])
	
	try:
		num = list(str(int("".join(arr)) + 1))
		for i, char in enumerate(num):
			num[i] = int(num[i])
		return num
	except ValueError:
		return None

	return None