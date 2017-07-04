package raindrops

import "strconv"

const testVersion = 3

func Convert(in int) string {
	var str = ""
	if in%3 == 0 {

		str += "Pling"
	}
	if in%5 == 0 {
		str += "Plang"
	}
	if in%7 == 0 {
		str += "Plong"
	}
	if str != "" {
		return str
	}
	return strconv.Itoa(in)
}
