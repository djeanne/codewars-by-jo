def reindeer(presents):
	assert(presents <= 180), "Too many presents!"
	if presents == 0:
		return 2
	elif presents > 0 and presents <= 30:
		return 3
	else:
		reindeers = float(presents) / 30
		if reindeers <= 6.00:
			if float(int(reindeers)) == reindeers:
				return int(reindeers)+2
			return int(reindeers)+3