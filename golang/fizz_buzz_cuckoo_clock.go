package kata

import (
	"fmt"
	"strings"
	"strconv"
)

func FizzBuzzCuckooClock(time string) string {
	var phrase string
	hour := time[0:2]
	minute := time[3:]

	h, err := strconv.Atoi(hour)
	if err != nil {
		fmt.Println("Error!")
	}
	m, err := strconv.Atoi(minute)
	if err != nil {
		fmt.Println("Error!")
	}

	switch {
	case m == 30:
		phrase = "Cuckoo"
	case m == 0 && h > 0 && h <= 12:
		phrase = strings.TrimSuffix((strings.Repeat("Cuckoo ", h)), " ")
	case m == 0 && h > 12:
		phrase = strings.TrimSuffix((strings.Repeat("Cuckoo ", (h-12))), " ")
	case m == 0 && h == 0:
		phrase = strings.TrimSuffix((strings.Repeat("Cuckoo ", 12)), " ")
	case m % 3 == 0 && m % 5 == 0 && m != 0:
		phrase = "Fizz Buzz"
	case m % 3 == 0 && m != 0:
		phrase = "Fizz"
	case m % 5 == 0 && m != 0:
		phrase = "Buzz"
	default:
		phrase = "tick"	
	}

	return phrase
}
