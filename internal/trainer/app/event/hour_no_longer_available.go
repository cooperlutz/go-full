package event

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type HourNoLongerAvailable struct {
	HourTime string
}

type HourNoLongerAvailableHandler struct {
	pubsub eeventdriven.IPubSubEventProcessor
}

func NewHourNoLongerAvailableHandler(
	pubsub eeventdriven.IPubSubEventProcessor,
) HourNoLongerAvailableHandler {
	handler := HourNoLongerAvailableHandler{
		pubsub: pubsub,
	}
	handler.Handle(pubsub.GetRouter())

	return handler
}

func (h HourNoLongerAvailableHandler) Handle(router *message.Router) *message.Handler {
	return router.AddConsumerHandler(
		"handle_hour_no_longer_available",
		"hours",
		h.pubsub.GetSubscriber(),
		eeventdriven.TraceConsumerHandler(func(msg *message.Message) error {
			var event HourNoLongerAvailable
			if err := json.Unmarshal(msg.Payload, &event); err != nil {
				return err
			}

			msg.Ack()

			return nil
		}),
	)
}
