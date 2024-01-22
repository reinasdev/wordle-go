package utils

import "time"

// CompareDate compares two dates
func CompareDate(date1 *time.Time, date2 *time.Time) (equal *bool) {
	equal = GetPointer(false)

	if date1 == nil || date2 == nil {
		return equal
	}

	if date1.Year() == date2.Year() && date1.Month() == date2.Month() && date1.Day() == date2.Day() {
		equal = GetPointer(true)
	}

	return equal
}
