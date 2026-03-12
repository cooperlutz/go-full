package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/billingandpayments/domain/billingandpayments"
)

type ProcessPayment struct {
	//
	//InvoiceId string,
	//
	//OwnerId string,
	//
	//AmountPaid float32,
	//
	//PaymentMethod string,
	//
	// TODO
}

type ProcessPaymentHandler struct {
	InvoiceRepo billingandpayments.InvoiceRepository

	PaymentRepo billingandpayments.PaymentRepository

	RefundRepo billingandpayments.RefundRepository
}

func NewProcessPaymentHandler(
	invoiceRepo billingandpayments.InvoiceRepository,

	paymentRepo billingandpayments.PaymentRepository,

	refundRepo billingandpayments.RefundRepository,
) ProcessPaymentHandler {
	return ProcessPaymentHandler{
		InvoiceRepo: invoiceRepo,

		PaymentRepo: paymentRepo,

		RefundRepo: refundRepo,
	}
}

func (h ProcessPaymentHandler) Handle(ctx context.Context, cmd ProcessPayment) error {
	// ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.command.process_payment.handle")
	// defer span.End()

	// TODO
	//err = h.InvoiceRepo.UpdateInvoice(ctx, uuid.MustParse(cmd.InvoiceId), func(i *billingandpayments.Invoice) (*billingandpayments.Invoice, error) {
	//
	//	 err := i.ProcessPayment(
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
	//	 err := p.ProcessPayment(
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
	//	 err := r.ProcessPayment(
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
