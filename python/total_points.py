def points(games):
	sumpoints = []
	for game in games:
		count = game.split()
		count1 = count[0].split(':')
		x = int(count1[0])
		y = int(count1[1])
		total = 0
		if x > y:
			total += 3
		elif x < y:
			total += 0
		elif x == y:
			total += 1
		sumpoints.append(total)
	return sum(sumpoints)