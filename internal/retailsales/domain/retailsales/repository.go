package retailsales

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type SalesOrderRepository interface {
	AddSalesOrder(ctx context.Context, salesorder *SalesOrder) error

	GetSalesOrder(ctx context.Context, id uuid.UUID) (*SalesOrder, error)

	UpdateSalesOrder(
		ctx context.Context,
		salesorderId uuid.UUID,
		updateFn func(e *SalesOrder) (*SalesOrder, error),
	) error
}

// MapToSalesOrder creates a SalesOrder domain object from the given parameters.
// This should ONLY BE USED when reconstructing an SalesOrder from its repository.
func MapToSalesOrder(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//orderId string,
	//
	//ownerId string,
	//
	//orderDate string,
	//
	//lineItems string,
	//
	//subtotal float32,
	//
	//discountAmount *float32,
	//
	//taxAmount float32,
	//
	//totalAmount float32,
	//
	//channel string,
	//
	//status string,
	//
) (*SalesOrder, error) {
	return &SalesOrder{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		//
		//orderId: orderId,
		//
		//ownerId: ownerId,
		//
		//orderDate: orderDate,
		//
		//lineItems: lineItems,
		//
		//subtotal: subtotal,
		//
		//discountAmount: discountAmount,
		//
		//taxAmount: taxAmount,
		//
		//totalAmount: totalAmount,
		//
		//channel: channel,
		//
		//status: status,
		//
		// TODO
	}, nil
}

type ShoppingCartRepository interface {
	AddShoppingCart(ctx context.Context, shoppingcart *ShoppingCart) error

	GetShoppingCart(ctx context.Context, id uuid.UUID) (*ShoppingCart, error)

	UpdateShoppingCart(
		ctx context.Context,
		shoppingcartId uuid.UUID,
		updateFn func(e *ShoppingCart) (*ShoppingCart, error),
	) error
}

// MapToShoppingCart creates a ShoppingCart domain object from the given parameters.
// This should ONLY BE USED when reconstructing an ShoppingCart from its repository.
func MapToShoppingCart(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//cartId string,
	//
	//ownerId string,
	//
	//items string,
	//
	//createdAt string,
	//
	//updatedAt string,
	//
) (*ShoppingCart, error) {
	return &ShoppingCart{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		//
		//cartId: cartId,
		//
		//ownerId: ownerId,
		//
		//items: items,
		//
		//createdAt: createdAt,
		//
		//updatedAt: updatedAt,
		//
		// TODO
	}, nil
}
