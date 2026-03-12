//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneNotification struct {
	NotificationID string
}

type FindOneNotificationReadModel interface {
	FindOneNotification(ctx context.Context, notificationId uuid.UUID) (Notification, error)
}

type FindOneNotificationHandler struct {
	readModel FindOneNotificationReadModel
}

func NewFindOneNotificationHandler(
	readModel FindOneNotificationReadModel,
) FindOneNotificationHandler {
	return FindOneNotificationHandler{readModel: readModel}
}

func (h FindOneNotificationHandler) Handle(ctx context.Context, qry FindOneNotification) (Notification, error) {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.app.query.find_one_notification.handle")
	defer span.End()

	notification, err := h.readModel.FindOneNotification(ctx, uuid.MustParse(qry.NotificationID))
	if err != nil {
		return Notification{}, err
	}

	return notification, nil
}
