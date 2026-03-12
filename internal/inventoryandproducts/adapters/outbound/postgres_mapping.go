package outbound

import (
	"github.com/cooperlutz/go-full/internal/inventoryandproducts/app/query"
	"github.com/cooperlutz/go-full/internal/inventoryandproducts/domain/inventoryandproducts"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

// toDomain maps the ProductProduct to the domain entity.
func (e InventoryandproductsProduct) toDomain() (*inventoryandproducts.Product, error) {
	return inventoryandproducts.MapToProduct(
		e.ProductID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.ProductId,
		//
		//e.Name,
		//
		//e.Description,
		//
		//e.Category,
		//
		//e.UnitPrice,
		//
		//e.Sku,
		//
		//e.SupplierId,
		//
		//e.IsActive,
		//
		// TODO
	)
}

// toQueryProduct maps the productProduct to the query.Product.
func (e InventoryandproductsProduct) toQueryProduct() (query.Product, error) {
	product, err := e.toDomain()
	if err != nil {
		return query.Product{}, err
	}

	return mapEntityProductToQuery(product), nil
}

// productProductsToQuery maps a slice of ProductProduct to a slice of query.Product entities.
func inventoryandproductsProductsToQuery(products []InventoryandproductsProduct) ([]query.Product, error) {
	var domainProducts []query.Product

	for _, product := range products {
		queryProduct, err := product.toQueryProduct()
		if err != nil {
			return nil, err
		}

		domainProducts = append(domainProducts, queryProduct)
	}

	return domainProducts, nil
}

// mapEntityProductToDB maps a domain Product entity to the ProductProduct database model.
func mapEntityProductToDB(product *inventoryandproducts.Product) InventoryandproductsProduct {
	createdAt := product.GetCreatedAtTime()
	updatedAt := product.GetUpdatedAtTime()

	return InventoryandproductsProduct{
		ProductID: pgxutil.UUIDToPgtypeUUID(product.GetIdUUID()),
		CreatedAt: pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt: pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:   product.IsDeleted(),
		DeletedAt: pgxutil.TimeToTimestampz(product.GetDeletedAtTime()),
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

// mapEntityProductToQuery maps a domain Product entity to a query.Product.
func mapEntityProductToQuery(product *inventoryandproducts.Product) query.Product {
	return query.Product{
		// TODO
	}
}

// toDomain maps the InventoryitemInventoryItem to the domain entity.
func (e InventoryandproductsInventoryItem) toDomain() (*inventoryandproducts.InventoryItem, error) {
	return inventoryandproducts.MapToInventoryItem(
		e.InventoryItemID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.InventoryItemId,
		//
		//e.ProductId,
		//
		//e.QuantityOnHand,
		//
		//e.ReorderThreshold,
		//
		//e.ReorderQuantity,
		//
		//e.LastRestockedDate,
		//
		// TODO
	)
}

// toQueryInventoryItem maps the inventoryitemInventoryItem to the query.InventoryItem.
func (e InventoryandproductsInventoryItem) toQueryInventoryItem() (query.InventoryItem, error) {
	inventoryitem, err := e.toDomain()
	if err != nil {
		return query.InventoryItem{}, err
	}

	return mapEntityInventoryItemToQuery(inventoryitem), nil
}

// inventoryitemInventoryItemsToQuery maps a slice of InventoryItemInventoryItem to a slice of query.InventoryItem entities.
func inventoryandproductsInventoryItemsToQuery(inventoryitems []InventoryandproductsInventoryItem) ([]query.InventoryItem, error) {
	var domainInventoryItems []query.InventoryItem

	for _, inventoryitem := range inventoryitems {
		queryInventoryItem, err := inventoryitem.toQueryInventoryItem()
		if err != nil {
			return nil, err
		}

		domainInventoryItems = append(domainInventoryItems, queryInventoryItem)
	}

	return domainInventoryItems, nil
}

// mapEntityInventoryItemToDB maps a domain InventoryItem entity to the InventoryItemInventoryItem database model.
func mapEntityInventoryItemToDB(inventoryitem *inventoryandproducts.InventoryItem) InventoryandproductsInventoryItem {
	createdAt := inventoryitem.GetCreatedAtTime()
	updatedAt := inventoryitem.GetUpdatedAtTime()

	return InventoryandproductsInventoryItem{
		InventoryItemID: pgxutil.UUIDToPgtypeUUID(inventoryitem.GetIdUUID()),
		CreatedAt:       pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:       pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:         inventoryitem.IsDeleted(),
		DeletedAt:       pgxutil.TimeToTimestampz(inventoryitem.GetDeletedAtTime()),
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

// mapEntityInventoryItemToQuery maps a domain InventoryItem entity to a query.InventoryItem.
func mapEntityInventoryItemToQuery(inventoryitem *inventoryandproducts.InventoryItem) query.InventoryItem {
	return query.InventoryItem{
		// TODO
	}
}

// toDomain maps the SupplierSupplier to the domain entity.
func (e InventoryandproductsSupplier) toDomain() (*inventoryandproducts.Supplier, error) {
	return inventoryandproducts.MapToSupplier(
		e.SupplierID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.SupplierId,
		//
		//e.Name,
		//
		//e.ContactName,
		//
		//e.Email,
		//
		//e.PhoneNumber,
		//
		//e.ProductCategories,
		//
		//e.Status,
		//
		// TODO
	)
}

// toQuerySupplier maps the supplierSupplier to the query.Supplier.
func (e InventoryandproductsSupplier) toQuerySupplier() (query.Supplier, error) {
	supplier, err := e.toDomain()
	if err != nil {
		return query.Supplier{}, err
	}

	return mapEntitySupplierToQuery(supplier), nil
}

// supplierSuppliersToQuery maps a slice of SupplierSupplier to a slice of query.Supplier entities.
func inventoryandproductsSuppliersToQuery(suppliers []InventoryandproductsSupplier) ([]query.Supplier, error) {
	var domainSuppliers []query.Supplier

	for _, supplier := range suppliers {
		querySupplier, err := supplier.toQuerySupplier()
		if err != nil {
			return nil, err
		}

		domainSuppliers = append(domainSuppliers, querySupplier)
	}

	return domainSuppliers, nil
}

// mapEntitySupplierToDB maps a domain Supplier entity to the SupplierSupplier database model.
func mapEntitySupplierToDB(supplier *inventoryandproducts.Supplier) InventoryandproductsSupplier {
	createdAt := supplier.GetCreatedAtTime()
	updatedAt := supplier.GetUpdatedAtTime()

	return InventoryandproductsSupplier{
		SupplierID: pgxutil.UUIDToPgtypeUUID(supplier.GetIdUUID()),
		CreatedAt:  pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:  pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:    supplier.IsDeleted(),
		DeletedAt:  pgxutil.TimeToTimestampz(supplier.GetDeletedAtTime()),
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

// mapEntitySupplierToQuery maps a domain Supplier entity to a query.Supplier.
func mapEntitySupplierToQuery(supplier *inventoryandproducts.Supplier) query.Supplier {
	return query.Supplier{
		// TODO
	}
}

// toDomain maps the PurchaseorderPurchaseOrder to the domain entity.
func (e InventoryandproductsPurchaseOrder) toDomain() (*inventoryandproducts.PurchaseOrder, error) {
	return inventoryandproducts.MapToPurchaseOrder(
		e.PurchaseOrderID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.PurchaseOrderId,
		//
		//e.SupplierId,
		//
		//e.OrderDate,
		//
		//e.ExpectedDeliveryDate,
		//
		//e.LineItems,
		//
		//e.TotalCost,
		//
		//e.Status,
		//
		// TODO
	)
}

// toQueryPurchaseOrder maps the purchaseorderPurchaseOrder to the query.PurchaseOrder.
func (e InventoryandproductsPurchaseOrder) toQueryPurchaseOrder() (query.PurchaseOrder, error) {
	purchaseorder, err := e.toDomain()
	if err != nil {
		return query.PurchaseOrder{}, err
	}

	return mapEntityPurchaseOrderToQuery(purchaseorder), nil
}

// purchaseorderPurchaseOrdersToQuery maps a slice of PurchaseOrderPurchaseOrder to a slice of query.PurchaseOrder entities.
func inventoryandproductsPurchaseOrdersToQuery(purchaseorders []InventoryandproductsPurchaseOrder) ([]query.PurchaseOrder, error) {
	var domainPurchaseOrders []query.PurchaseOrder

	for _, purchaseorder := range purchaseorders {
		queryPurchaseOrder, err := purchaseorder.toQueryPurchaseOrder()
		if err != nil {
			return nil, err
		}

		domainPurchaseOrders = append(domainPurchaseOrders, queryPurchaseOrder)
	}

	return domainPurchaseOrders, nil
}

// mapEntityPurchaseOrderToDB maps a domain PurchaseOrder entity to the PurchaseOrderPurchaseOrder database model.
func mapEntityPurchaseOrderToDB(purchaseorder *inventoryandproducts.PurchaseOrder) InventoryandproductsPurchaseOrder {
	createdAt := purchaseorder.GetCreatedAtTime()
	updatedAt := purchaseorder.GetUpdatedAtTime()

	return InventoryandproductsPurchaseOrder{
		PurchaseOrderID: pgxutil.UUIDToPgtypeUUID(purchaseorder.GetIdUUID()),
		CreatedAt:       pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:       pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:         purchaseorder.IsDeleted(),
		DeletedAt:       pgxutil.TimeToTimestampz(purchaseorder.GetDeletedAtTime()),
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

// mapEntityPurchaseOrderToQuery maps a domain PurchaseOrder entity to a query.PurchaseOrder.
func mapEntityPurchaseOrderToQuery(purchaseorder *inventoryandproducts.PurchaseOrder) query.PurchaseOrder {
	return query.PurchaseOrder{
		// TODO
	}
}
