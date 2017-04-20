package clock

import (
	"fmt"
	"math"
)

const testVersion = 4

type Clock struct {
	hour   int
	minute int
}

func New(hour int, minute int) Clock {
	h, m := getAdjustedMinuteAndReturnAccumulatedHour(minute)
	h = getAdjustedHour(h + hour)
	return Clock{h, m}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {

	return New(clock.hour, clock.minute+minutes)
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
		hour = getAdjustedHour(num/60) - 1
	}
	return hour, minute
}
