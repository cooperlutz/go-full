package event

type PingPongReceived struct {
	PingPongID string
	Message    string
}

func NewPingPongReceived(pingPongID, message string) PingPongReceived {
	return PingPongReceived{
		PingPongID: pingPongID,
		Message:    message,
	}
}
