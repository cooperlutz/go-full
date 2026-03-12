//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllVeterinariansReadModel interface {
	FindAllVeterinarians(ctx context.Context) ([]Veterinarian, error)
}

type FindAllVeterinariansHandler struct {
	readModel FindAllVeterinariansReadModel
}

func NewFindAllVeterinariansHandler(
	readModel FindAllVeterinariansReadModel,
) FindAllVeterinariansHandler {
	return FindAllVeterinariansHandler{readModel: readModel}
}

func (h FindAllVeterinariansHandler) Handle(ctx context.Context) ([]Veterinarian, error) {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.app.query.find_all_veterinarians.handle")
	defer span.End()

	exams, err := h.readModel.FindAllVeterinarians(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
