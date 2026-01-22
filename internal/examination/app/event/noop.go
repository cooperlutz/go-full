package event

import "github.com/ThreeDotsLabs/watermill/message"

// temporary no-op event and handler for testing purposes.
type NoOpEvent struct{}

// NoOpEventHandler is a handler that does nothing and acknowledges the message.
type NoOpEventHandler struct{}

// NewNoOpEventHandler creates a new NoOpEventHandler instance.
func NewNoOpEventHandler() NoOpEventHandler {
	return NoOpEventHandler{}
}

// Handle returns a message handler function that acknowledges the message without processing.
func (h NoOpEventHandler) Handle() message.NoPublishHandlerFunc {
	return func(msg *message.Message) error {
		msg.Ack()

		return nil
	}
}
