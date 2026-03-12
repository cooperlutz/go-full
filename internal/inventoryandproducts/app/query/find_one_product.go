//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneProduct struct {
	ProductID string
}

type FindOneProductReadModel interface {
	FindOneProduct(ctx context.Context, productId uuid.UUID) (Product, error)
}

type FindOneProductHandler struct {
	readModel FindOneProductReadModel
}

func NewFindOneProductHandler(
	readModel FindOneProductReadModel,
) FindOneProductHandler {
	return FindOneProductHandler{readModel: readModel}
}

func (h FindOneProductHandler) Handle(ctx context.Context, qry FindOneProduct) (Product, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.query.find_one_product.handle")
	defer span.End()

	product, err := h.readModel.FindOneProduct(ctx, uuid.MustParse(qry.ProductID))
	if err != nil {
		return Product{}, err
	}

	return product, nil
}
