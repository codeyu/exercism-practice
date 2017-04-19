// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import (
	"fmt"
	"math"
)

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock struct {
	Hour   int
	Minute int
}

func New(hour int, minute int) *Clock {
	var c = Clock{0, 0}
	c.Hour = hour
	c.Minute = minute
	return &c
}

func (clock Clock) String() string {
	hour, minute := getAdjustedMinuteAndReturnAccumulatedHour(clock.Minute)
	hour = getAdjustedHour(clock.Hour + hour)
	return fmt.Sprintf("%02d:%02d", hour, minute)
}

func (clock *Clock) Add(minutes int) Clock {
	clock.Minute += clock.Minute + minutes
	return *clock
}

func getAdjustedHour(num int) int {
	hour := 0
	switch {
	case num >= 0 && num < 24:
		hour = num

	case num >= 24:
		hour = int(math.Mod(float64(num), 24))

	case num < 0:
		hour = 24 - int(math.Mod(math.Abs(float64(num)), 24))
	}
	return hour
}

func getAdjustedMinuteAndReturnAccumulatedHour(num int) (int, int) {
	hour, minute := 0, 0
	switch {
	case num >= 0 && num < 60:
		minute = num

	case num >= 60:
		hour = getAdjustedHour(num / 60)
		minute = int(math.Mod(float64(num), 60))

	case num < 0:
		minute = 60 - int(math.Mod(math.Abs(float64(num)), 60))
		hour = getAdjustedHour(num / 60)
	}
	return hour, minute
}
