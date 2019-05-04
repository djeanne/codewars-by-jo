function abbrevName(name){
	let abbr = name.split(" ");
	return abbr[0][0].toUpperCase() + "." + abbr[1][0].toUpperCase()
}