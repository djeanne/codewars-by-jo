import itertools

def permutations(string):
	mut = []
	for x in itertools.permutations(string):
		y = ''.join(x)
		mut.append(y)
	mut = list(dict.fromkeys(mut))
	return mut

