package leap

const testVersion = 3

func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.
	if mod := year % 100; mod != 0 {
		return year%4 == 0
	}
	return year%400 == 0

}
