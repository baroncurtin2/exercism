package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	today := time.Now()
	appt, _ := time.Parse("January 1, 2006 15:04:05", date)

	return appt.Before(today)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	appt, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)

	return appt.Hour() >= 12 && appt.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appt := Schedule(date)
	apptDate := appt.Format("January 2, 2006, at 15:04.")

	return fmt.Sprintf("You have an appointment on %s, %s", appt.Weekday(), apptDate)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now()
	return time.Date(now.Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
