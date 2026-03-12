//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllPetsReadModel interface {
	FindAllPets(ctx context.Context) ([]Pet, error)
}

type FindAllPetsHandler struct {
	readModel FindAllPetsReadModel
}

func NewFindAllPetsHandler(
	readModel FindAllPetsReadModel,
) FindAllPetsHandler {
	return FindAllPetsHandler{readModel: readModel}
}

func (h FindAllPetsHandler) Handle(ctx context.Context) ([]Pet, error) {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.query.find_all_pets.handle")
	defer span.End()

	exams, err := h.readModel.FindAllPets(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
