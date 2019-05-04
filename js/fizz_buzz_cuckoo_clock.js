function fizzBuzzCuckooClock(time) {
	let h = time.slice(0, 2);
	let m = time.slice(3);
	let hour = parseInt(h, 10);
	let minute = parseInt(m, 10);

	let phrase;

	if(minute == 30) {
		phrase = "Cuckoo";
	} else if(minute == 0) {
		if(hour == 0) {
			phrase = "Cuckoo ".repeat(12).trim();
		} else if(hour > 0 && hour <= 12) {
			phrase = "Cuckoo ".repeat(hour).trim();
		} else {
			phrase = "Cuckoo ".repeat(hour - 12).trim();
		}
	} else if(minute % 3 == 0 && minute % 5 == 0) {
		phrase = "Fizz Buzz";
	} else if(minute % 3 == 0) {
		phrase = "Fizz";
	} else if(minute % 5 == 0) {
		phrase = "Buzz";
	} else {
		phrase = "tick"
	}

	return phrase
}