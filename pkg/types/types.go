package types

import "time"

type MeasureCountbyDateTime struct {
	DateTime time.Time
	Count    int
}
