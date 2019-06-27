int find(List integers) {
	if (integers[0] % 2 != integers[1] % 2) {
		if (integers[2] % 2 == integers[1] % 2) {
			return integers[0];
		} else {
			return integers[1];
		}
	} else {
		for (var integer in integers) {
			if (integer % 2 != integers[0] % 2) {
				return integer;
			}
		}
	}
}