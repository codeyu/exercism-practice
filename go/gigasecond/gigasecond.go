package gigasecond

import (
	"time"
)

const testVersion = 4

func AddGigasecond(in time.Time) (got time.Time) {
	gigasecond := time.Duration(1000000000) * time.Second
	got = in.Add(gigasecond)
	return
}
