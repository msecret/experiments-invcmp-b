package clock

import (
	"time"
)

var Now = func() time.Time { return time.Now() }

func NowForce(unix int) {
	Now = func() time.Time { return time.Unix(int64(unix), 0) }
}
