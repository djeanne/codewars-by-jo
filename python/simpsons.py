def namelist(names):
	simpsons = [name["name"] for name in names]

	if len(simpsons) == 0:
		return ""
	elif len(simpsons) == 1:
		return str((simpsons)[0])
	elif len(simpsons) == 2:
		return str((simpsons)[0]) + " & " + str((simpsons)[1])
	else:
		return ", ".join(simpsons[0:-2]) + ", " + str((simpsons)[-2]) + " & " + str((simpsons)[-1])