//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllProductsReadModel interface {
	FindAllProducts(ctx context.Context) ([]Product, error)
}

type FindAllProductsHandler struct {
	readModel FindAllProductsReadModel
}

func NewFindAllProductsHandler(
	readModel FindAllProductsReadModel,
) FindAllProductsHandler {
	return FindAllProductsHandler{readModel: readModel}
}

func (h FindAllProductsHandler) Handle(ctx context.Context) ([]Product, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.query.find_all_products.handle")
	defer span.End()

	exams, err := h.readModel.FindAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
