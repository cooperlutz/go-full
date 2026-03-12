package outbound

import (
	"github.com/cooperlutz/go-full/internal/retailsales/app/query"
	"github.com/cooperlutz/go-full/internal/retailsales/domain/retailsales"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

// toDomain maps the SalesorderSalesOrder to the domain entity.
func (e RetailsalesSalesOrder) toDomain() (*retailsales.SalesOrder, error) {
	return retailsales.MapToSalesOrder(
		e.SalesOrderID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.OrderId,
		//
		//e.OwnerId,
		//
		//e.OrderDate,
		//
		//e.LineItems,
		//
		//e.Subtotal,
		//
		//e.DiscountAmount,
		//
		//e.TaxAmount,
		//
		//e.TotalAmount,
		//
		//e.Channel,
		//
		//e.Status,
		//
		// TODO
	)
}

// toQuerySalesOrder maps the salesorderSalesOrder to the query.SalesOrder.
func (e RetailsalesSalesOrder) toQuerySalesOrder() (query.SalesOrder, error) {
	salesorder, err := e.toDomain()
	if err != nil {
		return query.SalesOrder{}, err
	}

	return mapEntitySalesOrderToQuery(salesorder), nil
}

// salesorderSalesOrdersToQuery maps a slice of SalesOrderSalesOrder to a slice of query.SalesOrder entities.
func retailsalesSalesOrdersToQuery(salesorders []RetailsalesSalesOrder) ([]query.SalesOrder, error) {
	var domainSalesOrders []query.SalesOrder

	for _, salesorder := range salesorders {
		querySalesOrder, err := salesorder.toQuerySalesOrder()
		if err != nil {
			return nil, err
		}

		domainSalesOrders = append(domainSalesOrders, querySalesOrder)
	}

	return domainSalesOrders, nil
}

// mapEntitySalesOrderToDB maps a domain SalesOrder entity to the SalesOrderSalesOrder database model.
func mapEntitySalesOrderToDB(salesorder *retailsales.SalesOrder) RetailsalesSalesOrder {
	createdAt := salesorder.GetCreatedAtTime()
	updatedAt := salesorder.GetUpdatedAtTime()

	return RetailsalesSalesOrder{
		SalesOrderID: pgxutil.UUIDToPgtypeUUID(salesorder.GetIdUUID()),
		CreatedAt:    pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:    pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:      salesorder.IsDeleted(),
		DeletedAt:    pgxutil.TimeToTimestampz(salesorder.GetDeletedAtTime()),
		//
		//OrderId: GetOrderId(),
		//
		//OwnerId: GetOwnerId(),
		//
		//OrderDate: GetOrderDate(),
		//
		//LineItems: GetLineItems(),
		//
		//Subtotal: GetSubtotal(),
		//
		//DiscountAmount: GetDiscountAmount(),
		//
		//TaxAmount: GetTaxAmount(),
		//
		//TotalAmount: GetTotalAmount(),
		//
		//Channel: GetChannel(),
		//
		//Status: GetStatus(),
		//
		// TODO
	}
}

// mapEntitySalesOrderToQuery maps a domain SalesOrder entity to a query.SalesOrder.
func mapEntitySalesOrderToQuery(salesorder *retailsales.SalesOrder) query.SalesOrder {
	return query.SalesOrder{
		// TODO
	}
}

// toDomain maps the ShoppingcartShoppingCart to the domain entity.
func (e RetailsalesShoppingCart) toDomain() (*retailsales.ShoppingCart, error) {
	return retailsales.MapToShoppingCart(
		e.ShoppingCartID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.CartId,
		//
		//e.OwnerId,
		//
		//e.Items,
		//
		//e.CreatedAt,
		//
		//e.UpdatedAt,
		//
		// TODO
	)
}

// toQueryShoppingCart maps the shoppingcartShoppingCart to the query.ShoppingCart.
func (e RetailsalesShoppingCart) toQueryShoppingCart() (query.ShoppingCart, error) {
	shoppingcart, err := e.toDomain()
	if err != nil {
		return query.ShoppingCart{}, err
	}

	return mapEntityShoppingCartToQuery(shoppingcart), nil
}

// shoppingcartShoppingCartsToQuery maps a slice of ShoppingCartShoppingCart to a slice of query.ShoppingCart entities.
func retailsalesShoppingCartsToQuery(shoppingcarts []RetailsalesShoppingCart) ([]query.ShoppingCart, error) {
	var domainShoppingCarts []query.ShoppingCart

	for _, shoppingcart := range shoppingcarts {
		queryShoppingCart, err := shoppingcart.toQueryShoppingCart()
		if err != nil {
			return nil, err
		}

		domainShoppingCarts = append(domainShoppingCarts, queryShoppingCart)
	}

	return domainShoppingCarts, nil
}

// mapEntityShoppingCartToDB maps a domain ShoppingCart entity to the ShoppingCartShoppingCart database model.
func mapEntityShoppingCartToDB(shoppingcart *retailsales.ShoppingCart) RetailsalesShoppingCart {
	createdAt := shoppingcart.GetCreatedAtTime()
	updatedAt := shoppingcart.GetUpdatedAtTime()

	return RetailsalesShoppingCart{
		ShoppingCartID: pgxutil.UUIDToPgtypeUUID(shoppingcart.GetIdUUID()),
		CreatedAt:      pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:      pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:        shoppingcart.IsDeleted(),
		DeletedAt:      pgxutil.TimeToTimestampz(shoppingcart.GetDeletedAtTime()),
		//
		//CartId: GetCartId(),
		//
		//OwnerId: GetOwnerId(),
		//
		//Items: GetItems(),
		//
		//CreatedAt: GetCreatedAt(),
		//
		//UpdatedAt: GetUpdatedAt(),
		//
		// TODO
	}
}

// mapEntityShoppingCartToQuery maps a domain ShoppingCart entity to a query.ShoppingCart.
func mapEntityShoppingCartToQuery(shoppingcart *retailsales.ShoppingCart) query.ShoppingCart {
	return query.ShoppingCart{
		// TODO
	}
}
