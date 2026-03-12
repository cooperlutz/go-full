package app

import (
	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/app/command"
	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/app/event"
	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/app/query"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Events   Events
}

type Commands struct {
	SendNotification command.SendNotificationHandler

	ScheduleNotification command.ScheduleNotificationHandler

	CancelScheduledNotification command.CancelScheduledNotificationHandler

	CreateNotificationTemplate command.CreateNotificationTemplateHandler
}

type Queries struct {
	FindAllNotifications query.FindAllNotificationsHandler
	FindOneNotification  query.FindOneNotificationHandler

	FindAllNotificationTemplates query.FindAllNotificationTemplatesHandler
	FindOneNotificationTemplate  query.FindOneNotificationTemplateHandler
}

type Events struct {
	NotificationSent event.NotificationSentHandler

	NotificationScheduled event.NotificationScheduledHandler

	NotificationCancelled event.NotificationCancelledHandler

	NotificationTemplateCreated event.NotificationTemplateCreatedHandler

	AppointmentScheduled event.AppointmentScheduledHandler

	AppointmentReminderScheduled event.AppointmentReminderScheduledHandler

	VaccinationDueReminderScheduled event.VaccinationDueReminderScheduledHandler

	LowStockAlertTriggered event.LowStockAlertTriggeredHandler
}

// NewApplication initializes the NotificationsAndCommunications application with its dependencies.
func NewApplication( //nolint:funlen // it's fine
	pgconn deebee.IDatabase,
	pubSub eeventdriven.IPubSubEventProcessor,
) (Application, error) {
	notificationRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	notificationtemplateRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	app := Application{
		Commands: Commands{
			SendNotification: command.NewSendNotificationHandler(

				notificationRepository,

				notificationtemplateRepository,
			),
			ScheduleNotification: command.NewScheduleNotificationHandler(

				notificationRepository,

				notificationtemplateRepository,
			),
			CancelScheduledNotification: command.NewCancelScheduledNotificationHandler(

				notificationRepository,

				notificationtemplateRepository,
			),
			CreateNotificationTemplate: command.NewCreateNotificationTemplateHandler(

				notificationRepository,

				notificationtemplateRepository,
			),
		},
		Queries: Queries{
			FindAllNotifications: query.NewFindAllNotificationsHandler(
				notificationRepository,
			),
			FindOneNotification: query.NewFindOneNotificationHandler(
				notificationRepository,
			),

			FindAllNotificationTemplates: query.NewFindAllNotificationTemplatesHandler(
				notificationtemplateRepository,
			),
			FindOneNotificationTemplate: query.NewFindOneNotificationTemplateHandler(
				notificationtemplateRepository,
			),
		},
		Events: Events{
			NotificationSent: event.NewNotificationSentHandler(
				pubSub,
			),

			NotificationScheduled: event.NewNotificationScheduledHandler(
				pubSub,
			),

			NotificationCancelled: event.NewNotificationCancelledHandler(
				pubSub,
			),

			NotificationTemplateCreated: event.NewNotificationTemplateCreatedHandler(
				pubSub,
			),

			AppointmentScheduled: event.NewAppointmentScheduledHandler(
				pubSub,
			),

			AppointmentReminderScheduled: event.NewAppointmentReminderScheduledHandler(
				pubSub,
			),

			VaccinationDueReminderScheduled: event.NewVaccinationDueReminderScheduledHandler(
				pubSub,
			),

			LowStockAlertTriggered: event.NewLowStockAlertTriggeredHandler(
				pubSub,
			),
		},
	}

	return app, nil
}
