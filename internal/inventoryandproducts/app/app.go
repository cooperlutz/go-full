package app

import (
	"github.com/cooperlutz/go-full/internal/inventoryandproducts/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/inventoryandproducts/app/command"
	"github.com/cooperlutz/go-full/internal/inventoryandproducts/app/event"
	"github.com/cooperlutz/go-full/internal/inventoryandproducts/app/query"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Events   Events
}

type Commands struct {
	AddProductToCatalog command.AddProductToCatalogHandler

	UpdateProductPrice command.UpdateProductPriceHandler

	RestockInventory command.RestockInventoryHandler

	DeductInventory command.DeductInventoryHandler

	PlacePurchaseOrder command.PlacePurchaseOrderHandler

	ReceivePurchaseOrder command.ReceivePurchaseOrderHandler
}

type Queries struct {
	FindAllProducts query.FindAllProductsHandler
	FindOneProduct  query.FindOneProductHandler

	FindAllInventoryItems query.FindAllInventoryItemsHandler
	FindOneInventoryItem  query.FindOneInventoryItemHandler

	FindAllSuppliers query.FindAllSuppliersHandler
	FindOneSupplier  query.FindOneSupplierHandler

	FindAllPurchaseOrders query.FindAllPurchaseOrdersHandler
	FindOnePurchaseOrder  query.FindOnePurchaseOrderHandler
}

type Events struct {
	ProductAddedToCatalog event.ProductAddedToCatalogHandler

	ProductPriceUpdated event.ProductPriceUpdatedHandler

	InventoryRestocked event.InventoryRestockedHandler

	InventoryDeducted event.InventoryDeductedHandler

	LowStockAlertTriggered event.LowStockAlertTriggeredHandler

	PurchaseOrderPlaced event.PurchaseOrderPlacedHandler

	PurchaseOrderReceived event.PurchaseOrderReceivedHandler
}

// NewApplication initializes the InventoryAndProducts application with its dependencies.
func NewApplication( //nolint:funlen // it's fine
	pgconn deebee.IDatabase,
	pubSub eeventdriven.IPubSubEventProcessor,
) (Application, error) {
	productRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	inventoryitemRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	supplierRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	purchaseorderRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	app := Application{
		Commands: Commands{
			AddProductToCatalog: command.NewAddProductToCatalogHandler(

				productRepository,

				inventoryitemRepository,

				supplierRepository,

				purchaseorderRepository,
			),
			UpdateProductPrice: command.NewUpdateProductPriceHandler(

				productRepository,

				inventoryitemRepository,

				supplierRepository,

				purchaseorderRepository,
			),
			RestockInventory: command.NewRestockInventoryHandler(

				productRepository,

				inventoryitemRepository,

				supplierRepository,

				purchaseorderRepository,
			),
			DeductInventory: command.NewDeductInventoryHandler(

				productRepository,

				inventoryitemRepository,

				supplierRepository,

				purchaseorderRepository,
			),
			PlacePurchaseOrder: command.NewPlacePurchaseOrderHandler(

				productRepository,

				inventoryitemRepository,

				supplierRepository,

				purchaseorderRepository,
			),
			ReceivePurchaseOrder: command.NewReceivePurchaseOrderHandler(

				productRepository,

				inventoryitemRepository,

				supplierRepository,

				purchaseorderRepository,
			),
		},
		Queries: Queries{
			FindAllProducts: query.NewFindAllProductsHandler(
				productRepository,
			),
			FindOneProduct: query.NewFindOneProductHandler(
				productRepository,
			),

			FindAllInventoryItems: query.NewFindAllInventoryItemsHandler(
				inventoryitemRepository,
			),
			FindOneInventoryItem: query.NewFindOneInventoryItemHandler(
				inventoryitemRepository,
			),

			FindAllSuppliers: query.NewFindAllSuppliersHandler(
				supplierRepository,
			),
			FindOneSupplier: query.NewFindOneSupplierHandler(
				supplierRepository,
			),

			FindAllPurchaseOrders: query.NewFindAllPurchaseOrdersHandler(
				purchaseorderRepository,
			),
			FindOnePurchaseOrder: query.NewFindOnePurchaseOrderHandler(
				purchaseorderRepository,
			),
		},
		Events: Events{
			ProductAddedToCatalog: event.NewProductAddedToCatalogHandler(
				pubSub,
			),

			ProductPriceUpdated: event.NewProductPriceUpdatedHandler(
				pubSub,
			),

			InventoryRestocked: event.NewInventoryRestockedHandler(
				pubSub,
			),

			InventoryDeducted: event.NewInventoryDeductedHandler(
				pubSub,
			),

			LowStockAlertTriggered: event.NewLowStockAlertTriggeredHandler(
				pubSub,
			),

			PurchaseOrderPlaced: event.NewPurchaseOrderPlacedHandler(
				pubSub,
			),

			PurchaseOrderReceived: event.NewPurchaseOrderReceivedHandler(
				pubSub,
			),
		},
	}

	return app, nil
}
