package raindrops

import "strconv"

const testVersion = 3

func Convert(n int) string {

	var s string
	if n%3 == 0 {
		s += "Pling"
	}
	if n%5 == 0 {
		s += "Plang"
	}
	if n%7 == 0 {
		s += "Plong"
	}
	if len(s) != 0 {
		return s
	}
	return strconv.Itoa(n)
}
