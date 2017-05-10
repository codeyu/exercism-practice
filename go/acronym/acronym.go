package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 2

func Abbreviate(phrase string) (acr string) {
	acronym := []string{}
	re := regexp.MustCompile("[A-Z]+[a-z]*|[a-z]+")
	words := re.FindAllString(phrase, -1)
	for _, w := range words {
		acronym = append(acronym, w[0:1])
	}
	acr = strings.ToUpper(strings.Join(acronym, ""))
	return
}
