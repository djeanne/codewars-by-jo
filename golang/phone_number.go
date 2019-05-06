package kata

import (
		"strconv"
		"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	digits := make([]string, 0)
	for _, number := range numbers {
		d := uint64(number)
		digit := strconv.FormatUint(d, 10)
		digits = append(digits, digit)
	}
	num := strings.Join(digits, "")
	phone := "(" + string(num[0:3]) + ") " + string(num[3:6]) + "-" + string(num[6:len(num)])
	return phone
}