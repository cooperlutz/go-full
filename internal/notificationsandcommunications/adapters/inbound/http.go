package inbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/app"
	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// HttpAdapter represents the HTTP server for the NotificationsAndCommunications module.
type HttpAdapter struct {
	app app.Application
}

// NewHttpAdapter creates a new HttpAdapter instance with the provided NotificationsAndCommunications application.
func NewHttpAdapter(application app.Application) HttpAdapter {
	return HttpAdapter{
		app: application,
	}
}

// StrictHandler returns a strict HTTP handler for the NotificationsAndCommunications module.
func (h HttpAdapter) StrictHandler() ServerInterface {
	return NewStrictHandler(h, nil)
}

// (GET /v1/notifications).
func (h HttpAdapter) FindAllNotifications(ctx context.Context, request FindAllNotificationsRequestObject) (FindAllNotificationsResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "notification.adapters.inbound.http.find_all_notifications")
	defer span.End()

	notification, err := h.app.Queries.FindAllNotifications.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseNotifications []Notification
	for _, e := range notification {
		responseNotifications = append(responseNotifications, queryNotificationToHttpNotification(e))
	}

	return FindAllNotifications200JSONResponse(responseNotifications), nil
}

// (GET /v1/notification/{notificationId}).
func (h HttpAdapter) FindOneNotification(ctx context.Context, request FindOneNotificationRequestObject) (FindOneNotificationResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_notification")
	defer span.End()

	notification, err := h.app.Queries.FindOneNotification.Handle(ctx, query.FindOneNotification{NotificationID: request.NotificationId})
	if err != nil {
		return nil, err
	}

	return FindOneNotification200JSONResponse(queryNotificationToHttpNotification(notification)), nil
}

func queryNotificationToHttpNotification(e query.Notification) Notification {
	return Notification{
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

// (GET /v1/notificationtemplates).
func (h HttpAdapter) FindAllNotificationTemplates(ctx context.Context, request FindAllNotificationTemplatesRequestObject) (FindAllNotificationTemplatesResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "notificationtemplate.adapters.inbound.http.find_all_notificationtemplates")
	defer span.End()

	notificationtemplate, err := h.app.Queries.FindAllNotificationTemplates.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseNotificationTemplates []NotificationTemplate
	for _, e := range notificationtemplate {
		responseNotificationTemplates = append(responseNotificationTemplates, queryNotificationTemplateToHttpNotificationTemplate(e))
	}

	return FindAllNotificationTemplates200JSONResponse(responseNotificationTemplates), nil
}

// (GET /v1/notificationtemplate/{notification_templateId}).
func (h HttpAdapter) FindOneNotificationTemplate(ctx context.Context, request FindOneNotificationTemplateRequestObject) (FindOneNotificationTemplateResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_notification_template")
	defer span.End()

	notificationtemplate, err := h.app.Queries.FindOneNotificationTemplate.Handle(ctx, query.FindOneNotificationTemplate{NotificationTemplateID: request.NotificationTemplateId})
	if err != nil {
		return nil, err
	}

	return FindOneNotificationTemplate200JSONResponse(queryNotificationTemplateToHttpNotificationTemplate(notificationtemplate)), nil
}

func queryNotificationTemplateToHttpNotificationTemplate(e query.NotificationTemplate) NotificationTemplate {
	return NotificationTemplate{
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
