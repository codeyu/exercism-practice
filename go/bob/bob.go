package bob

import (
	"strings"
	"unicode"
)

const testVersion = 3

func Hey(input string) string {
	input = strings.TrimSpace(input)
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
	} else if strings.HasSuffix(input, "?") {
		return "Sure."
	} else if input == "" {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}
