package inbound

import (
	"github.com/cooperlutz/go-full/internal/reporting/app"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

func RegisterEventHandlers(events app.Events, pubSub *eeventdriven.BasePgsqlPubSubProcessor) {
	router := pubSub.GetRouter()

	router.AddConsumerHandler(
		"reporting_exam_started_handler",
		"examination.exam_started",
		pubSub.GetSubscriber(),
		events.ExamStarted.Handle(),
	)
	router.AddConsumerHandler(
		"reporting_exam_submitted_handler",
		"examination.exam_submitted",
		pubSub.GetSubscriber(),
		events.ExamSubmitted.Handle(),
	)
	router.AddConsumerHandler(
		"reporting_grading_started_handler",
		"grading.grading_started",
		pubSub.GetSubscriber(),
		events.GradingStarted.Handle(),
	)
	router.AddConsumerHandler(
		"reporting_grading_completed_handler",
		"grading.grading_completed",
		pubSub.GetSubscriber(),
		events.GradingCompleted.Handle(),
	)
}
