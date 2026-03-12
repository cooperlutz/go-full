//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneVaccinationRecord struct {
	VaccinationRecordID string
}

type FindOneVaccinationRecordReadModel interface {
	FindOneVaccinationRecord(ctx context.Context, vaccinationrecordId uuid.UUID) (VaccinationRecord, error)
}

type FindOneVaccinationRecordHandler struct {
	readModel FindOneVaccinationRecordReadModel
}

func NewFindOneVaccinationRecordHandler(
	readModel FindOneVaccinationRecordReadModel,
) FindOneVaccinationRecordHandler {
	return FindOneVaccinationRecordHandler{readModel: readModel}
}

func (h FindOneVaccinationRecordHandler) Handle(ctx context.Context, qry FindOneVaccinationRecord) (VaccinationRecord, error) {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.query.find_one_vaccination_record.handle")
	defer span.End()

	vaccinationrecord, err := h.readModel.FindOneVaccinationRecord(ctx, uuid.MustParse(qry.VaccinationRecordID))
	if err != nil {
		return VaccinationRecord{}, err
	}

	return vaccinationrecord, nil
}
