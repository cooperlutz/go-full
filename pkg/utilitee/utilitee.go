package utilitee

import "time"

// RightNow is a simple utility function to ensure consistency of returning a "now" timestamp being localized to UTC.
func RightNow() time.Time {
	return time.Now().UTC()
}

func StrPtr(s string) *string {
	return &s
}

func BoolPtr(b bool) *bool {
	return &b
}
