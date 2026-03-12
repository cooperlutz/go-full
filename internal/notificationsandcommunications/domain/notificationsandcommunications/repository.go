package notificationsandcommunications

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type NotificationRepository interface {
	AddNotification(ctx context.Context, notification *Notification) error

	GetNotification(ctx context.Context, id uuid.UUID) (*Notification, error)

	UpdateNotification(
		ctx context.Context,
		notificationId uuid.UUID,
		updateFn func(e *Notification) (*Notification, error),
	) error
}

// MapToNotification creates a Notification domain object from the given parameters.
// This should ONLY BE USED when reconstructing an Notification from its repository.
func MapToNotification(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//notificationId string,
	//
	//recipientId string,
	//
	//recipientType string,
	//
	//channel string,
	//
	//subject *string,
	//
	//messageBody string,
	//
	//status string,
	//
	//sentAt *string,
	//
	//notificationType string,
	//
) (*Notification, error) {
	return &Notification{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		//
		//notificationId: notificationId,
		//
		//recipientId: recipientId,
		//
		//recipientType: recipientType,
		//
		//channel: channel,
		//
		//subject: subject,
		//
		//messageBody: messageBody,
		//
		//status: status,
		//
		//sentAt: sentAt,
		//
		//notificationType: notificationType,
		//
		// TODO
	}, nil
}

type NotificationTemplateRepository interface {
	AddNotificationTemplate(ctx context.Context, notificationtemplate *NotificationTemplate) error

	GetNotificationTemplate(ctx context.Context, id uuid.UUID) (*NotificationTemplate, error)

	UpdateNotificationTemplate(
		ctx context.Context,
		notificationtemplateId uuid.UUID,
		updateFn func(e *NotificationTemplate) (*NotificationTemplate, error),
	) error
}

// MapToNotificationTemplate creates a NotificationTemplate domain object from the given parameters.
// This should ONLY BE USED when reconstructing an NotificationTemplate from its repository.
func MapToNotificationTemplate(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//templateId string,
	//
	//name string,
	//
	//notificationType string,
	//
	//channel string,
	//
	//subjectTemplate *string,
	//
	//bodyTemplate string,
	//
	//isActive bool,
	//
) (*NotificationTemplate, error) {
	return &NotificationTemplate{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		//
		//templateId: templateId,
		//
		//name: name,
		//
		//notificationType: notificationType,
		//
		//channel: channel,
		//
		//subjectTemplate: subjectTemplate,
		//
		//bodyTemplate: bodyTemplate,
		//
		//isActive: isActive,
		//
		// TODO
	}, nil
}
