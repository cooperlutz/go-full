//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllTelemedicineSessionsReadModel interface {
	FindAllTelemedicineSessions(ctx context.Context) ([]TelemedicineSession, error)
}

type FindAllTelemedicineSessionsHandler struct {
	readModel FindAllTelemedicineSessionsReadModel
}

func NewFindAllTelemedicineSessionsHandler(
	readModel FindAllTelemedicineSessionsReadModel,
) FindAllTelemedicineSessionsHandler {
	return FindAllTelemedicineSessionsHandler{readModel: readModel}
}

func (h FindAllTelemedicineSessionsHandler) Handle(ctx context.Context) ([]TelemedicineSession, error) {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.query.find_all_telemedicine_sessions.handle")
	defer span.End()

	exams, err := h.readModel.FindAllTelemedicineSessions(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
