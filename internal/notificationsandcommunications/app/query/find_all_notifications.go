//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllNotificationsReadModel interface {
	FindAllNotifications(ctx context.Context) ([]Notification, error)
}

type FindAllNotificationsHandler struct {
	readModel FindAllNotificationsReadModel
}

func NewFindAllNotificationsHandler(
	readModel FindAllNotificationsReadModel,
) FindAllNotificationsHandler {
	return FindAllNotificationsHandler{readModel: readModel}
}

func (h FindAllNotificationsHandler) Handle(ctx context.Context) ([]Notification, error) {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.app.query.find_all_notifications.handle")
	defer span.End()

	exams, err := h.readModel.FindAllNotifications(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
