package hash

import "math"

func str2int(s string) int {
	bb := make([]byte, len(s))
	copy(bb, s)

	res := 0

	for _, b := range bb {
		res += int(b)
	}

	return res
}

// Takes any string and returns int from 0 to sz.
func HashString(s string, c float64, sz int) int {
	k := float64(str2int(s))
	fract := ((k * c) - math.Floor(k*c))
	return int(math.Floor(float64(sz) * fract))
}
