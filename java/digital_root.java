public class DRoot {
	public static int digital_root(int n) {

		while (n > 9) {
			int sum = 0;
			String stringN = String.valueOf(n);
			String[] arrN = stringN.split("");

			
			for(int i = 0; i < arrN.length; i++) {
				int digit = Integer.valueOf(arrN[i]);
				sum += digit;
			}

			n = sum;
		}

		return n;
	}
}