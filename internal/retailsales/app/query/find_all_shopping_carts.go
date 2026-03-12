//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllShoppingCartsReadModel interface {
	FindAllShoppingCarts(ctx context.Context) ([]ShoppingCart, error)
}

type FindAllShoppingCartsHandler struct {
	readModel FindAllShoppingCartsReadModel
}

func NewFindAllShoppingCartsHandler(
	readModel FindAllShoppingCartsReadModel,
) FindAllShoppingCartsHandler {
	return FindAllShoppingCartsHandler{readModel: readModel}
}

func (h FindAllShoppingCartsHandler) Handle(ctx context.Context) ([]ShoppingCart, error) {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.app.query.find_all_shopping_carts.handle")
	defer span.End()

	exams, err := h.readModel.FindAllShoppingCarts(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
