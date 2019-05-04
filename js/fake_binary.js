function fakeBin(x){
	let fbin = "";
	for(let i = 0; i < x.length; i++) {
		let intchar = parseInt(x[i], 10);
		let strchar = "";

		if (intchar < 5) {
			strchar = "0";
		} else {
			strchar = "1";
		}
		fbin = fbin + strchar;
	}
	return fbin
}