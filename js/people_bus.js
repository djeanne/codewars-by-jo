var number = function(busStops){
	let passengers = 0;
	for (let i = 0; i < busStops.length; i++) {
		let gotIn = busStops[i][0];
		let gotOut = busStops[i][1];
		passengers = passengers + gotIn - gotOut;
		}
	return passengers
}