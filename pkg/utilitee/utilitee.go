package utilitee

import "time"

// RightNow is a simple utility function to ensure consistency of returning a "now" timestamp being localized to UTC.
func RightNow() time.Time {
	return time.Now().UTC()
}

// StrPtr returns a pointer to the given string value.
func StrPtr(s string) *string {
	return &s
}

// BoolPtr returns a pointer to the given bool value.
func BoolPtr(b bool) *bool {
	return &b
}

// TimePtr returns a pointer to the given time.Time value.
func TimePtr(t time.Time) *time.Time {
	return &t
}

func IntPtr(i int) *int {
	return &i
}
