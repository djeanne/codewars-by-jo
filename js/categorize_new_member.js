function openOrSenior(data){
	let status = [];
	for(let i = 0; i < data.length; i++) {
		let age = data[i][0]
		let handicap = data[i][1]
		if(age >= 55 && handicap > 7) {
			status.push("Senior");
		} else {
			status.push("Open");
		}
	}
	return status
}