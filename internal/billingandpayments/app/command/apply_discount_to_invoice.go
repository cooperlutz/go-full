package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/billingandpayments/domain/billingandpayments"
)

type ApplyDiscountToInvoice struct {
	//
	//InvoiceId string,
	//
	//DiscountAmount float32,
	//
	//DiscountReason string,
	//
	// TODO
}

type ApplyDiscountToInvoiceHandler struct {
	InvoiceRepo billingandpayments.InvoiceRepository

	PaymentRepo billingandpayments.PaymentRepository

	RefundRepo billingandpayments.RefundRepository
}

func NewApplyDiscountToInvoiceHandler(
	invoiceRepo billingandpayments.InvoiceRepository,

	paymentRepo billingandpayments.PaymentRepository,

	refundRepo billingandpayments.RefundRepository,
) ApplyDiscountToInvoiceHandler {
	return ApplyDiscountToInvoiceHandler{
		InvoiceRepo: invoiceRepo,

		PaymentRepo: paymentRepo,

		RefundRepo: refundRepo,
	}
}

func (h ApplyDiscountToInvoiceHandler) Handle(ctx context.Context, cmd ApplyDiscountToInvoice) error {
	// ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.command.apply_discount_to_invoice.handle")
	// defer span.End()

	// TODO
	//err = h.InvoiceRepo.UpdateInvoice(ctx, uuid.MustParse(cmd.InvoiceId), func(i *billingandpayments.Invoice) (*billingandpayments.Invoice, error) {
	//
	//	 err := i.ApplyDiscountToInvoice(
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
	//	 err := p.ApplyDiscountToInvoice(
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
	//	 err := r.ApplyDiscountToInvoice(
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
