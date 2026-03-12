package outbound

import (
	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/app/query"
	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/domain/notificationsandcommunications"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

// toDomain maps the NotificationNotification to the domain entity.
func (e NotificationsandcommunicationsNotification) toDomain() (*notificationsandcommunications.Notification, error) {
	return notificationsandcommunications.MapToNotification(
		e.NotificationID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.NotificationId,
		//
		//e.RecipientId,
		//
		//e.RecipientType,
		//
		//e.Channel,
		//
		//e.Subject,
		//
		//e.MessageBody,
		//
		//e.Status,
		//
		//e.SentAt,
		//
		//e.NotificationType,
		//
		// TODO
	)
}

// toQueryNotification maps the notificationNotification to the query.Notification.
func (e NotificationsandcommunicationsNotification) toQueryNotification() (query.Notification, error) {
	notification, err := e.toDomain()
	if err != nil {
		return query.Notification{}, err
	}

	return mapEntityNotificationToQuery(notification), nil
}

// notificationNotificationsToQuery maps a slice of NotificationNotification to a slice of query.Notification entities.
func notificationsandcommunicationsNotificationsToQuery(notifications []NotificationsandcommunicationsNotification) ([]query.Notification, error) {
	var domainNotifications []query.Notification

	for _, notification := range notifications {
		queryNotification, err := notification.toQueryNotification()
		if err != nil {
			return nil, err
		}

		domainNotifications = append(domainNotifications, queryNotification)
	}

	return domainNotifications, nil
}

// mapEntityNotificationToDB maps a domain Notification entity to the NotificationNotification database model.
func mapEntityNotificationToDB(notification *notificationsandcommunications.Notification) NotificationsandcommunicationsNotification {
	createdAt := notification.GetCreatedAtTime()
	updatedAt := notification.GetUpdatedAtTime()

	return NotificationsandcommunicationsNotification{
		NotificationID: pgxutil.UUIDToPgtypeUUID(notification.GetIdUUID()),
		CreatedAt:      pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:      pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:        notification.IsDeleted(),
		DeletedAt:      pgxutil.TimeToTimestampz(notification.GetDeletedAtTime()),
		//
		//NotificationId: GetNotificationId(),
		//
		//RecipientId: GetRecipientId(),
		//
		//RecipientType: GetRecipientType(),
		//
		//Channel: GetChannel(),
		//
		//Subject: GetSubject(),
		//
		//MessageBody: GetMessageBody(),
		//
		//Status: GetStatus(),
		//
		//SentAt: GetSentAt(),
		//
		//NotificationType: GetNotificationType(),
		//
		// TODO
	}
}

// mapEntityNotificationToQuery maps a domain Notification entity to a query.Notification.
func mapEntityNotificationToQuery(notification *notificationsandcommunications.Notification) query.Notification {
	return query.Notification{
		// TODO
	}
}

// toDomain maps the NotificationtemplateNotificationTemplate to the domain entity.
func (e NotificationsandcommunicationsNotificationTemplate) toDomain() (*notificationsandcommunications.NotificationTemplate, error) {
	return notificationsandcommunications.MapToNotificationTemplate(
		e.NotificationTemplateID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.TemplateId,
		//
		//e.Name,
		//
		//e.NotificationType,
		//
		//e.Channel,
		//
		//e.SubjectTemplate,
		//
		//e.BodyTemplate,
		//
		//e.IsActive,
		//
		// TODO
	)
}

// toQueryNotificationTemplate maps the notificationtemplateNotificationTemplate to the query.NotificationTemplate.
func (e NotificationsandcommunicationsNotificationTemplate) toQueryNotificationTemplate() (query.NotificationTemplate, error) {
	notificationtemplate, err := e.toDomain()
	if err != nil {
		return query.NotificationTemplate{}, err
	}

	return mapEntityNotificationTemplateToQuery(notificationtemplate), nil
}

// notificationtemplateNotificationTemplatesToQuery maps a slice of NotificationTemplateNotificationTemplate to a slice of query.NotificationTemplate entities.
func notificationsandcommunicationsNotificationTemplatesToQuery(notificationtemplates []NotificationsandcommunicationsNotificationTemplate) ([]query.NotificationTemplate, error) {
	var domainNotificationTemplates []query.NotificationTemplate

	for _, notificationtemplate := range notificationtemplates {
		queryNotificationTemplate, err := notificationtemplate.toQueryNotificationTemplate()
		if err != nil {
			return nil, err
		}

		domainNotificationTemplates = append(domainNotificationTemplates, queryNotificationTemplate)
	}

	return domainNotificationTemplates, nil
}

// mapEntityNotificationTemplateToDB maps a domain NotificationTemplate entity to the NotificationTemplateNotificationTemplate database model.
func mapEntityNotificationTemplateToDB(notificationtemplate *notificationsandcommunications.NotificationTemplate) NotificationsandcommunicationsNotificationTemplate {
	createdAt := notificationtemplate.GetCreatedAtTime()
	updatedAt := notificationtemplate.GetUpdatedAtTime()

	return NotificationsandcommunicationsNotificationTemplate{
		NotificationTemplateID: pgxutil.UUIDToPgtypeUUID(notificationtemplate.GetIdUUID()),
		CreatedAt:              pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:              pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:                notificationtemplate.IsDeleted(),
		DeletedAt:              pgxutil.TimeToTimestampz(notificationtemplate.GetDeletedAtTime()),
		//
		//TemplateId: GetTemplateId(),
		//
		//Name: GetName(),
		//
		//NotificationType: GetNotificationType(),
		//
		//Channel: GetChannel(),
		//
		//SubjectTemplate: GetSubjectTemplate(),
		//
		//BodyTemplate: GetBodyTemplate(),
		//
		//IsActive: GetIsActive(),
		//
		// TODO
	}
}

// mapEntityNotificationTemplateToQuery maps a domain NotificationTemplate entity to a query.NotificationTemplate.
func mapEntityNotificationTemplateToQuery(notificationtemplate *notificationsandcommunications.NotificationTemplate) query.NotificationTemplate {
	return query.NotificationTemplate{
		// TODO
	}
}
