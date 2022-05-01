package chapter01

import (
	"fmt"
	"time"
)

func NextMonth(t time.Time) time.Time {
	year1, month1, day := t.Date()
	first := time.Date(year1, month1, 1, 0, 0, 0, 0, time.UTC)

	year2, month1, _ := first.AddDate(0, 1, 0).Date()
	nextMonthTime := time.Date(year2, month1, day, 0, 0, 0, 0, time.UTC)

	if month1 != nextMonthTime.Month() {
		return first.AddDate(0, 2, -1)
	}

	return nextMonthTime

}

func dateTimeExample() {
	now := time.Now()

	tz, _ := time.LoadLocation("America/Los_Angeles")
	future := time.Date(2015, time.October, 21, 7, 28, 0, 0, tz)

	fmt.Println(now.String())
	fmt.Println(future.Format(time.RFC3339Nano))

	var seconds int = 10
	tenSeconds := time.Duration(seconds) * time.Second

	past := time.Date(1999, time.November, 12, 6, 38, 0, 0, time.UTC)
	dur := time.Now().Sub(past)

	fmt.Println(tenSeconds)
	fmt.Println(dur)

	jst, _ := time.LoadLocation("Asia/Tokyo")
	now2 := time.Date(2021, 5, 31, 20, 56, 00, 000, jst)
	wrongNextMonth := now2.AddDate(0, 1, 0)
	correctNextMonth := NextMonth(now2)
	fmt.Printf("wrongNextMonth: %v\n", wrongNextMonth)
	fmt.Printf("correctNextMonth: %v\n", correctNextMonth)

	normal := time.Date(2021, 7, 31, 00, 00, 00, 000, jst)
	fmt.Printf("normal: %v\n", normal)
}
