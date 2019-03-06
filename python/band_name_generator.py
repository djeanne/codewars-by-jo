def band_name_generator(name):
    if name[0] == name[-1]:
    	return name[0].upper() + name[1:] + name[1:]
    else:
    	return "The " + name[0].upper() + name[1:]