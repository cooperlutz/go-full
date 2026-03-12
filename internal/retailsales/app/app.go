package app

import (
	"github.com/cooperlutz/go-full/internal/retailsales/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/retailsales/app/command"
	"github.com/cooperlutz/go-full/internal/retailsales/app/event"
	"github.com/cooperlutz/go-full/internal/retailsales/app/query"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Events   Events
}

type Commands struct {
	AddItemToCart command.AddItemToCartHandler

	RemoveItemFromCart command.RemoveItemFromCartHandler

	PlaceSalesOrder command.PlaceSalesOrderHandler

	FulfillSalesOrder command.FulfillSalesOrderHandler

	CancelSalesOrder command.CancelSalesOrderHandler
}

type Queries struct {
	FindAllSalesOrders query.FindAllSalesOrdersHandler
	FindOneSalesOrder  query.FindOneSalesOrderHandler

	FindAllShoppingCarts query.FindAllShoppingCartsHandler
	FindOneShoppingCart  query.FindOneShoppingCartHandler
}

type Events struct {
	ItemAddedToCart event.ItemAddedToCartHandler

	ItemRemovedFromCart event.ItemRemovedFromCartHandler

	SalesOrderPlaced event.SalesOrderPlacedHandler

	InventoryDeductionRequested event.InventoryDeductionRequestedHandler

	SalesOrderFulfilled event.SalesOrderFulfilledHandler

	SalesOrderCancelled event.SalesOrderCancelledHandler

	InventoryRestockRequested event.InventoryRestockRequestedHandler
}

// NewApplication initializes the RetailSales application with its dependencies.
func NewApplication( //nolint:funlen // it's fine
	pgconn deebee.IDatabase,
	pubSub eeventdriven.IPubSubEventProcessor,
) (Application, error) {
	salesorderRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	shoppingcartRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	app := Application{
		Commands: Commands{
			AddItemToCart: command.NewAddItemToCartHandler(

				salesorderRepository,

				shoppingcartRepository,
			),
			RemoveItemFromCart: command.NewRemoveItemFromCartHandler(

				salesorderRepository,

				shoppingcartRepository,
			),
			PlaceSalesOrder: command.NewPlaceSalesOrderHandler(

				salesorderRepository,

				shoppingcartRepository,
			),
			FulfillSalesOrder: command.NewFulfillSalesOrderHandler(

				salesorderRepository,

				shoppingcartRepository,
			),
			CancelSalesOrder: command.NewCancelSalesOrderHandler(

				salesorderRepository,

				shoppingcartRepository,
			),
		},
		Queries: Queries{
			FindAllSalesOrders: query.NewFindAllSalesOrdersHandler(
				salesorderRepository,
			),
			FindOneSalesOrder: query.NewFindOneSalesOrderHandler(
				salesorderRepository,
			),

			FindAllShoppingCarts: query.NewFindAllShoppingCartsHandler(
				shoppingcartRepository,
			),
			FindOneShoppingCart: query.NewFindOneShoppingCartHandler(
				shoppingcartRepository,
			),
		},
		Events: Events{
			ItemAddedToCart: event.NewItemAddedToCartHandler(
				pubSub,
			),

			ItemRemovedFromCart: event.NewItemRemovedFromCartHandler(
				pubSub,
			),

			SalesOrderPlaced: event.NewSalesOrderPlacedHandler(
				pubSub,
			),

			InventoryDeductionRequested: event.NewInventoryDeductionRequestedHandler(
				pubSub,
			),

			SalesOrderFulfilled: event.NewSalesOrderFulfilledHandler(
				pubSub,
			),

			SalesOrderCancelled: event.NewSalesOrderCancelledHandler(
				pubSub,
			),

			InventoryRestockRequested: event.NewInventoryRestockRequestedHandler(
				pubSub,
			),
		},
	}

	return app, nil
}
