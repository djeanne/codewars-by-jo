function roundToNext5(n){
	let a = n % 5;
	if(n == 0) {
		b = 0;
	} else if(a ==0){
		b = 5 * Math.floor(n / 5)
	} else {
		b = 5 * (Math.floor(n / 5) + 1)
	}
	return b
}