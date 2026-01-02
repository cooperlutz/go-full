package eeventdriven

// ErrPubSubHandlersNotImplemented is returned when Pub/Sub handlers are not implemented.
type ErrPubSubHandlersNotImplemented struct{}

// Error returns the error message.
func (e ErrPubSubHandlersNotImplemented) Error() string {
	return "handlers not implemented"
}
