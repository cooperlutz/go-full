package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/retailsales/domain/retailsales"
)

type AddItemToCart struct {
	//
	//CartId string,
	//
	//OwnerId string,
	//
	//ProductId string,
	//
	//Quantity int,
	//
	// TODO
}

type AddItemToCartHandler struct {
	SalesOrderRepo retailsales.SalesOrderRepository

	ShoppingCartRepo retailsales.ShoppingCartRepository
}

func NewAddItemToCartHandler(
	salesorderRepo retailsales.SalesOrderRepository,

	shoppingcartRepo retailsales.ShoppingCartRepository,
) AddItemToCartHandler {
	return AddItemToCartHandler{
		SalesOrderRepo: salesorderRepo,

		ShoppingCartRepo: shoppingcartRepo,
	}
}

func (h AddItemToCartHandler) Handle(ctx context.Context, cmd AddItemToCart) error {
	// ctx, span := telemetree.AddSpan(ctx, "retailsales.app.command.add_item_to_cart.handle")
	// defer span.End()

	// TODO
	//err = h.SalesOrderRepo.UpdateSalesOrder(ctx, uuid.MustParse(cmd.SalesOrderId), func(s *retailsales.SalesOrder) (*retailsales.SalesOrder, error) {
	//
	//	 err := s.AddItemToCart(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return s, nil
	//})
	//if err != nil {
	//	return err
	//}

	// TODO
	//err = h.ShoppingCartRepo.UpdateShoppingCart(ctx, uuid.MustParse(cmd.ShoppingCartId), func(s *retailsales.ShoppingCart) (*retailsales.ShoppingCart, error) {
	//
	//	 err := s.AddItemToCart(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return s, nil
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}
