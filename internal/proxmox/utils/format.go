package utils

import "fmt"

// FormatSecondsHMS converts a duration in seconds to HH:MM:SS.
// Hours are not capped (e.g., 100:05:09 for long uptimes).
func FormatSecondsHMS(seconds int64) string {
	var rowFormat = "%02d:%02d:%02d"

	if seconds <= 0 {
		return fmt.Sprintf(rowFormat, 0, 0, 0)
	}

	h := seconds / 3600
	rem := seconds % 3600
	m := rem / 60
	s := rem % 60
	return fmt.Sprintf(rowFormat, h, m, s)
}
