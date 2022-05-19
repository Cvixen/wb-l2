package cur_time

import (
	"github.com/beevik/ntp"
	"time"
)

func GetTime() (time.Time, error) {
	currentTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return time.Time{}, err
	}
	return currentTime, nil
}
