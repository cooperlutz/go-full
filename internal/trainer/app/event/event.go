package event

import "github.com/cooperlutz/go-full/pkg/eeventdriven"

type TrainerHourMadeAvailable struct {
	HourTime string
}

type TrainerHourMadeAvailableHandler struct {
	pubsub eeventdriven.IPubSubEventProcessor
}

func NewTrainerHourMadeAvailableHandler(
	pubsub eeventdriven.IPubSubEventProcessor,
) TrainerHourMadeAvailableHandler {
	return TrainerHourMadeAvailableHandler{
		pubsub: pubsub,
	}
}

func (h TrainerHourMadeAvailableHandler) Handle(event TrainerHourMadeAvailable) error {
	err := h.pubsub.EmitEvent("hours", event)
	if err != nil {
		return err
	}

	return nil
}
