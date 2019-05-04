function list(names){
	let simpsons = names.map(i => Object.values(i));
	let x = simpsons.length;
	switch(x) {
		case 0:
		return ""
		break;
		case 1:
		return simpsons[0][0]
		break
		case 2:
		return simpsons[0] + " & " + simpsons[1]
		break
		default:
		return simpsons.slice(0, x-2).join(", ") + ", " + simpsons[x-2] + " & " + simpsons[x-1];
		}
}