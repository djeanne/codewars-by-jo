def mxdiflg(a1, a2):
	if len(a1) == 0 or len(a2) == 0:
		return -1
	else:
		x = []
		for item in a1:
			x.append(len(item))
		y = []	
		for item in a2:
			y.append(len(item))
		z = []
		for a in x:
			for b in y:
				z.append(abs(a - b))
		
		return max(z)