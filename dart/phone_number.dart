String createPhoneNumber(List numbers) {
	
	List<String> phoneNumber = [];
	for(var number in numbers) {
		String char = number.toString();
		phoneNumber.add(char);
	}

	String partOne = phoneNumber.sublist(0, 3).join();
	String partTwo = phoneNumber.sublist(3, 6).join();
	String partThree = phoneNumber.sublist(6).join();
	String opPar = "(";
	String clPar = ") ";
	String hyp = "-";

	String phNum = opPar + partOne + clPar + partTwo + hyp + partThree;
	return phNum;
}