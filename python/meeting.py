def meeting(s):
	sorted_guests = []
	guests = s.split(';')
	split_names = []
	for guest in guests:
		names = guest.split(':')
		split_names.append(names)

	formatted_names = []

	for split_name in split_names:
		first_name = split_name[0].upper()
		last_name = split_name[1].upper()

		formatted_name = []

		formatted_name.append(last_name)
		formatted_name.append(first_name)

		formatted_names.append(formatted_name)

	sorted_guests = sorted(formatted_names, key = lambda x: (x[0], x[1]))
	
	new_list = []
	for sorted_guest in sorted_guests:
		new_guest = "(" + ", ".join(sorted_guest) + ")"
		new_list.append(new_guest)
	
	guest_list = ''.join(new_list)
	return guest_list




