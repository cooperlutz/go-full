//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneSupplier struct {
	SupplierID string
}

type FindOneSupplierReadModel interface {
	FindOneSupplier(ctx context.Context, supplierId uuid.UUID) (Supplier, error)
}

type FindOneSupplierHandler struct {
	readModel FindOneSupplierReadModel
}

func NewFindOneSupplierHandler(
	readModel FindOneSupplierReadModel,
) FindOneSupplierHandler {
	return FindOneSupplierHandler{readModel: readModel}
}

func (h FindOneSupplierHandler) Handle(ctx context.Context, qry FindOneSupplier) (Supplier, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.query.find_one_supplier.handle")
	defer span.End()

	supplier, err := h.readModel.FindOneSupplier(ctx, uuid.MustParse(qry.SupplierID))
	if err != nil {
		return Supplier{}, err
	}

	return supplier, nil
}
