public class FizzBuzzCuckooClock {
	public static String fizzBuzzCuckooClock(String time) {
		
		String h = time.substring(0, 2);
		String m = time.substring(3);

		String phrase;

		int hour = Integer.valueOf(h);
		int minute = Integer.valueOf(m);

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
			phrase = "tick";
		}
		return phrase;
	}
}