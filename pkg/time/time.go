package time

import (
	"log"
	"time"
)

var TehranLoc *time.Location

func init() {
	l, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		log.Fatal(err)
	}
	TehranLoc = l
}

func AddMinutes(minute uint, isTehran bool) time.Time {
	now := time.Now()
	if isTehran {
		now = now.In(TehranLoc)
	}
	return now.Add(time.Minute * time.Duration(minute))
}
