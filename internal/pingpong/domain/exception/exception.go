package exception

type ErrPingPongMsgValidation struct{}

func (e ErrPingPongMsgValidation) Error() string {
	return "ya gotta send a ping or a pong"
}
