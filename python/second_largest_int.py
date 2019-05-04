def find_2nd_largest(arr):
	new = [x for x in arr if type(x) == int]
	new.sort(reverse = True)
	new.pop(0)
	if new[0] == new[1]:
		return None
	return new[0]