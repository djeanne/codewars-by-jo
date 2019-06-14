public class Kata {
	public static String createPhoneNumber(int[] numbers) {
		String phoneNum = "";
		for(int i = 0; i < 10; i++) {
			String n = String.valueOf(numbers[i]);
			phoneNum += n;
		}

	String partOne = phoneNum.substring(0, 3);
	String partTwo = phoneNum.substring(3, 6);
	String partThree = phoneNum.substring(6);

	String phoneNumber = String.format("(%s) %s-%s", partOne, partTwo, partThree);

	return phoneNumber;
	}
}