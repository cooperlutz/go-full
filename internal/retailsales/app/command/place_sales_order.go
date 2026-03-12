package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/retailsales/domain/retailsales"
)

type PlaceSalesOrder struct {
	//
	//OwnerId string,
	//
	//LineItems string,
	//
	//Channel string,
	//
	//DiscountAmount *float32,
	//
	// TODO
}

type PlaceSalesOrderHandler struct {
	SalesOrderRepo retailsales.SalesOrderRepository

	ShoppingCartRepo retailsales.ShoppingCartRepository
}

func NewPlaceSalesOrderHandler(
	salesorderRepo retailsales.SalesOrderRepository,

	shoppingcartRepo retailsales.ShoppingCartRepository,
) PlaceSalesOrderHandler {
	return PlaceSalesOrderHandler{
		SalesOrderRepo: salesorderRepo,

		ShoppingCartRepo: shoppingcartRepo,
	}
}

func (h PlaceSalesOrderHandler) Handle(ctx context.Context, cmd PlaceSalesOrder) error {
	// ctx, span := telemetree.AddSpan(ctx, "retailsales.app.command.place_sales_order.handle")
	// defer span.End()

	// TODO
	//err = h.SalesOrderRepo.UpdateSalesOrder(ctx, uuid.MustParse(cmd.SalesOrderId), func(s *retailsales.SalesOrder) (*retailsales.SalesOrder, error) {
	//
	//	 err := s.PlaceSalesOrder(
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
	//	 err := s.PlaceSalesOrder(
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
