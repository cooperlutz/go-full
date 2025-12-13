package exception

// ErrPingPongMsgValidation is returned when a PingPong message fails validation.
type ErrPingPongMsgValidation struct{}

// Error implements the error interface for ErrPingPongMsgValidation.
func (e ErrPingPongMsgValidation) Error() string {
	return "ya gotta send a ping or a pong"
}
