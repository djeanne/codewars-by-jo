package kata

import "math"

func EquableTriangle(a, b, c int) bool {
	af, bf, cf := float64(a), float64(b), float64(c)
	perimeter := af + bf + cf
	s := perimeter / 2
	area := math.Sqrt(s * ((s - af) * (s - bf) * (s - cf)))

	if perimeter == area {
		return true
	} else {
		return false
	}
}