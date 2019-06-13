public class Maskify {
	public static String maskify(String str) {
		//throw new UnsupportedOperationException("Unimplemented");
		int len = str.length();
		if(len < 5) {
			return str;
		} else {
			String hashes = str.substring((len - 4));
			String newString = "";
			String hash = "#";
			for(int i = 0; i < len - 4; i++) {
				newString += hash;
			}
			String str1 = newString.concat(hashes);
			return str1;
		}
	}
}