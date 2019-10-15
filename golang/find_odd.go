package kata

func FindOdd(seq []int) int {

	var odd int

	store := make(map[int]int)

	for _, y := range seq {
		store[y] = 0
	}
	
	for _, x := range seq {
		_, ok := store[x]
		if ok {
			store[x] += 1
		} else {
			continue
		}
	}

	for k, v := range store {
		if v % 2 != 0 {
			odd = k
		}
	}

	return odd
}