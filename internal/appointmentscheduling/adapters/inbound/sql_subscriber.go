package inbound

import (
	"github.com/cooperlutz/go-full/internal/appointmentscheduling/app"
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
		"appointmentscheduling_appointment_scheduled_handler",
		"appointmentscheduling.appointment_scheduled",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"appointmentscheduling_appointment_rescheduled_handler",
		"appointmentscheduling.appointment_rescheduled",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"appointmentscheduling_appointment_canceled_handler",
		"appointmentscheduling.appointment_canceled",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"appointmentscheduling_appointment_confirmed_handler",
		"appointmentscheduling.appointment_confirmed",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"appointmentscheduling_appointment_completed_handler",
		"appointmentscheduling.appointment_completed",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"appointmentscheduling_appointment_reminder_scheduled_handler",
		"appointmentscheduling.appointment_reminder_scheduled",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"appointmentscheduling_telemedicine_session_started_handler",
		"appointmentscheduling.telemedicine_session_started",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"appointmentscheduling_telemedicine_session_ended_handler",
		"appointmentscheduling.telemedicine_session_ended",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)
}
