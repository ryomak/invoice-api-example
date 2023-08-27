package time

import "time"

var Now = func() time.Time {
	return time.Now()
}

func SetFakeNow(t time.Time) {
	Now = func() time.Time {
		return t
	}
}
