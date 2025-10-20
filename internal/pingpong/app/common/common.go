package common

import "time"

type PingPongResult struct {
	Message string
}

type PingPongRawResult struct {
	ID        string
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
	DeletedAt *time.Time
}
