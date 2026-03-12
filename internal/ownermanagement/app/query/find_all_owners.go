//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllOwnersReadModel interface {
	FindAllOwners(ctx context.Context) ([]Owner, error)
}

type FindAllOwnersHandler struct {
	readModel FindAllOwnersReadModel
}

func NewFindAllOwnersHandler(
	readModel FindAllOwnersReadModel,
) FindAllOwnersHandler {
	return FindAllOwnersHandler{readModel: readModel}
}

func (h FindAllOwnersHandler) Handle(ctx context.Context) ([]Owner, error) {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.query.find_all_owners.handle")
	defer span.End()

	exams, err := h.readModel.FindAllOwners(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
