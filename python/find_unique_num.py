from collections import Counter

def find_uniq(arr):
	unique = [k for k, v in Counter(arr).items() if v == 1]
	n = unique[0]
	return n