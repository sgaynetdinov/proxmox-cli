package utils

import "fmt"

func FormatOptionalUptime(seconds int64, available bool) string {
	if !available {
		return "n/a"
	}

	return formatSecondsHMS(seconds)
}

// formatSecondsHMS converts seconds into a compact uptime string.
func formatSecondsHMS(seconds int64) string {
	if seconds <= 0 {
		return "00:00:00"
	}
	h := seconds / 3600
	rem := seconds % 3600
	m := rem / 60
	s := rem % 60
	d := h / 24
	h = h % 24
	if d > 0 {
		return fmt.Sprintf("%dd %02d:%02d:%02d", d, h, m, s)
	}

	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
