package retailsales

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type ShoppingCart struct {
	*baseentitee.EntityMetadata
	//
	//cartId string
	//
	//ownerId string
	//
	//items string
	//
	//createdAt string
	//
	//updatedAt string
	//
	// TODO
}

func NewShoppingCart(
// cartId string,
//
// ownerId string,
//
// items string,
//
// createdAt string,
//
// updatedAt string,
) *ShoppingCart {
	return &ShoppingCart{
		EntityMetadata: baseentitee.NewEntityMetadata(),
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
	}
}
