package billingandpayments

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type Invoice struct {
	*baseentitee.EntityMetadata
	//
	//invoiceId string
	//
	//ownerId string
	//
	//petId *string
	//
	//appointmentId *string
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
	//status string
	//
	//issuedDate string
	//
	//dueDate string
	//
	// TODO
}

func NewInvoice(
// invoiceId string,
//
// ownerId string,
//
// petId *string,
//
// appointmentId *string,
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
// status string,
//
// issuedDate string,
//
// dueDate string,
) *Invoice {
	return &Invoice{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		//
		//invoiceId: invoiceId,
		//
		//ownerId: ownerId,
		//
		//petId: petId,
		//
		//appointmentId: appointmentId,
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
		//status: status,
		//
		//issuedDate: issuedDate,
		//
		//dueDate: dueDate,
		//
	}
}
