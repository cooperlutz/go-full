package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/retailsales/domain/retailsales"
)

type RemoveItemFromCart struct {
	//
	//CartId string,
	//
	//ProductId string,
	//
	// TODO
}

type RemoveItemFromCartHandler struct {
	SalesOrderRepo retailsales.SalesOrderRepository

	ShoppingCartRepo retailsales.ShoppingCartRepository
}

func NewRemoveItemFromCartHandler(
	salesorderRepo retailsales.SalesOrderRepository,

	shoppingcartRepo retailsales.ShoppingCartRepository,
) RemoveItemFromCartHandler {
	return RemoveItemFromCartHandler{
		SalesOrderRepo: salesorderRepo,

		ShoppingCartRepo: shoppingcartRepo,
	}
}

func (h RemoveItemFromCartHandler) Handle(ctx context.Context, cmd RemoveItemFromCart) error {
	// ctx, span := telemetree.AddSpan(ctx, "retailsales.app.command.remove_item_from_cart.handle")
	// defer span.End()

	// TODO
	//err = h.SalesOrderRepo.UpdateSalesOrder(ctx, uuid.MustParse(cmd.SalesOrderId), func(s *retailsales.SalesOrder) (*retailsales.SalesOrder, error) {
	//
	//	 err := s.RemoveItemFromCart(
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
	//	 err := s.RemoveItemFromCart(
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
