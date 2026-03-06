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
	switch {
	case seconds <= 0:
		return "just now"
	case seconds < 60:
		return "just now"
	case seconds < 3600:
		return fmt.Sprintf("%dm", seconds/60)
	}

	h := seconds / 3600
	rem := seconds % 3600
	m := rem / 60
	d := h / 24
	h = h % 24
	return fmt.Sprintf("%dd %02d:%02d", d, h, m)
}
