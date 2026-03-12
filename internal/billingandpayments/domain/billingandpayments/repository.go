package billingandpayments

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type InvoiceRepository interface {
	AddInvoice(ctx context.Context, invoice *Invoice) error

	GetInvoice(ctx context.Context, id uuid.UUID) (*Invoice, error)

	UpdateInvoice(
		ctx context.Context,
		invoiceId uuid.UUID,
		updateFn func(e *Invoice) (*Invoice, error),
	) error
}

// MapToInvoice creates a Invoice domain object from the given parameters.
// This should ONLY BE USED when reconstructing an Invoice from its repository.
func MapToInvoice(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//invoiceId string,
	//
	//ownerId string,
	//
	//petId *string,
	//
	//appointmentId *string,
	//
	//lineItems string,
	//
	//subtotal float32,
	//
	//discountAmount *float32,
	//
	//taxAmount float32,
	//
	//totalAmount float32,
	//
	//status string,
	//
	//issuedDate string,
	//
	//dueDate string,
	//
) (*Invoice, error) {
	return &Invoice{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
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
		// TODO
	}, nil
}

type PaymentRepository interface {
	AddPayment(ctx context.Context, payment *Payment) error

	GetPayment(ctx context.Context, id uuid.UUID) (*Payment, error)

	UpdatePayment(
		ctx context.Context,
		paymentId uuid.UUID,
		updateFn func(e *Payment) (*Payment, error),
	) error
}

// MapToPayment creates a Payment domain object from the given parameters.
// This should ONLY BE USED when reconstructing an Payment from its repository.
func MapToPayment(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//paymentId string,
	//
	//invoiceId string,
	//
	//ownerId string,
	//
	//amountPaid float32,
	//
	//paymentMethod string,
	//
	//paymentDate string,
	//
	//status string,
	//
) (*Payment, error) {
	return &Payment{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
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
		// TODO
	}, nil
}

type RefundRepository interface {
	AddRefund(ctx context.Context, refund *Refund) error

	GetRefund(ctx context.Context, id uuid.UUID) (*Refund, error)

	UpdateRefund(
		ctx context.Context,
		refundId uuid.UUID,
		updateFn func(e *Refund) (*Refund, error),
	) error
}

// MapToRefund creates a Refund domain object from the given parameters.
// This should ONLY BE USED when reconstructing an Refund from its repository.
func MapToRefund(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//refundId string,
	//
	//paymentId string,
	//
	//ownerId string,
	//
	//refundAmount float32,
	//
	//reason string,
	//
	//refundDate string,
	//
	//status string,
	//
) (*Refund, error) {
	return &Refund{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
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
		// TODO
	}, nil
}
