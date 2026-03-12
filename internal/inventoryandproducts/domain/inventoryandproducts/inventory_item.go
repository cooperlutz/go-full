package inventoryandproducts

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type InventoryItem struct {
	*baseentitee.EntityMetadata
	//
	//inventoryItemId string
	//
	//productId string
	//
	//quantityOnHand int32
	//
	//reorderThreshold int32
	//
	//reorderQuantity int32
	//
	//lastRestockedDate *string
	//
	// TODO
}

func NewInventoryItem(
// inventoryItemId string,
//
// productId string,
//
// quantityOnHand int32,
//
// reorderThreshold int32,
//
// reorderQuantity int32,
//
// lastRestockedDate *string,
) *InventoryItem {
	return &InventoryItem{
		EntityMetadata: baseentitee.NewEntityMetadata(),
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
	}
}
