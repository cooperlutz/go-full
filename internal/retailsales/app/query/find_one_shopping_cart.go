//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneShoppingCart struct {
	ShoppingCartID string
}

type FindOneShoppingCartReadModel interface {
	FindOneShoppingCart(ctx context.Context, shoppingcartId uuid.UUID) (ShoppingCart, error)
}

type FindOneShoppingCartHandler struct {
	readModel FindOneShoppingCartReadModel
}

func NewFindOneShoppingCartHandler(
	readModel FindOneShoppingCartReadModel,
) FindOneShoppingCartHandler {
	return FindOneShoppingCartHandler{readModel: readModel}
}

func (h FindOneShoppingCartHandler) Handle(ctx context.Context, qry FindOneShoppingCart) (ShoppingCart, error) {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.app.query.find_one_shopping_cart.handle")
	defer span.End()

	shoppingcart, err := h.readModel.FindOneShoppingCart(ctx, uuid.MustParse(qry.ShoppingCartID))
	if err != nil {
		return ShoppingCart{}, err
	}

	return shoppingcart, nil
}
