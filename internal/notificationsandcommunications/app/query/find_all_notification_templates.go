//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllNotificationTemplatesReadModel interface {
	FindAllNotificationTemplates(ctx context.Context) ([]NotificationTemplate, error)
}

type FindAllNotificationTemplatesHandler struct {
	readModel FindAllNotificationTemplatesReadModel
}

func NewFindAllNotificationTemplatesHandler(
	readModel FindAllNotificationTemplatesReadModel,
) FindAllNotificationTemplatesHandler {
	return FindAllNotificationTemplatesHandler{readModel: readModel}
}

func (h FindAllNotificationTemplatesHandler) Handle(ctx context.Context) ([]NotificationTemplate, error) {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.app.query.find_all_notification_templates.handle")
	defer span.End()

	exams, err := h.readModel.FindAllNotificationTemplates(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
