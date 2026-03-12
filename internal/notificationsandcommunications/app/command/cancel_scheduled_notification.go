package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/domain/notificationsandcommunications"
)

type CancelScheduledNotification struct {
	//
	//NotificationId string,
	//
	// TODO
}

type CancelScheduledNotificationHandler struct {
	NotificationRepo notificationsandcommunications.NotificationRepository

	NotificationTemplateRepo notificationsandcommunications.NotificationTemplateRepository
}

func NewCancelScheduledNotificationHandler(
	notificationRepo notificationsandcommunications.NotificationRepository,

	notificationtemplateRepo notificationsandcommunications.NotificationTemplateRepository,
) CancelScheduledNotificationHandler {
	return CancelScheduledNotificationHandler{
		NotificationRepo: notificationRepo,

		NotificationTemplateRepo: notificationtemplateRepo,
	}
}

func (h CancelScheduledNotificationHandler) Handle(ctx context.Context, cmd CancelScheduledNotification) error {
	// ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.app.command.cancel_scheduled_notification.handle")
	// defer span.End()

	// TODO
	//err = h.NotificationRepo.UpdateNotification(ctx, uuid.MustParse(cmd.NotificationId), func(n *notificationsandcommunications.Notification) (*notificationsandcommunications.Notification, error) {
	//
	//	 err := n.CancelScheduledNotification(
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
	//	 err := n.CancelScheduledNotification(
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
