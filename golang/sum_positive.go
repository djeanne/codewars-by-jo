package kata

func PositiveSum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	} else {

		var positives []int

		for _, num := range numbers {
			if num > 0 {
				positives = append(positives, num)
			}
		}

		sum := 0

		for _, n := range positives {
			sum += n
		}

		return sum
	}
}
