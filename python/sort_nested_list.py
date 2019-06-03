def sort_nested_list(A):

	def chunk(lst, cap):
		for i in range(0, len(lst), cap):
			yield lst[i:i+cap]

	innermost = []

	for B in A:
		for C in B:
			for D in C:
				innermost.append(D)
		global m
		m = len(C)
	global z
	z = len(B)

	innermost.sort()
	
	try:
		middle = list(chunk(innermost, m))
		result = list(chunk(middle, z))
	except ValueError:
		return A

	return result