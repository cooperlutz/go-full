//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllVaccinationRecordsReadModel interface {
	FindAllVaccinationRecords(ctx context.Context) ([]VaccinationRecord, error)
}

type FindAllVaccinationRecordsHandler struct {
	readModel FindAllVaccinationRecordsReadModel
}

func NewFindAllVaccinationRecordsHandler(
	readModel FindAllVaccinationRecordsReadModel,
) FindAllVaccinationRecordsHandler {
	return FindAllVaccinationRecordsHandler{readModel: readModel}
}

func (h FindAllVaccinationRecordsHandler) Handle(ctx context.Context) ([]VaccinationRecord, error) {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.query.find_all_vaccination_records.handle")
	defer span.End()

	exams, err := h.readModel.FindAllVaccinationRecords(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
