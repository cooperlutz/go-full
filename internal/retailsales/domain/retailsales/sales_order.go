package retailsales

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type SalesOrder struct {
	*baseentitee.EntityMetadata
	//
	//orderId string
	//
	//ownerId string
	//
	//orderDate string
	//
	//lineItems string
	//
	//subtotal float32
	//
	//discountAmount *float32
	//
	//taxAmount float32
	//
	//totalAmount float32
	//
	//channel string
	//
	//status string
	//
	// TODO
}

func NewSalesOrder(
// orderId string,
//
// ownerId string,
//
// orderDate string,
//
// lineItems string,
//
// subtotal float32,
//
// discountAmount *float32,
//
// taxAmount float32,
//
// totalAmount float32,
//
// channel string,
//
// status string,
) *SalesOrder {
	return &SalesOrder{
		EntityMetadata: baseentitee.NewEntityMetadata(),
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
	}
}
