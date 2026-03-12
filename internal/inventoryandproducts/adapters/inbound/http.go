package inbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/inventoryandproducts/app"
	"github.com/cooperlutz/go-full/internal/inventoryandproducts/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// HttpAdapter represents the HTTP server for the InventoryAndProducts module.
type HttpAdapter struct {
	app app.Application
}

// NewHttpAdapter creates a new HttpAdapter instance with the provided InventoryAndProducts application.
func NewHttpAdapter(application app.Application) HttpAdapter {
	return HttpAdapter{
		app: application,
	}
}

// StrictHandler returns a strict HTTP handler for the InventoryAndProducts module.
func (h HttpAdapter) StrictHandler() ServerInterface {
	return NewStrictHandler(h, nil)
}

// (GET /v1/products).
func (h HttpAdapter) FindAllProducts(ctx context.Context, request FindAllProductsRequestObject) (FindAllProductsResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "product.adapters.inbound.http.find_all_products")
	defer span.End()

	product, err := h.app.Queries.FindAllProducts.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseProducts []Product
	for _, e := range product {
		responseProducts = append(responseProducts, queryProductToHttpProduct(e))
	}

	return FindAllProducts200JSONResponse(responseProducts), nil
}

// (GET /v1/product/{productId}).
func (h HttpAdapter) FindOneProduct(ctx context.Context, request FindOneProductRequestObject) (FindOneProductResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_product")
	defer span.End()

	product, err := h.app.Queries.FindOneProduct.Handle(ctx, query.FindOneProduct{ProductID: request.ProductId})
	if err != nil {
		return nil, err
	}

	return FindOneProduct200JSONResponse(queryProductToHttpProduct(product)), nil
}

func queryProductToHttpProduct(e query.Product) Product {
	return Product{
		//
		//ProductId: GetProductId(),
		//
		//Name: GetName(),
		//
		//Description: GetDescription(),
		//
		//Category: GetCategory(),
		//
		//UnitPrice: GetUnitPrice(),
		//
		//Sku: GetSku(),
		//
		//SupplierId: GetSupplierId(),
		//
		//IsActive: GetIsActive(),
		//
		// TODO
	}
}

// (GET /v1/inventoryitems).
func (h HttpAdapter) FindAllInventoryItems(ctx context.Context, request FindAllInventoryItemsRequestObject) (FindAllInventoryItemsResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryitem.adapters.inbound.http.find_all_inventoryitems")
	defer span.End()

	inventoryitem, err := h.app.Queries.FindAllInventoryItems.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseInventoryItems []InventoryItem
	for _, e := range inventoryitem {
		responseInventoryItems = append(responseInventoryItems, queryInventoryItemToHttpInventoryItem(e))
	}

	return FindAllInventoryItems200JSONResponse(responseInventoryItems), nil
}

// (GET /v1/inventoryitem/{inventory_itemId}).
func (h HttpAdapter) FindOneInventoryItem(ctx context.Context, request FindOneInventoryItemRequestObject) (FindOneInventoryItemResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_inventory_item")
	defer span.End()

	inventoryitem, err := h.app.Queries.FindOneInventoryItem.Handle(ctx, query.FindOneInventoryItem{InventoryItemID: request.InventoryItemId})
	if err != nil {
		return nil, err
	}

	return FindOneInventoryItem200JSONResponse(queryInventoryItemToHttpInventoryItem(inventoryitem)), nil
}

func queryInventoryItemToHttpInventoryItem(e query.InventoryItem) InventoryItem {
	return InventoryItem{
		//
		//InventoryItemId: GetInventoryItemId(),
		//
		//ProductId: GetProductId(),
		//
		//QuantityOnHand: GetQuantityOnHand(),
		//
		//ReorderThreshold: GetReorderThreshold(),
		//
		//ReorderQuantity: GetReorderQuantity(),
		//
		//LastRestockedDate: GetLastRestockedDate(),
		//
		// TODO
	}
}

// (GET /v1/suppliers).
func (h HttpAdapter) FindAllSuppliers(ctx context.Context, request FindAllSuppliersRequestObject) (FindAllSuppliersResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "supplier.adapters.inbound.http.find_all_suppliers")
	defer span.End()

	supplier, err := h.app.Queries.FindAllSuppliers.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseSuppliers []Supplier
	for _, e := range supplier {
		responseSuppliers = append(responseSuppliers, querySupplierToHttpSupplier(e))
	}

	return FindAllSuppliers200JSONResponse(responseSuppliers), nil
}

// (GET /v1/supplier/{supplierId}).
func (h HttpAdapter) FindOneSupplier(ctx context.Context, request FindOneSupplierRequestObject) (FindOneSupplierResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_supplier")
	defer span.End()

	supplier, err := h.app.Queries.FindOneSupplier.Handle(ctx, query.FindOneSupplier{SupplierID: request.SupplierId})
	if err != nil {
		return nil, err
	}

	return FindOneSupplier200JSONResponse(querySupplierToHttpSupplier(supplier)), nil
}

func querySupplierToHttpSupplier(e query.Supplier) Supplier {
	return Supplier{
		//
		//SupplierId: GetSupplierId(),
		//
		//Name: GetName(),
		//
		//ContactName: GetContactName(),
		//
		//Email: GetEmail(),
		//
		//PhoneNumber: GetPhoneNumber(),
		//
		//ProductCategories: GetProductCategories(),
		//
		//Status: GetStatus(),
		//
		// TODO
	}
}

// (GET /v1/purchaseorders).
func (h HttpAdapter) FindAllPurchaseOrders(ctx context.Context, request FindAllPurchaseOrdersRequestObject) (FindAllPurchaseOrdersResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "purchaseorder.adapters.inbound.http.find_all_purchaseorders")
	defer span.End()

	purchaseorder, err := h.app.Queries.FindAllPurchaseOrders.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responsePurchaseOrders []PurchaseOrder
	for _, e := range purchaseorder {
		responsePurchaseOrders = append(responsePurchaseOrders, queryPurchaseOrderToHttpPurchaseOrder(e))
	}

	return FindAllPurchaseOrders200JSONResponse(responsePurchaseOrders), nil
}

// (GET /v1/purchaseorder/{purchase_orderId}).
func (h HttpAdapter) FindOnePurchaseOrder(ctx context.Context, request FindOnePurchaseOrderRequestObject) (FindOnePurchaseOrderResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_purchase_order")
	defer span.End()

	purchaseorder, err := h.app.Queries.FindOnePurchaseOrder.Handle(ctx, query.FindOnePurchaseOrder{PurchaseOrderID: request.PurchaseOrderId})
	if err != nil {
		return nil, err
	}

	return FindOnePurchaseOrder200JSONResponse(queryPurchaseOrderToHttpPurchaseOrder(purchaseorder)), nil
}

func queryPurchaseOrderToHttpPurchaseOrder(e query.PurchaseOrder) PurchaseOrder {
	return PurchaseOrder{
		//
		//PurchaseOrderId: GetPurchaseOrderId(),
		//
		//SupplierId: GetSupplierId(),
		//
		//OrderDate: GetOrderDate(),
		//
		//ExpectedDeliveryDate: GetExpectedDeliveryDate(),
		//
		//LineItems: GetLineItems(),
		//
		//TotalCost: GetTotalCost(),
		//
		//Status: GetStatus(),
		//
		// TODO
	}
}
