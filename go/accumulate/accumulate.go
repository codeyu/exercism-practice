package accumulate

const testVersion = 1

func Accumulate(col []string, myFunc func(string) string) []string{
	rel := []string{}
	for _, c := range col{
		rel = append(rel, myFunc(c))
	}
	return rel
}
