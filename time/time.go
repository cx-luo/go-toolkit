// Package timeutil provides time manipulation utilities
package timeutil

import (
	"time"
)

const (
	// Common time formats
	FormatDate      = "2006-01-02"
	FormatTime      = "15:04:05"
	FormatDateTime  = "2006-01-02 15:04:05"
	FormatDateTimeT = "2006-01-02T15:04:05"
	FormatISO8601   = "2006-01-02T15:04:05Z07:00"
)

// Now returns the current time
func Now() time.Time {
	return time.Now()
}

// NowUnix returns the current Unix timestamp
func NowUnix() int64 {
	return time.Now().Unix()
}

// NowUnixNano returns the current Unix timestamp in nanoseconds
func NowUnixNano() int64 {
	return time.Now().UnixNano()
}

// Format formats a time.Time to string using the given format
func Format(t time.Time, format string) string {
	return t.Format(format)
}

// Parse parses a string to time.Time using the given format
func Parse(timeStr, format string) (time.Time, error) {
	return time.Parse(format, timeStr)
}

// ParseInLocation parses a string to time.Time in a specific location
func ParseInLocation(timeStr, format string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(format, timeStr, loc)
}

// Today returns the start of today (00:00:00)
func Today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// Tomorrow returns the start of tomorrow
func Tomorrow() time.Time {
	return Today().AddDate(0, 0, 1)
}

// Yesterday returns the start of yesterday
func Yesterday() time.Time {
	return Today().AddDate(0, 0, -1)
}

// StartOfDay returns the start of the day for the given time
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the end of the day for the given time
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// StartOfWeek returns the start of the week (Monday) for the given time
func StartOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7 // Sunday becomes 7
	}
	daysToMonday := weekday - 1
	return StartOfDay(t.AddDate(0, 0, -daysToMonday))
}

// EndOfWeek returns the end of the week (Sunday) for the given time
func EndOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	daysToSunday := 7 - weekday
	return EndOfDay(t.AddDate(0, 0, daysToSunday))
}

// StartOfMonth returns the start of the month for the given time
func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth returns the end of the month for the given time
func EndOfMonth(t time.Time) time.Time {
	return StartOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// StartOfYear returns the start of the year for the given time
func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear returns the end of the year for the given time
func EndOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 23, 59, 59, 999999999, t.Location())
}

// AddDays adds days to a time
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddMonths adds months to a time
func AddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// AddYears adds years to a time
func AddYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

// DiffDays calculates the difference in days between two times
func DiffDays(t1, t2 time.Time) int {
	return int(t1.Sub(t2).Hours() / 24)
}

// DiffHours calculates the difference in hours between two times
func DiffHours(t1, t2 time.Time) int64 {
	return int64(t1.Sub(t2).Hours())
}

// DiffMinutes calculates the difference in minutes between two times
func DiffMinutes(t1, t2 time.Time) int64 {
	return int64(t1.Sub(t2).Minutes())
}

// DiffSeconds calculates the difference in seconds between two times
func DiffSeconds(t1, t2 time.Time) int64 {
	return int64(t1.Sub(t2).Seconds())
}

// IsSameDay checks if two times are on the same day
func IsSameDay(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}

// IsSameWeek checks if two times are in the same week
func IsSameWeek(t1, t2 time.Time) bool {
	s1, s2 := StartOfWeek(t1), StartOfWeek(t2)
	return IsSameDay(s1, s2)
}

// IsSameMonth checks if two times are in the same month
func IsSameMonth(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month()
}

// IsSameYear checks if two times are in the same year
func IsSameYear(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year()
}

// UnixToTime converts Unix timestamp to time.Time
func UnixToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// UnixNanoToTime converts Unix nanosecond timestamp to time.Time
func UnixNanoToTime(nsec int64) time.Time {
	return time.Unix(0, nsec)
}

// TimeToUnix converts time.Time to Unix timestamp
func TimeToUnix(t time.Time) int64 {
	return t.Unix()
}

