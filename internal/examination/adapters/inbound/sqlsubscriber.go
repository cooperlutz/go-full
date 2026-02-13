package inbound

import (
	"github.com/cooperlutz/go-full/internal/examination/app"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

func RegisterEventHandler(events app.Events, pubSub *eeventdriven.BasePgsqlPubSubProcessor) {
	router := pubSub.GetRouter()

	router.AddConsumerHandler(
		"examination_examsubmitted_handler",
		"examination.examsubmitted",
		pubSub.GetSubscriber(),
		events.NoOp.Handle(),
	)

	router.AddConsumerHandler(
		"examination_examstarted_handler",
		"examination.examstarted",
		pubSub.GetSubscriber(),
		events.NoOp.Handle(),
	)
}
