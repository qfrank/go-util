package goutil

import (
	"time"
	"strconv"
)

func Convert2Time(timestampInMills int64) *time.Time {
	if timestampInMills <= 0 {
		return nil
	}
	secs := timestampInMills / 1000
	s := strconv.FormatInt(timestampInMills, 10)
	var t int64 = 0
	if len(s) >= 3 {
		t, _ = strconv.ParseInt(s[len(s)-3:], 10, 64)
	} else {
		t, _ = strconv.ParseInt(s, 10, 64)
	}
	nsec := t * 1000000
	tt := time.Unix(secs, nsec)
	return &tt
}
