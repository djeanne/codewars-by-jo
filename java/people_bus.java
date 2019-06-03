import java.util.ArrayList;

class Metro {

	public static int countPassengers(ArrayList<int[]> stops) {
		int passengers = 0;
		int stopsCount = stops.size();
		
		for (int i = 0; i < stopsCount; i++) {
			int gotIn = stops.get(i)[0];
			int gotOut = stops.get(i)[1];
			passengers = passengers + gotIn - gotOut;
		}

		return passengers;
	}
}
  