//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneTelemedicineSession struct {
	TelemedicineSessionID string
}

type FindOneTelemedicineSessionReadModel interface {
	FindOneTelemedicineSession(ctx context.Context, telemedicinesessionId uuid.UUID) (TelemedicineSession, error)
}

type FindOneTelemedicineSessionHandler struct {
	readModel FindOneTelemedicineSessionReadModel
}

func NewFindOneTelemedicineSessionHandler(
	readModel FindOneTelemedicineSessionReadModel,
) FindOneTelemedicineSessionHandler {
	return FindOneTelemedicineSessionHandler{readModel: readModel}
}

func (h FindOneTelemedicineSessionHandler) Handle(ctx context.Context, qry FindOneTelemedicineSession) (TelemedicineSession, error) {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.query.find_one_telemedicine_session.handle")
	defer span.End()

	telemedicinesession, err := h.readModel.FindOneTelemedicineSession(ctx, uuid.MustParse(qry.TelemedicineSessionID))
	if err != nil {
		return TelemedicineSession{}, err
	}

	return telemedicinesession, nil
}
