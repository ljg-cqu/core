package utils

import (
	"strconv"
	"time"
)

func NowTimestamp13() string {
	timeString := strconv.Itoa(int(time.Now().UnixMilli()))
	return timeString
}
