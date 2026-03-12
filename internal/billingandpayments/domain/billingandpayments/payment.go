package billingandpayments

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type Payment struct {
	*baseentitee.EntityMetadata
	//
	//paymentId string
	//
	//invoiceId string
	//
	//ownerId string
	//
	//amountPaid float32
	//
	//paymentMethod string
	//
	//paymentDate string
	//
	//status string
	//
	// TODO
}

func NewPayment(
// paymentId string,
//
// invoiceId string,
//
// ownerId string,
//
// amountPaid float32,
//
// paymentMethod string,
//
// paymentDate string,
//
// status string,
) *Payment {
	return &Payment{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		//
		//paymentId: paymentId,
		//
		//invoiceId: invoiceId,
		//
		//ownerId: ownerId,
		//
		//amountPaid: amountPaid,
		//
		//paymentMethod: paymentMethod,
		//
		//paymentDate: paymentDate,
		//
		//status: status,
		//
	}
}
