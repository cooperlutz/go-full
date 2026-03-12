package inventoryandproducts

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type PurchaseOrder struct {
	*baseentitee.EntityMetadata
	//
	//purchaseOrderId string
	//
	//supplierId string
	//
	//orderDate string
	//
	//expectedDeliveryDate *string
	//
	//lineItems string
	//
	//totalCost float32
	//
	//status string
	//
	// TODO
}

func NewPurchaseOrder(
// purchaseOrderId string,
//
// supplierId string,
//
// orderDate string,
//
// expectedDeliveryDate *string,
//
// lineItems string,
//
// totalCost float32,
//
// status string,
) *PurchaseOrder {
	return &PurchaseOrder{
		EntityMetadata: baseentitee.NewEntityMetadata(),
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
	}
}
