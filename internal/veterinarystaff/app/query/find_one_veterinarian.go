//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneVeterinarian struct {
	VeterinarianID string
}

type FindOneVeterinarianReadModel interface {
	FindOneVeterinarian(ctx context.Context, veterinarianId uuid.UUID) (Veterinarian, error)
}

type FindOneVeterinarianHandler struct {
	readModel FindOneVeterinarianReadModel
}

func NewFindOneVeterinarianHandler(
	readModel FindOneVeterinarianReadModel,
) FindOneVeterinarianHandler {
	return FindOneVeterinarianHandler{readModel: readModel}
}

func (h FindOneVeterinarianHandler) Handle(ctx context.Context, qry FindOneVeterinarian) (Veterinarian, error) {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.app.query.find_one_veterinarian.handle")
	defer span.End()

	veterinarian, err := h.readModel.FindOneVeterinarian(ctx, uuid.MustParse(qry.VeterinarianID))
	if err != nil {
		return Veterinarian{}, err
	}

	return veterinarian, nil
}
