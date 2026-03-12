//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllSuppliersReadModel interface {
	FindAllSuppliers(ctx context.Context) ([]Supplier, error)
}

type FindAllSuppliersHandler struct {
	readModel FindAllSuppliersReadModel
}

func NewFindAllSuppliersHandler(
	readModel FindAllSuppliersReadModel,
) FindAllSuppliersHandler {
	return FindAllSuppliersHandler{readModel: readModel}
}

func (h FindAllSuppliersHandler) Handle(ctx context.Context) ([]Supplier, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.query.find_all_suppliers.handle")
	defer span.End()

	exams, err := h.readModel.FindAllSuppliers(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
