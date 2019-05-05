<?

function fizzBuzzCuckooClock($time) {
	$hour = intval(substr($time, 0, 2));
	$minute  = intval(substr($time, 3));

	if ($minute == 30) {
		$phrase = "Cuckoo";
	} elseif ($minute == 0) {
		if($hour == 0 or $hour == 12) {
			$phrase = rtrim(str_repeat("Cuckoo ", 12), " ");
		} elseif ($hour > 0 and $hour < 12) {
			$phrase = rtrim(str_repeat("Cuckoo ", $hour), " ");
		} else {
			$phrase = rtrim(str_repeat("Cuckoo ", $hour - 12), " ");
		}
	} elseif ($minute % 3 == 0 and $minute % 5 == 0) {
		$phrase = "Fizz Buzz";
	} elseif ($minute % 3 == 0) {
		$phrase = "Fizz";
	} elseif ($minute % 5 == 0) {
		$phrase = "Buzz";
	} else {
		$phrase = "tick";
	}
	return $phrase;
}