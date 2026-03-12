package inventoryandproducts

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type Supplier struct {
	*baseentitee.EntityMetadata
	//
	//supplierId string
	//
	//name string
	//
	//contactName *string
	//
	//email string
	//
	//phoneNumber string
	//
	//productCategories string
	//
	//status string
	//
	// TODO
}

func NewSupplier(
// supplierId string,
//
// name string,
//
// contactName *string,
//
// email string,
//
// phoneNumber string,
//
// productCategories string,
//
// status string,
) *Supplier {
	return &Supplier{
		EntityMetadata: baseentitee.NewEntityMetadata(),
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
	}
}
