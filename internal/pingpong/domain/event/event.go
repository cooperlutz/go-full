package event

type PingPongReceived struct {
	PingID  string
	Message string
}

func NewPingPongReceived(pingID, message string) PingPongReceived {
	return PingPongReceived{
		PingID:  pingID,
		Message: message,
	}
}
