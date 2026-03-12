package inventoryandproducts

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type ProductRepository interface {
	AddProduct(ctx context.Context, product *Product) error

	GetProduct(ctx context.Context, id uuid.UUID) (*Product, error)

	UpdateProduct(
		ctx context.Context,
		productId uuid.UUID,
		updateFn func(e *Product) (*Product, error),
	) error
}

// MapToProduct creates a Product domain object from the given parameters.
// This should ONLY BE USED when reconstructing an Product from its repository.
func MapToProduct(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//productId string,
	//
	//name string,
	//
	//description *string,
	//
	//category string,
	//
	//unitPrice float32,
	//
	//sku string,
	//
	//supplierId *string,
	//
	//isActive bool,
	//
) (*Product, error) {
	return &Product{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		//
		//productId: productId,
		//
		//name: name,
		//
		//description: description,
		//
		//category: category,
		//
		//unitPrice: unitPrice,
		//
		//sku: sku,
		//
		//supplierId: supplierId,
		//
		//isActive: isActive,
		//
		// TODO
	}, nil
}

type InventoryItemRepository interface {
	AddInventoryItem(ctx context.Context, inventoryitem *InventoryItem) error

	GetInventoryItem(ctx context.Context, id uuid.UUID) (*InventoryItem, error)

	UpdateInventoryItem(
		ctx context.Context,
		inventoryitemId uuid.UUID,
		updateFn func(e *InventoryItem) (*InventoryItem, error),
	) error
}

// MapToInventoryItem creates a InventoryItem domain object from the given parameters.
// This should ONLY BE USED when reconstructing an InventoryItem from its repository.
func MapToInventoryItem(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//inventoryItemId string,
	//
	//productId string,
	//
	//quantityOnHand int32,
	//
	//reorderThreshold int32,
	//
	//reorderQuantity int32,
	//
	//lastRestockedDate *string,
	//
) (*InventoryItem, error) {
	return &InventoryItem{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		//
		//inventoryItemId: inventoryItemId,
		//
		//productId: productId,
		//
		//quantityOnHand: quantityOnHand,
		//
		//reorderThreshold: reorderThreshold,
		//
		//reorderQuantity: reorderQuantity,
		//
		//lastRestockedDate: lastRestockedDate,
		//
		// TODO
	}, nil
}

type SupplierRepository interface {
	AddSupplier(ctx context.Context, supplier *Supplier) error

	GetSupplier(ctx context.Context, id uuid.UUID) (*Supplier, error)

	UpdateSupplier(
		ctx context.Context,
		supplierId uuid.UUID,
		updateFn func(e *Supplier) (*Supplier, error),
	) error
}

// MapToSupplier creates a Supplier domain object from the given parameters.
// This should ONLY BE USED when reconstructing an Supplier from its repository.
func MapToSupplier(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//supplierId string,
	//
	//name string,
	//
	//contactName *string,
	//
	//email string,
	//
	//phoneNumber string,
	//
	//productCategories string,
	//
	//status string,
	//
) (*Supplier, error) {
	return &Supplier{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		//
		//supplierId: supplierId,
		//
		//name: name,
		//
		//contactName: contactName,
		//
		//email: email,
		//
		//phoneNumber: phoneNumber,
		//
		//productCategories: productCategories,
		//
		//status: status,
		//
		// TODO
	}, nil
}

type PurchaseOrderRepository interface {
	AddPurchaseOrder(ctx context.Context, purchaseorder *PurchaseOrder) error

	GetPurchaseOrder(ctx context.Context, id uuid.UUID) (*PurchaseOrder, error)

	UpdatePurchaseOrder(
		ctx context.Context,
		purchaseorderId uuid.UUID,
		updateFn func(e *PurchaseOrder) (*PurchaseOrder, error),
	) error
}

// MapToPurchaseOrder creates a PurchaseOrder domain object from the given parameters.
// This should ONLY BE USED when reconstructing an PurchaseOrder from its repository.
func MapToPurchaseOrder(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//purchaseOrderId string,
	//
	//supplierId string,
	//
	//orderDate string,
	//
	//expectedDeliveryDate *string,
	//
	//lineItems string,
	//
	//totalCost float32,
	//
	//status string,
	//
) (*PurchaseOrder, error) {
	return &PurchaseOrder{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		//
		//purchaseOrderId: purchaseOrderId,
		//
		//supplierId: supplierId,
		//
		//orderDate: orderDate,
		//
		//expectedDeliveryDate: expectedDeliveryDate,
		//
		//lineItems: lineItems,
		//
		//totalCost: totalCost,
		//
		//status: status,
		//
		// TODO
	}, nil
}
