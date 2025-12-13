package common

import "time"

// PingPongResult represents the result of a PingPong operation.
type PingPongResult struct {
	Message string
}

// PingPongRawResult represents the raw result of a PingPong operation.
type PingPongRawResult struct {
	ID        string
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
	DeletedAt *time.Time
}
