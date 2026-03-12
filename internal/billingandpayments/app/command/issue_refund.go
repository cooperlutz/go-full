package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/billingandpayments/domain/billingandpayments"
)

type IssueRefund struct {
	//
	//PaymentId string,
	//
	//RefundAmount float32,
	//
	//Reason string,
	//
	// TODO
}

type IssueRefundHandler struct {
	InvoiceRepo billingandpayments.InvoiceRepository

	PaymentRepo billingandpayments.PaymentRepository

	RefundRepo billingandpayments.RefundRepository
}

func NewIssueRefundHandler(
	invoiceRepo billingandpayments.InvoiceRepository,

	paymentRepo billingandpayments.PaymentRepository,

	refundRepo billingandpayments.RefundRepository,
) IssueRefundHandler {
	return IssueRefundHandler{
		InvoiceRepo: invoiceRepo,

		PaymentRepo: paymentRepo,

		RefundRepo: refundRepo,
	}
}

func (h IssueRefundHandler) Handle(ctx context.Context, cmd IssueRefund) error {
	// ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.command.issue_refund.handle")
	// defer span.End()

	// TODO
	//err = h.InvoiceRepo.UpdateInvoice(ctx, uuid.MustParse(cmd.InvoiceId), func(i *billingandpayments.Invoice) (*billingandpayments.Invoice, error) {
	//
	//	 err := i.IssueRefund(
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
	//	 err := p.IssueRefund(
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
	//	 err := r.IssueRefund(
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
