def number(bus_stops):
	passengers = 0
	for bus_stop in bus_stops:
		got_in = bus_stop[0]
		got_out = bus_stop[1]
		passengers = passengers + got_in - got_out
	return passengers