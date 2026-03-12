//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneMedicalRecord struct {
	MedicalRecordID string
}

type FindOneMedicalRecordReadModel interface {
	FindOneMedicalRecord(ctx context.Context, medicalrecordId uuid.UUID) (MedicalRecord, error)
}

type FindOneMedicalRecordHandler struct {
	readModel FindOneMedicalRecordReadModel
}

func NewFindOneMedicalRecordHandler(
	readModel FindOneMedicalRecordReadModel,
) FindOneMedicalRecordHandler {
	return FindOneMedicalRecordHandler{readModel: readModel}
}

func (h FindOneMedicalRecordHandler) Handle(ctx context.Context, qry FindOneMedicalRecord) (MedicalRecord, error) {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.query.find_one_medical_record.handle")
	defer span.End()

	medicalrecord, err := h.readModel.FindOneMedicalRecord(ctx, uuid.MustParse(qry.MedicalRecordID))
	if err != nil {
		return MedicalRecord{}, err
	}

	return medicalrecord, nil
}
