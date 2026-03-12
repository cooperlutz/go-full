package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/billingandpayments/domain/billingandpayments"
)

type VoidInvoice struct {
	//
	//InvoiceId string,
	//
	//Reason string,
	//
	// TODO
}

type VoidInvoiceHandler struct {
	InvoiceRepo billingandpayments.InvoiceRepository

	PaymentRepo billingandpayments.PaymentRepository

	RefundRepo billingandpayments.RefundRepository
}

func NewVoidInvoiceHandler(
	invoiceRepo billingandpayments.InvoiceRepository,

	paymentRepo billingandpayments.PaymentRepository,

	refundRepo billingandpayments.RefundRepository,
) VoidInvoiceHandler {
	return VoidInvoiceHandler{
		InvoiceRepo: invoiceRepo,

		PaymentRepo: paymentRepo,

		RefundRepo: refundRepo,
	}
}

func (h VoidInvoiceHandler) Handle(ctx context.Context, cmd VoidInvoice) error {
	// ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.command.void_invoice.handle")
	// defer span.End()

	// TODO
	//err = h.InvoiceRepo.UpdateInvoice(ctx, uuid.MustParse(cmd.InvoiceId), func(i *billingandpayments.Invoice) (*billingandpayments.Invoice, error) {
	//
	//	 err := i.VoidInvoice(
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
	//	 err := p.VoidInvoice(
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
	//	 err := r.VoidInvoice(
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
