package billingandpayments

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type Refund struct {
	*baseentitee.EntityMetadata
	//
	//refundId string
	//
	//paymentId string
	//
	//ownerId string
	//
	//refundAmount float32
	//
	//reason string
	//
	//refundDate string
	//
	//status string
	//
	// TODO
}

func NewRefund(
// refundId string,
//
// paymentId string,
//
// ownerId string,
//
// refundAmount float32,
//
// reason string,
//
// refundDate string,
//
// status string,
) *Refund {
	return &Refund{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		//
		//refundId: refundId,
		//
		//paymentId: paymentId,
		//
		//ownerId: ownerId,
		//
		//refundAmount: refundAmount,
		//
		//reason: reason,
		//
		//refundDate: refundDate,
		//
		//status: status,
		//
	}
}
