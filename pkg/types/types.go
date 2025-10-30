package types

import "time"

type MeasureCountbyDateTimeMetric struct {
	DateTime time.Time
	Count    int
}

type QuantityMetric struct {
	Quantity int64
}
