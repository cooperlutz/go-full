package event

type PingReceived struct {
	PingID  string
	Message string
}

func NewPingPongReceived(pingID, message string) PingReceived {
	return PingReceived{
		PingID:  pingID,
		Message: message,
	}
}
