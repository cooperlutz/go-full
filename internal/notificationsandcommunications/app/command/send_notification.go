package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/domain/notificationsandcommunications"
)

type SendNotification struct {
	//
	//RecipientId string,
	//
	//RecipientType string,
	//
	//Channel string,
	//
	//NotificationType string,
	//
	//TemplateId *string,
	//
	//MessageBody *string,
	//
	// TODO
}

type SendNotificationHandler struct {
	NotificationRepo notificationsandcommunications.NotificationRepository

	NotificationTemplateRepo notificationsandcommunications.NotificationTemplateRepository
}

func NewSendNotificationHandler(
	notificationRepo notificationsandcommunications.NotificationRepository,

	notificationtemplateRepo notificationsandcommunications.NotificationTemplateRepository,
) SendNotificationHandler {
	return SendNotificationHandler{
		NotificationRepo: notificationRepo,

		NotificationTemplateRepo: notificationtemplateRepo,
	}
}

func (h SendNotificationHandler) Handle(ctx context.Context, cmd SendNotification) error {
	// ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.app.command.send_notification.handle")
	// defer span.End()

	// TODO
	//err = h.NotificationRepo.UpdateNotification(ctx, uuid.MustParse(cmd.NotificationId), func(n *notificationsandcommunications.Notification) (*notificationsandcommunications.Notification, error) {
	//
	//	 err := n.SendNotification(
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
	//	 err := n.SendNotification(
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
