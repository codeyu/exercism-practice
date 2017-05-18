package pangram

import (
	"regexp"
	"strings"
)

const testVersion = 1

func IsPangram(input string) bool {

	reg := regexp.MustCompile("[a-z]")
	letter := reg.FindAllString(strings.ToLower(input), -1)
	UniqueSlice(&letter)
	return len(letter) == 26
}
func UniqueSlice(slice *[]string) {
	found := make(map[string]bool)
	total := 0
	for i, val := range *slice {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*slice)[total] = (*slice)[i]
			total++
		}
	}

	*slice = (*slice)[:total]
}
