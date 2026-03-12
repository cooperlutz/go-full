package inbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/retailsales/app"
	"github.com/cooperlutz/go-full/internal/retailsales/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// HttpAdapter represents the HTTP server for the RetailSales module.
type HttpAdapter struct {
	app app.Application
}

// NewHttpAdapter creates a new HttpAdapter instance with the provided RetailSales application.
func NewHttpAdapter(application app.Application) HttpAdapter {
	return HttpAdapter{
		app: application,
	}
}

// StrictHandler returns a strict HTTP handler for the RetailSales module.
func (h HttpAdapter) StrictHandler() ServerInterface {
	return NewStrictHandler(h, nil)
}

// (GET /v1/salesorders).
func (h HttpAdapter) FindAllSalesOrders(ctx context.Context, request FindAllSalesOrdersRequestObject) (FindAllSalesOrdersResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "salesorder.adapters.inbound.http.find_all_salesorders")
	defer span.End()

	salesorder, err := h.app.Queries.FindAllSalesOrders.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseSalesOrders []SalesOrder
	for _, e := range salesorder {
		responseSalesOrders = append(responseSalesOrders, querySalesOrderToHttpSalesOrder(e))
	}

	return FindAllSalesOrders200JSONResponse(responseSalesOrders), nil
}

// (GET /v1/salesorder/{sales_orderId}).
func (h HttpAdapter) FindOneSalesOrder(ctx context.Context, request FindOneSalesOrderRequestObject) (FindOneSalesOrderResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_sales_order")
	defer span.End()

	salesorder, err := h.app.Queries.FindOneSalesOrder.Handle(ctx, query.FindOneSalesOrder{SalesOrderID: request.SalesOrderId})
	if err != nil {
		return nil, err
	}

	return FindOneSalesOrder200JSONResponse(querySalesOrderToHttpSalesOrder(salesorder)), nil
}

func querySalesOrderToHttpSalesOrder(e query.SalesOrder) SalesOrder {
	return SalesOrder{
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

// (GET /v1/shoppingcarts).
func (h HttpAdapter) FindAllShoppingCarts(ctx context.Context, request FindAllShoppingCartsRequestObject) (FindAllShoppingCartsResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "shoppingcart.adapters.inbound.http.find_all_shoppingcarts")
	defer span.End()

	shoppingcart, err := h.app.Queries.FindAllShoppingCarts.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseShoppingCarts []ShoppingCart
	for _, e := range shoppingcart {
		responseShoppingCarts = append(responseShoppingCarts, queryShoppingCartToHttpShoppingCart(e))
	}

	return FindAllShoppingCarts200JSONResponse(responseShoppingCarts), nil
}

// (GET /v1/shoppingcart/{shopping_cartId}).
func (h HttpAdapter) FindOneShoppingCart(ctx context.Context, request FindOneShoppingCartRequestObject) (FindOneShoppingCartResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_shopping_cart")
	defer span.End()

	shoppingcart, err := h.app.Queries.FindOneShoppingCart.Handle(ctx, query.FindOneShoppingCart{ShoppingCartID: request.ShoppingCartId})
	if err != nil {
		return nil, err
	}

	return FindOneShoppingCart200JSONResponse(queryShoppingCartToHttpShoppingCart(shoppingcart)), nil
}

func queryShoppingCartToHttpShoppingCart(e query.ShoppingCart) ShoppingCart {
	return ShoppingCart{
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
