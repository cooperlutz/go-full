package inbound

import (
	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/app"
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
		"notificationsandcommunications_notification_sent_handler",
		"notificationsandcommunications.notification_sent",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"notificationsandcommunications_notification_scheduled_handler",
		"notificationsandcommunications.notification_scheduled",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"notificationsandcommunications_notification_canceled_handler",
		"notificationsandcommunications.notification_canceled",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)

	// noop handler to initialize the topic table
	router.AddConsumerHandler(
		"notificationsandcommunications_notification_template_created_handler",
		"notificationsandcommunications.notification_template_created",
		s.pubSub.GetSubscriber(),
		eeventdriven.NewNoOpEventHandler().Handle(),
	)
}
