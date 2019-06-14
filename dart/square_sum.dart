int squareSum(List numbers) {
	int squareSum = 0;
	for(var number in numbers) {
		number = number * number;
		squareSum += number;
	}
	return squareSum;
}