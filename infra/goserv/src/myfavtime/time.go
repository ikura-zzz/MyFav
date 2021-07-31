package myfavtime

import "time"

func GetTimeString() string {
	layout := "2006-01-02 15:04:05"
	now := time.Now()
	return now.Format(layout)
}
