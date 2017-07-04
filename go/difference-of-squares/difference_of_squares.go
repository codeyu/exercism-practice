package diffsquares

const testVersion = 1

func SquareOfSums(firstN int) int {
	sum := 0
	for next := firstN; next > 0; next-- {
		sum += next
	}
	return sum * sum
}

func SumOfSquares(firstN int) int {
	sum := 0
	for next := firstN; next > 0; next-- {
		sum += next * next
	}
	return sum
}

func Difference(firstN int) int {
	return SquareOfSums(firstN) - SumOfSquares(firstN)
}
