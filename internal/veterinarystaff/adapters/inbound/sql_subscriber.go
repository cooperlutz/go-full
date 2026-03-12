package inbound

import (
	"github.com/cooperlutz/go-full/internal/veterinarystaff/app"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type SqlSubscriberAdapter struct {
	app    app.Events
	pubSub *eeventdriven.BasePgsqlPubSubProcessor
}

func NewSqlSubscriberAdapter(events app.Events, pubSub *eeventdriven.BasePgsqlPubSubProcessor) SqlSubscriberAdapter {
	return SqlSubscriberAdapter{
		app:    events,
		pubSub: pubSub,
	}
}

func (s SqlSubscriberAdapter) RegisterEventHandlers() {
	router := s.pubSub.GetRouter()

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"veterinarystaff_veterinarian_onboarded_handler",
		"veterinarystaff.veterinarian_onboarded",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"veterinarystaff_veterinarian_profile_updated_handler",
		"veterinarystaff.veterinarian_profile_updated",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"veterinarystaff_staff_availability_updated_handler",
		"veterinarystaff.staff_availability_updated",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"veterinarystaff_staff_member_deactivated_handler",
		"veterinarystaff.staff_member_deactivated",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)
}
