package raindrops
import "strconv"
const testVersion = 3

func Convert(in int) string{
	var str = ""
	var has_factor = false
	if in % 3 == 0{
		has_factor = true
		str += "Pling"
	}
	if in % 5 ==0{
		has_factor = true
		str += "Plang"
	}
	if in % 7 == 0{
		has_factor = true
		str += "Plong"
	}
	if has_factor{
		return str
	}
	return strconv.Itoa(in)
}
