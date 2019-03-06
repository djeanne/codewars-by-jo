def openOrSenior(data):
	status = []
	for member in data:
		age = member[0]
		handicap = member[1]
		if age >= 55 and handicap > 7:
			status.append('Senior')
		else:
			status.append('Open')
	return status