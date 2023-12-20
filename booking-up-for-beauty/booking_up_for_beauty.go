package booking

import (
	"time"
)

func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	t.Format("2006-01-02 15:04:05 -0700 MST")
	return t
}

func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return t.Before(time.Now())
}

func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return t.Hour() >= 12 && t.Hour() < 18
}

func Description(date string) string {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return "You have an appointment on " + t.Format("Monday, January 2, 2006, at 15:04") + "."
}

func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
