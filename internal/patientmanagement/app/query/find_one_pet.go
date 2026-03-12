//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOnePet struct {
	PetID string
}

type FindOnePetReadModel interface {
	FindOnePet(ctx context.Context, petId uuid.UUID) (Pet, error)
}

type FindOnePetHandler struct {
	readModel FindOnePetReadModel
}

func NewFindOnePetHandler(
	readModel FindOnePetReadModel,
) FindOnePetHandler {
	return FindOnePetHandler{readModel: readModel}
}

func (h FindOnePetHandler) Handle(ctx context.Context, qry FindOnePet) (Pet, error) {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.query.find_one_pet.handle")
	defer span.End()

	pet, err := h.readModel.FindOnePet(ctx, uuid.MustParse(qry.PetID))
	if err != nil {
		return Pet{}, err
	}

	return pet, nil
}
