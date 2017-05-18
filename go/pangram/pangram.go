package pangram

import (
	"regexp"
	"strings"
)

const testVersion = 1

func IsPangram3(input string) bool {
	// Lowercase everything for ease of checking.
	input = strings.ToLower(input)

	// List of all letters.
	letters := "abcdefghijklmnopqrstuvwxyz"

	// If we find any letters that aren't included in the string, then we'll
	// return false.
	allIncluded := true
	for _, l := range letters {
		if !strings.Contains(input, string(l)) {
			allIncluded = false
		}
	}

	return allIncluded
}
func IsPangram2(input string) bool {
	m := make(map[byte]bool)
	letters := strings.ToLower(input)
	for i := 0; i < len(letters); i++ {
		if letters[i] >= 'a' && letters[i] <= 'z' {
			m[letters[i]] = true
		}
	}
	return len(m) == 26
}
func IsPangram(orig string) bool {
	//create map of alphabet letters
	ab := "abcdefghijklmnopqrstuvwxyz"
	orig = strings.ToLower(orig)
	for _, val := range ab {
		if !strings.ContainsRune(orig, val) {
			return false
		}
	}
	return true
}
func IsPangram_self(input string) bool {

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
