package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/billingandpayments/domain/billingandpayments"
)

type GenerateInvoice struct {
	//
	//OwnerId string,
	//
	//PetId *string,
	//
	//AppointmentId *string,
	//
	//LineItems string,
	//
	//DiscountAmount *float32,
	//
	// TODO
}

type GenerateInvoiceHandler struct {
	InvoiceRepo billingandpayments.InvoiceRepository

	PaymentRepo billingandpayments.PaymentRepository

	RefundRepo billingandpayments.RefundRepository
}

func NewGenerateInvoiceHandler(
	invoiceRepo billingandpayments.InvoiceRepository,

	paymentRepo billingandpayments.PaymentRepository,

	refundRepo billingandpayments.RefundRepository,
) GenerateInvoiceHandler {
	return GenerateInvoiceHandler{
		InvoiceRepo: invoiceRepo,

		PaymentRepo: paymentRepo,

		RefundRepo: refundRepo,
	}
}

func (h GenerateInvoiceHandler) Handle(ctx context.Context, cmd GenerateInvoice) error {
	// ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.command.generate_invoice.handle")
	// defer span.End()

	// TODO
	//err = h.InvoiceRepo.UpdateInvoice(ctx, uuid.MustParse(cmd.InvoiceId), func(i *billingandpayments.Invoice) (*billingandpayments.Invoice, error) {
	//
	//	 err := i.GenerateInvoice(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return i, nil
	//})
	//if err != nil {
	//	return err
	//}

	// TODO
	//err = h.PaymentRepo.UpdatePayment(ctx, uuid.MustParse(cmd.PaymentId), func(p *billingandpayments.Payment) (*billingandpayments.Payment, error) {
	//
	//	 err := p.GenerateInvoice(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return p, nil
	//})
	//if err != nil {
	//	return err
	//}

	// TODO
	//err = h.RefundRepo.UpdateRefund(ctx, uuid.MustParse(cmd.RefundId), func(r *billingandpayments.Refund) (*billingandpayments.Refund, error) {
	//
	//	 err := r.GenerateInvoice(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return r, nil
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}
