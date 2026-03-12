package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/domain/notificationsandcommunications"
)

type ScheduleNotification struct {
	//
	//RecipientId string,
	//
	//RecipientType string,
	//
	//Channel string,
	//
	//NotificationType string,
	//
	//ScheduledSendTime string,
	//
	//TemplateId *string,
	//
	// TODO
}

type ScheduleNotificationHandler struct {
	NotificationRepo notificationsandcommunications.NotificationRepository

	NotificationTemplateRepo notificationsandcommunications.NotificationTemplateRepository
}

func NewScheduleNotificationHandler(
	notificationRepo notificationsandcommunications.NotificationRepository,

	notificationtemplateRepo notificationsandcommunications.NotificationTemplateRepository,
) ScheduleNotificationHandler {
	return ScheduleNotificationHandler{
		NotificationRepo: notificationRepo,

		NotificationTemplateRepo: notificationtemplateRepo,
	}
}

func (h ScheduleNotificationHandler) Handle(ctx context.Context, cmd ScheduleNotification) error {
	// ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.app.command.schedule_notification.handle")
	// defer span.End()

	// TODO
	//err = h.NotificationRepo.UpdateNotification(ctx, uuid.MustParse(cmd.NotificationId), func(n *notificationsandcommunications.Notification) (*notificationsandcommunications.Notification, error) {
	//
	//	 err := n.ScheduleNotification(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return n, nil
	//})
	//if err != nil {
	//	return err
	//}

	// TODO
	//err = h.NotificationTemplateRepo.UpdateNotificationTemplate(ctx, uuid.MustParse(cmd.NotificationTemplateId), func(n *notificationsandcommunications.NotificationTemplate) (*notificationsandcommunications.NotificationTemplate, error) {
	//
	//	 err := n.ScheduleNotification(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return n, nil
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}
