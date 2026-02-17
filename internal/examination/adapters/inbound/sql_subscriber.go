package inbound

import (
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

func RegisterEventHandlers(pubSub *eeventdriven.BasePgsqlPubSubProcessor) {
	router := pubSub.GetRouter()

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"examination_exam_started_handler",
		"examination.exam_started",
		pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)
}
