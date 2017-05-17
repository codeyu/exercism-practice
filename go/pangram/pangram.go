package pangram
import (
	"regexp"
	"strings"
)
const testVersion = 1

func IsPangram(input string) bool｛
	reg := regexp.MustCompile("[a-zA-Z]")
	letter := reg.FindAllString(input, -1)
｝
