package timeutil

import "time"

const (
	DateMinute = "2006-01-02 15:04"
	DateHour   = "2006-01-02 15"
	DateMonth  = "2006-01"
	DateYear   = "2006"
)

func FormatDateTime(source time.Time) string {
	return source.Format(time.DateTime)
}

func FormatMinuteTime(source time.Time) string {
	return source.Format(DateMinute)
}

func FormatDateHour(source time.Time) string {
	return source.Format(DateHour)
}

func FormatDateOnly(source time.Time) string {
	return source.Format(time.DateOnly)
}

func FormatDateMonth(source time.Time) string {
	return source.Format(DateMonth)
}

func FormatDateYear(source time.Time) string {
	return source.Format(DateYear)
}

func ParseDateTime(source string) (time.Time, error) {
	return time.ParseInLocation(time.DateTime, source, time.Local)
}

func ParseDateMinute(source string) (time.Time, error) {
	return time.ParseInLocation(DateMinute, source, time.Local)
}

func ParseDateHour(source string) (time.Time, error) {
	return time.ParseInLocation(DateHour, source, time.Local)
}

func ParseDateOnly(source string) (time.Time, error) {
	return time.ParseInLocation(time.DateOnly, source, time.Local)
}

func ParseDateMonth(source string) (time.Time, error) {
	return time.ParseInLocation(DateMonth, source, time.Local)
}

func ParseDateYear(source string) (time.Time, error) {
	return time.ParseInLocation(DateYear, source, time.Local)
}
