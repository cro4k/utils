/**
 * Created by 93201 on 2017/5/27.
 */
package utils

import (
	"errors"
	"fmt"
	"time"
)

func Quarter(v ...time.Time) int {
	t := time.Now()
	if len(v) > 0 {
		t = v[0]
	}
	switch t.Month() {
	case time.January, time.February, time.March:
		return 1
	case time.April, time.May, time.June:
		return 2
	case time.July, time.August, time.September:
		return 3
	case time.October, time.November, time.December:
		return 4
	}
	return 0
}

func IsToday(t time.Time) bool {
	y, m, d := time.Now().Date()
	y1, m1, d1 := t.Date()
	return y == y1 && m == m1 && d == d1
}

func Today(args ...time.Time) time.Time {
	y, m, d := time.Now().Date()
	if len(args) > 0 {
		y, m, d = args[0].Date()
	}
	//"2006-01-02T15:04:05Z07:00"
	t, _ := time.Parse(time.RFC3339, fmt.Sprintf("%d-%02d-%02dT00:00:00+08:00", y, int(m), d))
	return t
}

func TodayEnd(args ...time.Time) time.Time {
	y, m, d := time.Now().Date()
	if len(args) > 0 {
		y, m, d = args[0].Date()
	}
	t, _ := time.Parse(time.RFC3339, fmt.Sprintf("%d-%02d-%02dT23:59:59+08:00", y, int(m), d))
	return t
}

func ParseTime(args ...interface{}) (time.Time, error) {
	if len(args) < 3 {
		return time.Time{}, errors.New("arguments not enough")
	}
	args = append(args, 0, 0, 0)
	t, err := time.Parse(time.RFC3339, fmt.Sprintf("%v-%02v-%02vT%02v:%02v:%02v+08:00", args[:6]...))
	return t, err
}

func Week() [7]time.Time {
	now := time.Now()
	res := [7]time.Time{}
	for i := 0; i < 7; i++ {
		y, m, d := now.Add(-24 * time.Hour * time.Duration(i)).Date()
		res[i], _ = ParseTime(y, int(m), d)
	}
	return res
}

func ThisWeek() (time.Time, time.Time) {
	y, m, d := time.Now().Date()
	now, _ := ParseTime(y, int(m), d)
	wd := now.Weekday()
	n := int(wd)
	if n == 0 {
		n = 7
	}
	monday := now.Add(time.Duration(1-n) * 24 * time.Hour)
	sunday := monday.Add(7*24*time.Hour - 1)
	return monday, sunday
}

func ThisMonth() (time.Time, time.Time) {
	y, m, _ := time.Now().Date()
	ed := 31
	switch m {
	case time.February:
		if y%4 == 0 && y%400 != 0 {
			ed = 29
		} else {
			ed = 28
		}
	case time.April, time.June, time.September, time.November:
		ed = 30
	}
	start, _ := ParseTime(y, int(m), 1)
	end, _ := ParseTime(y, int(m), ed, 23, 59, 59)
	return start, end
}
