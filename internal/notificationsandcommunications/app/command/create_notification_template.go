package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/domain/notificationsandcommunications"
)

type CreateNotificationTemplate struct {
	//
	//Name string,
	//
	//NotificationType string,
	//
	//Channel string,
	//
	//SubjectTemplate *string,
	//
	//BodyTemplate string,
	//
	// TODO
}

type CreateNotificationTemplateHandler struct {
	NotificationRepo notificationsandcommunications.NotificationRepository

	NotificationTemplateRepo notificationsandcommunications.NotificationTemplateRepository
}

func NewCreateNotificationTemplateHandler(
	notificationRepo notificationsandcommunications.NotificationRepository,

	notificationtemplateRepo notificationsandcommunications.NotificationTemplateRepository,
) CreateNotificationTemplateHandler {
	return CreateNotificationTemplateHandler{
		NotificationRepo: notificationRepo,

		NotificationTemplateRepo: notificationtemplateRepo,
	}
}

func (h CreateNotificationTemplateHandler) Handle(ctx context.Context, cmd CreateNotificationTemplate) error {
	// ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.app.command.create_notification_template.handle")
	// defer span.End()

	// TODO
	//err = h.NotificationRepo.UpdateNotification(ctx, uuid.MustParse(cmd.NotificationId), func(n *notificationsandcommunications.Notification) (*notificationsandcommunications.Notification, error) {
	//
	//	 err := n.CreateNotificationTemplate(
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
	//	 err := n.CreateNotificationTemplate(
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
