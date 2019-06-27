public class PangramChecker {
	public boolean check(String sentence){
		return sentence.toLowerCase()
			.replaceAll("[^a-z]", "")
			.replaceAll("(.)(?=.*\\1)", "")
			.length() == 26;
	}
}