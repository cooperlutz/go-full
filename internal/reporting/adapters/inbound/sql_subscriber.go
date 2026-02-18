package inbound

import (
	"github.com/cooperlutz/go-full/internal/reporting/app"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type Subscriber struct {
	app    app.Events
	pubSub *eeventdriven.BasePgsqlPubSubProcessor
}

func NewSubscriber(events app.Events, pubSub *eeventdriven.BasePgsqlPubSubProcessor) Subscriber {
	return Subscriber{
		app:    events,
		pubSub: pubSub,
	}
}

func (s Subscriber) RegisterEventHandlers() {
	s.pubSub.GetRouter().AddConsumerHandler(
		"reporting_exam_started_handler",
		"reporting.exam_started",
		s.pubSub.GetSubscriber(),
		s.app.ExamStarted.Handle(),
	)
	s.pubSub.GetRouter().AddConsumerHandler(
		"reporting_exam_submitted_handler",
		"reporting.exam_submitted",
		s.pubSub.GetSubscriber(),
		s.app.ExamSubmitted.Handle(),
	)
	s.pubSub.GetRouter().AddConsumerHandler(
		"reporting_grading_started_handler",
		"reporting.grading_started",
		s.pubSub.GetSubscriber(),
		s.app.GradingStarted.Handle(),
	)
	s.pubSub.GetRouter().AddConsumerHandler(
		"reporting_grading_completed_handler",
		"reporting.grading_completed",
		s.pubSub.GetSubscriber(),
		s.app.GradingCompleted.Handle(),
	)
}
