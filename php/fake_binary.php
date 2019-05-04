<?

function fake_bin(string $s): string {
	$sbin = str_split($s);
	$fbin = "";
	$slength = count($sbin);
	for ($i = 0; $i < $slength; $i++) {
		if (intval($sbin[$i]) < 5) {
			$fbin = $fbin . "0";
		} else {
			$fbin = $fbin . "1";
		}
	}
	return $fbin;
}