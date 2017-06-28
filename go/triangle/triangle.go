package triangle

import "math"

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	result := NaT
	if a > 0 && b > 0 && c > 0 && !math.IsInf(a, 1) && !math.IsInf(b, 1) && !math.IsInf(c, 1) && a+b >= c && a+c >= b && b+c >= a {
		if a == b && a == c && b == c {
			result = Equ
		} else if a == b || a == c || b == c {
			result = Iso
		} else {
			result = Sca
		}

	}
	return result
}

type Kind int

const (
	NaT Kind = 0
	Equ      = 1
	Iso      = 2
	Sca      = 3
)
