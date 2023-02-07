package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	const layout = "1/2/2006 15:04:05"

	return parseTime(layout, date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	const layout = "January 2, 2006 15:04:05"
	parsedTime := parseTime(layout, date)

	return parsedTime.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	const layout = "Monday, January 2, 2006 15:04:05"
	appointmentTime := parseTime(layout, date)
	afternoonMinTime := time.Date(appointmentTime.Year(), appointmentTime.Month(), appointmentTime.Day(), 12, 0, 0, 0, appointmentTime.Location())
	afternoonMaxTime := time.Date(appointmentTime.Year(), appointmentTime.Month(), appointmentTime.Day(), 18, 0, 0, 0, appointmentTime.Location())

	return (appointmentTime.After(afternoonMinTime) || appointmentTime.Equal(afternoonMinTime)) && appointmentTime.Before(afternoonMaxTime)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	const layout = "Monday, January 2, 2006, at 15:04"

	return "You have an appointment on " + Schedule(date).Format(layout) + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now()

	return time.Date(now.Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}

// parseTime parses given timestring using given layout and ignores any errors
func parseTime(layout, date string) time.Time {
	parsedTime, _ := time.Parse(layout, date)

	return parsedTime
}
