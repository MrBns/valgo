package is

import "strings"

// IsANSIC validates whether the string matches ANSIC time format.
// Format: "Mon Jan _2 15:04:05 2006"
func IsANSIC(v string) bool {
	return ansicRegex.MatchString(v)
}

// IsUnixDate validates whether the string matches Unix date format.
// Format: "Mon Jan _2 15:04:05 MST 2006"
func IsUnixDate(v string) bool {
	return unixDateRegex.MatchString(v)
}

// IsRubyDate validates whether the string matches Ruby date format.
// Format: "Mon Jan 02 15:04:05 -0700 2006"
func IsRubyDate(v string) bool {
	return rubyDateRegex.MatchString(v)
}

// IsRFC822 validates whether the string matches RFC822 time format.
// Format: "02 Jan 06 15:04 MST"
func IsRFC822(v string) bool {
	return rfc822Regex.MatchString(v)
}

// IsRFC822Z validates whether the string matches RFC822Z time format.
// Format: "02 Jan 06 15:04 -0700"
func IsRFC822Z(v string) bool {
	return rfc822ZRegex.MatchString(v)
}

// IsRFC850 validates whether the string matches RFC850 time format.
// Format: "Monday, 02-Jan-06 15:04:05 MST"
func IsRFC850(v string) bool {
	return rfc850Regex.MatchString(v)
}

// IsRFC1123 validates whether the string matches RFC1123 time format.
// Format: "Mon, 02 Jan 2006 15:04:05 MST"
func IsRFC1123(v string) bool {
	return rfc1123Regex.MatchString(v)
}

// IsRFC1123Z validates whether the string matches RFC1123Z time format.
// Format: "Mon, 02 Jan 2006 15:04:05 -0700"
func IsRFC1123Z(v string) bool {
	return rfc1123ZRegex.MatchString(v)
}

// IsRFC3339 validates whether the string matches RFC3339 time format.
// Format: "2006-01-02T15:04:05Z07:00"
func IsRFC3339(v string) bool {
	return rfc3339Regex.MatchString(v)
}

// IsRFC3339Nano validates whether the string matches RFC3339Nano time format.
// Format: "2006-01-02T15:04:05.999999999Z07:00"
func IsRFC3339Nano(v string) bool {
	// Precheck: minimum length is 29 (2006-01-02T15:04:05Z07:00)
	// maximum length is 39 (with 9 nanosecond digits)
	if len(v) < 20 || len(v) > 39 {
		return false
	}

	// Precheck: must contain 'T' and 'Z' or +/- for timezone
	if !strings.ContainsAny(v, "T") {
		return false
	}

	return rfc3339NanoRegex.MatchString(v)
}

// IsKitchen validates whether the string matches Kitchen time format.
// Format: "3:04PM"
func IsKitchen(v string) bool {
	return kitchenRegex.MatchString(v)
}

// IsStamp validates whether the string matches Stamp time format.
// Format: "Jan _2 15:04:05"
func IsStamp(v string) bool {
	return stampRegex.MatchString(v)
}

// IsStampMilli validates whether the string matches StampMilli time format.
// Format: "Jan _2 15:04:05.000"
func IsStampMilli(v string) bool {
	return stampMilliRegex.MatchString(v)
}

// IsStampMicro validates whether the string matches StampMicro time format.
// Format: "Jan _2 15:04:05.000000"
func IsStampMicro(v string) bool {
	return stampMicroRegex.MatchString(v)
}

// IsStampNano validates whether the string matches StampNano time format.
// Format: "Jan _2 15:04:05.000000000"
func IsStampNano(v string) bool {
	return stampNanoRegex.MatchString(v)
}

// IsDateTime validates whether the string matches DateTime format.
// Format: "2006-01-02 15:04:05"
func IsDateTime(v string) bool {
	return dateTimeRegex.MatchString(v)
}

// IsTimeOnly validates whether the string matches TimeOnly format.
// Format: "15:04:05"
func IsTimeOnly(v string) bool {
	return timeOnlyRegex.MatchString(v)
}
