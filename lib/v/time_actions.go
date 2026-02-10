package v

import (
	"fmt"
	"time"
)

// timeAction implements TimePipeAction for time.Time validation.
type timeAction struct {
	errorMsg func(v time.Time) string
	validate func(v time.Time) bool
}

// Run executes the validation function on the given time.Time value.
// Returns an error if validation fails.
func (action *timeAction) Run(value time.Time) error {
	if !action.validate(value) {
		return fmt.Errorf("%s", action.errorMsg(value))
	}
	return nil
}

// CustomTime creates a custom time validator using the provided validation function.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	CustomTime(func(v time.Time) bool { return v.Hour() >= 9 && v.Hour() < 17 })
func CustomTime(fn func(value time.Time) bool, option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("invalid time", v, option...)
		},
		validate: fn,
	}
}

// Before validates that a time.Time value is strictly before the specified time.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	Before(time.Now()) // validates v < now
func Before(t time.Time, option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time must be before "+t.String(), v, option...)
		},
		validate: func(v time.Time) bool {
			return v.Before(t)
		},
	}
}

// After validates that a time.Time value is strictly after the specified time.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	After(time.Now()) // validates v > now
func After(t time.Time, option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time must be after "+t.String(), v, option...)
		},
		validate: func(v time.Time) bool {
			return v.After(t)
		},
	}
}

// Between validates that a time.Time value is strictly between the specified start and end times.
// The optional ActionOptions parameter can be used to customize the error message.
//
// Example:
//
//	Between(startDate, endDate) // validates startDate < v < endDate
func Between(start time.Time, end time.Time, option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time must be between "+start.String()+" and "+end.String(), v, option...)
		},
		validate: func(v time.Time) bool {
			return v.After(start) && v.Before(end)
		},
	}
}

// BeforeNow validates that a time.Time value is in the past (before the current time).
// The optional ActionOptions parameter can be used to customize the error message.
func BeforeNow(option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time must be in the past", v, option...)
		},
		validate: func(v time.Time) bool {
			return v.Before(time.Now())
		},
	}
}

// AfterNow validates that a time.Time value is in the future (after the current time).
// The optional ActionOptions parameter can be used to customize the error message.
func AfterNow(option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time must be in the future", v, option...)
		},
		validate: func(v time.Time) bool {
			return v.After(time.Now())
		},
	}
}

// NotZero validates that a time.Time value is not the zero value.
// The optional ActionOptions parameter can be used to customize the error message.
func NotZero(option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time cannot be zero value", v, option...)
		},
		validate: func(v time.Time) bool {
			return !v.IsZero()
		},
	}
}

// SameDay validates that a time.Time value is on the same day as the specified time.
// The comparison is done using year, month, and day only, ignoring the time component.
// The optional ActionOptions parameter can be used to customize the error message.
func SameDay(t time.Time, option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time must be on the same day as "+t.String(), v, option...)
		},
		validate: func(v time.Time) bool {
			return v.Year() == t.Year() && v.Month() == t.Month() && v.Day() == t.Day()
		},
	}
}

// SameMonth validates that a time.Time value is in the same month as the specified time.
// The comparison is done using year and month only, ignoring the day and time components.
// The optional ActionOptions parameter can be used to customize the error message.
func SameMonth(t time.Time, option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time must be in the same month as "+t.String(), v, option...)
		},
		validate: func(v time.Time) bool {
			return v.Year() == t.Year() && v.Month() == t.Month()
		},
	}
}

// SameYear validates that a time.Time value is in the same year as the specified time.
// The comparison is done using year only, ignoring all other components.
// The optional ActionOptions parameter can be used to customize the error message.
func SameYear(t time.Time, option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time must be in the same year as "+t.String(), v, option...)
		},
		validate: func(v time.Time) bool {
			return v.Year() == t.Year()
		},
	}
}

// MinDate validates that a time.Time value is not before the specified minimum date.
// The comparison is inclusive (v >= minDate).
// Edge cases: zero values are considered valid if they match the condition.
// The optional ActionOptions parameter can be used to customize the error message.
func MinDate(minDate time.Time, option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time must be on or after "+minDate.String(), v, option...)
		},
		validate: func(v time.Time) bool {
			return v.After(minDate) || v.Equal(minDate)
		},
	}
}

// MaxDate validates that a time.Time value is not after the specified maximum date.
// The comparison is inclusive (v <= maxDate).
// Edge cases: zero values are considered valid if they match the condition.
// The optional ActionOptions parameter can be used to customize the error message.
func MaxDate(maxDate time.Time, option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time must be on or before "+maxDate.String(), v, option...)
		},
		validate: func(v time.Time) bool {
			return v.Before(maxDate) || v.Equal(maxDate)
		},
	}
}

// Equal validates that a time.Time value equals the specified time exactly (including nanoseconds).
// The optional ActionOptions parameter can be used to customize the error message.
//
// Edge case consideration: This comparison includes nanosecond precision, so times parsed from
// different sources may not be equal due to nanosecond differences.
func Equal(t time.Time, option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time must equal "+t.String(), v, option...)
		},
		validate: func(v time.Time) bool {
			return v.Equal(t)
		},
	}
}

// NotEqual validates that a time.Time value does not equal the specified time exactly (including nanoseconds).
// The optional ActionOptions parameter can be used to customize the error message.
//
// Edge case consideration: This comparison includes nanosecond precision.
func NotEqual(t time.Time, option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time must not equal "+t.String(), v, option...)
		},
		validate: func(v time.Time) bool {
			return !v.Equal(t)
		},
	}
}

// OldOf validates that a time.Time value is at least the specified number of days old.
// The value must be at least `days` days before the current time.
//
// Example:
//
//	OldOf(3) // time must be at least 3 days old
//
// Edge cases:
// - Negative days: treated as 0
// - DST transitions: calculations account for timezone changes
// - Leap years: handled correctly by Go's time package
func OldOf(days int, option ...ActionOptionFace) TimePipeAction {
	if days < 0 {
		days = 0
	}
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg(fmt.Sprintf("time must be at least %d days old", days), v, option...)
		},
		validate: func(v time.Time) bool {
			cutoff := time.Now().AddDate(0, 0, -days)
			return v.Before(cutoff)
		},
	}
}

// NewOf validates that a time.Time value is at least the specified number of days in the future.
// The value must be at least `days` days after the current time.
//
// Example:
//
//	NewOf(2) // time must be at least 2 days in the future
//
// Edge cases:
// - Negative days: treated as 0
// - DST transitions: calculations account for timezone changes
// - Leap years: handled correctly by Go's time package
func NewOf(days int, option ...ActionOptionFace) TimePipeAction {
	if days < 0 {
		days = 0
	}
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg(fmt.Sprintf("time must be at least %d days in the future", days), v, option...)
		},
		validate: func(v time.Time) bool {
			cutoff := time.Now().AddDate(0, 0, days)
			return v.After(cutoff)
		},
	}
}

// SameWeek validates that a time.Time value is in the same week as the specified time.
// The comparison uses ISO 8601 week numbering (Monday is the first day of the week).
// The optional ActionOptions parameter can be used to customize the error message.
//
// Edge cases:
// - Week boundaries: correctly handles transitions at Monday 00:00
// - Year boundaries: ISO week can span across calendar year boundaries
// - First/last week: handled correctly per ISO 8601
func SameWeek(t time.Time, option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time must be in the same week as "+t.String(), v, option...)
		},
		validate: func(v time.Time) bool {
			vYear, vWeek := v.ISOWeek()
			tYear, tWeek := t.ISOWeek()
			return vYear == tYear && vWeek == tWeek
		},
	}
}

// IsWeekday validates that a time.Time value falls on a weekday (Monday through Friday).
// The optional ActionOptions parameter can be used to customize the error message.
//
// Edge cases:
// - Timezone is preserved: validation is done in the time's local location
func IsWeekday(option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time must fall on a weekday (Monday-Friday)", v, option...)
		},
		validate: func(v time.Time) bool {
			day := v.Weekday()
			return day >= time.Monday && day <= time.Friday
		},
	}
}

// IsTimezone validates that a time.Time value has a valid timezone offset.
// Checks if the timezone offset is within valid ranges (-12:00 to +14:00).
// The optional ActionOptions parameter can be used to customize the error message.
//
// Edge cases:
// - Handles non-standard offsets like +5:30 (India), +5:45 (Nepal), +9:30 (Australia), -3:30 (Newfoundland)
// - Zero offset (UTC) is valid
// - Times without location info are considered UTC and valid
func IsTimezone(option ...ActionOptionFace) TimePipeAction {
	return &timeAction{
		errorMsg: func(v time.Time) string {
			return extractMsg("time has invalid timezone offset", v, option...)
		},
		validate: func(v time.Time) bool {
			_, offset := v.Zone()
			// offset is in seconds, valid range is -12:00 to +14:00
			// -12:00 = -43200 seconds, +14:00 = 50400 seconds
			return offset >= -43200 && offset <= 50400
		},
	}
}
