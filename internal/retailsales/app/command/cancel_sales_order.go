package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/retailsales/domain/retailsales"
)

type CancelSalesOrder struct {
	//
	//OrderId string,
	//
	//Reason string,
	//
	// TODO
}

type CancelSalesOrderHandler struct {
	SalesOrderRepo retailsales.SalesOrderRepository

	ShoppingCartRepo retailsales.ShoppingCartRepository
}

func NewCancelSalesOrderHandler(
	salesorderRepo retailsales.SalesOrderRepository,

	shoppingcartRepo retailsales.ShoppingCartRepository,
) CancelSalesOrderHandler {
	return CancelSalesOrderHandler{
		SalesOrderRepo: salesorderRepo,

		ShoppingCartRepo: shoppingcartRepo,
	}
}

func (h CancelSalesOrderHandler) Handle(ctx context.Context, cmd CancelSalesOrder) error {
	// ctx, span := telemetree.AddSpan(ctx, "retailsales.app.command.cancel_sales_order.handle")
	// defer span.End()

	// TODO
	//err = h.SalesOrderRepo.UpdateSalesOrder(ctx, uuid.MustParse(cmd.SalesOrderId), func(s *retailsales.SalesOrder) (*retailsales.SalesOrder, error) {
	//
	//	 err := s.CancelSalesOrder(
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
	//	 err := s.CancelSalesOrder(
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
