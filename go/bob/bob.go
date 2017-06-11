package bob

import (
	"strings"
	"unicode"
)

const testVersion = 3

func Hay(input string) string {
	isAllUpper := false
	for _, ch := range input {
		if unicode.IsLetter(ch) {
			if unicode.IsUpper(ch) {
				isAllUpper = true
			} else {
				isAllUpper = false
				break
			}
		}
		continue
	}
	if isAllUpper {
		return "Whoa, chill out!"
	} else if input[len(input)-1] == '?' {
		return "Sure."
	} else if len(strings.TrimSpace(input)) == 0 {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}
