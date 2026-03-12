//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneOwner struct {
	OwnerID string
}

type FindOneOwnerReadModel interface {
	FindOneOwner(ctx context.Context, ownerId uuid.UUID) (Owner, error)
}

type FindOneOwnerHandler struct {
	readModel FindOneOwnerReadModel
}

func NewFindOneOwnerHandler(
	readModel FindOneOwnerReadModel,
) FindOneOwnerHandler {
	return FindOneOwnerHandler{readModel: readModel}
}

func (h FindOneOwnerHandler) Handle(ctx context.Context, qry FindOneOwner) (Owner, error) {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.query.find_one_owner.handle")
	defer span.End()

	owner, err := h.readModel.FindOneOwner(ctx, uuid.MustParse(qry.OwnerID))
	if err != nil {
		return Owner{}, err
	}

	return owner, nil
}
