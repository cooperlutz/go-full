package inventoryandproducts

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type Product struct {
	*baseentitee.EntityMetadata
	//
	//productId string
	//
	//name string
	//
	//description *string
	//
	//category string
	//
	//unitPrice float32
	//
	//sku string
	//
	//supplierId *string
	//
	//isActive bool
	//
	// TODO
}

func NewProduct(
// productId string,
//
// name string,
//
// description *string,
//
// category string,
//
// unitPrice float32,
//
// sku string,
//
// supplierId *string,
//
// isActive bool,
) *Product {
	return &Product{
		EntityMetadata: baseentitee.NewEntityMetadata(),
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
	}
}
