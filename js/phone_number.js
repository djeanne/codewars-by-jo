function createPhoneNumber(numbers){
	let partOne = numbers.slice(0, 3).join("");
	let partTwo = numbers.slice(3, 6).join("");
	let partThree = numbers.slice(6).join("");
	let phoneNumber = "(".concat(partOne, ") ", partTwo, "-", partThree)
	return phoneNumber
}