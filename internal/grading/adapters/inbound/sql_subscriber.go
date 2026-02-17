package inbound

import (
	"github.com/cooperlutz/go-full/internal/grading/app"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

func RegisterEventHandlers(events app.Events, pubSub *eeventdriven.BasePgsqlPubSubProcessor) {
	router := pubSub.GetRouter()

	router.AddConsumerHandler(
		"grading_exam_submitted_handler",
		"examination.exam_submitted",
		pubSub.GetSubscriber(),
		events.ExamSubmitted.Handle(
			events.GradingStarted,
			events.GradingCompleted,
		),
	)
}
