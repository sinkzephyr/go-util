package util

import (
	"fmt"
	"time"
)

const TimeLayout = "2006-1-2"

//fmt.Println("first date:", FirstDateOfWeek())
//fmt.Println("first date next week:", FirstDateOfNextWeek())

//var cdate = "2020-2-4"
//fmt.Println("first date custom date:", FirstDateOfWeekByDate(cdate))
//fmt.Println("first date next week by custom date:", FirstDateOfNextWeekByDate(cdate))
func FirstDateOfWeek() (weekMonday string) {
	return GetFirstDateOfWeek(false, "").Format(TimeLayout)
}
func FirstDateOfNextWeek() (weekMonday string) {
	return GetFirstDateOfWeek(false, "").AddDate(0, 0, 7).Format(TimeLayout)
}

func FirstDateOfWeekByDate(dateStr string) (weekMonday string) {
	return GetFirstDateOfWeek(true, dateStr).Format(TimeLayout)
}

func FirstDateOfNextWeekByDate(dateStr string) (weekMonday string) {
	return GetFirstDateOfWeek(true, dateStr).AddDate(0, 0, 7).Format(TimeLayout)
}


/**
获取本周周一的日期
*/
func GetFirstDateOfWeek(customDate bool, dateStr string) (weekMonday time.Time) {
	var now time.Time
	if customDate == true {
		if dateStr == "" {
			fmt.Println("customDate error")
			return
		}
		var parseErr error
		now, parseErr = time.ParseInLocation(TimeLayout, dateStr, time.Local)
		if parseErr != nil {
			fmt.Println("Time parse error: ", dateStr, parseErr)
			return
		}
	} else {
		now = time.Now()
	}

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStartDate
}
