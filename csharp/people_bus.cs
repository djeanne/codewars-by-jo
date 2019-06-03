using System;
using System.Collections.Generic;

public class Kata
	{
		public static int Number(List<int[]> peopleListInOut)
		{
			int passengers = 0;
			
			foreach (var el in peopleListInOut)
			{
				int gotIn = el[0];
				int gotOut = el[1];
				passengers = passengers + gotIn - gotOut;
			}

			return passengers;
		}
	}