package utils

import "testing"

func TestFormatSecondsHMS(t *testing.T) {
	tests := []struct {
		name    string
		seconds int64
		want    string
	}{
		{name: "negative", seconds: -1, want: "00:00:00"},
		{name: "zero", seconds: 0, want: "00:00:00"},
		{name: "one second", seconds: 1, want: "00:00:01"},
		{name: "forty five seconds", seconds: 45, want: "00:00:45"},
		{name: "five minutes", seconds: 5 * 60, want: "00:05:00"},
		{name: "fifty nine minutes", seconds: 59*60 + 59, want: "00:59:59"},
		{name: "one hour", seconds: 3600, want: "01:00:00"},
		{name: "six hours", seconds: 6*3600 + 23*60, want: "06:23:00"},
		{name: "twelve hours", seconds: 12*3600 + 5*60 + 9, want: "12:05:09"},
		{name: "one second before day", seconds: 23*3600 + 59*60 + 59, want: "23:59:59"},
		{name: "one day", seconds: 24 * 3600, want: "1d 00:00:00"},
		{name: "one day five minutes", seconds: 24*3600 + 5*60, want: "1d 00:05:00"},
		{name: "three hundred hours", seconds: 300 * 3600, want: "12d 12:00:00"},
		{name: "four hundred days", seconds: 400*24*3600 + 3*3600 + 10*60, want: "400d 03:10:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatSecondsHMS(tt.seconds)
			if got != tt.want {
				t.Fatalf("formatSecondsHMS(%d) = %q, want %q", tt.seconds, got, tt.want)
			}
		})
	}
}

func TestFormatOptionalUptime(t *testing.T) {
	tests := []struct {
		name      string
		seconds   int64
		available bool
		want      string
	}{
		{name: "not available", seconds: 3600, available: false, want: "n/a"},
		{name: "available zero", seconds: 0, available: true, want: "00:00:00"},
		{name: "available five minutes", seconds: 5 * 60, available: true, want: "00:05:00"},
		{name: "available six hours", seconds: 6*3600 + 23*60, available: true, want: "06:23:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatOptionalUptime(tt.seconds, tt.available)
			if got != tt.want {
				t.Fatalf("FormatOptionalUptime(%d, %t) = %q, want %q", tt.seconds, tt.available, got, tt.want)
			}
		})
	}
}
