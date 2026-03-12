//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllMedicalRecordsReadModel interface {
	FindAllMedicalRecords(ctx context.Context) ([]MedicalRecord, error)
}

type FindAllMedicalRecordsHandler struct {
	readModel FindAllMedicalRecordsReadModel
}

func NewFindAllMedicalRecordsHandler(
	readModel FindAllMedicalRecordsReadModel,
) FindAllMedicalRecordsHandler {
	return FindAllMedicalRecordsHandler{readModel: readModel}
}

func (h FindAllMedicalRecordsHandler) Handle(ctx context.Context) ([]MedicalRecord, error) {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.query.find_all_medical_records.handle")
	defer span.End()

	exams, err := h.readModel.FindAllMedicalRecords(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
