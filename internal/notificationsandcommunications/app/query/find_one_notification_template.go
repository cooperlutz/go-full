//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneNotificationTemplate struct {
	NotificationTemplateID string
}

type FindOneNotificationTemplateReadModel interface {
	FindOneNotificationTemplate(ctx context.Context, notificationtemplateId uuid.UUID) (NotificationTemplate, error)
}

type FindOneNotificationTemplateHandler struct {
	readModel FindOneNotificationTemplateReadModel
}

func NewFindOneNotificationTemplateHandler(
	readModel FindOneNotificationTemplateReadModel,
) FindOneNotificationTemplateHandler {
	return FindOneNotificationTemplateHandler{readModel: readModel}
}

func (h FindOneNotificationTemplateHandler) Handle(ctx context.Context, qry FindOneNotificationTemplate) (NotificationTemplate, error) {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.app.query.find_one_notification_template.handle")
	defer span.End()

	notificationtemplate, err := h.readModel.FindOneNotificationTemplate(ctx, uuid.MustParse(qry.NotificationTemplateID))
	if err != nil {
		return NotificationTemplate{}, err
	}

	return notificationtemplate, nil
}
