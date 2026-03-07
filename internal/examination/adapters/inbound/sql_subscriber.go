package inbound

import (
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type SqlSubscriberAdapter struct {
	pubSub *eeventdriven.BasePgsqlPubSubProcessor
}

func NewSqlSubscriberAdapter(
	pubSub *eeventdriven.BasePgsqlPubSubProcessor,
) SqlSubscriberAdapter {
	return SqlSubscriberAdapter{
		pubSub: pubSub,
	}
}

func (s SqlSubscriberAdapter) RegisterEventHandlers() {
	router := s.pubSub.GetRouter()

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"examination_exam_started_handler",
		"examination.exam_started",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)
	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"examination_exam_submitted_handler",
		"examination.exam_submitted",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)
}
