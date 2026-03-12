package inbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/billingandpayments/app"
	"github.com/cooperlutz/go-full/internal/billingandpayments/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// HttpAdapter represents the HTTP server for the BillingAndPayments module.
type HttpAdapter struct {
	app app.Application
}

// NewHttpAdapter creates a new HttpAdapter instance with the provided BillingAndPayments application.
func NewHttpAdapter(application app.Application) HttpAdapter {
	return HttpAdapter{
		app: application,
	}
}

// StrictHandler returns a strict HTTP handler for the BillingAndPayments module.
func (h HttpAdapter) StrictHandler() ServerInterface {
	return NewStrictHandler(h, nil)
}

// (GET /v1/invoices).
func (h HttpAdapter) FindAllInvoices(ctx context.Context, request FindAllInvoicesRequestObject) (FindAllInvoicesResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "invoice.adapters.inbound.http.find_all_invoices")
	defer span.End()

	invoice, err := h.app.Queries.FindAllInvoices.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseInvoices []Invoice
	for _, e := range invoice {
		responseInvoices = append(responseInvoices, queryInvoiceToHttpInvoice(e))
	}

	return FindAllInvoices200JSONResponse(responseInvoices), nil
}

// (GET /v1/invoice/{invoiceId}).
func (h HttpAdapter) FindOneInvoice(ctx context.Context, request FindOneInvoiceRequestObject) (FindOneInvoiceResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_invoice")
	defer span.End()

	invoice, err := h.app.Queries.FindOneInvoice.Handle(ctx, query.FindOneInvoice{InvoiceID: request.InvoiceId})
	if err != nil {
		return nil, err
	}

	return FindOneInvoice200JSONResponse(queryInvoiceToHttpInvoice(invoice)), nil
}

func queryInvoiceToHttpInvoice(e query.Invoice) Invoice {
	return Invoice{
		//
		//InvoiceId: GetInvoiceId(),
		//
		//OwnerId: GetOwnerId(),
		//
		//PetId: GetPetId(),
		//
		//AppointmentId: GetAppointmentId(),
		//
		//LineItems: GetLineItems(),
		//
		//Subtotal: GetSubtotal(),
		//
		//DiscountAmount: GetDiscountAmount(),
		//
		//TaxAmount: GetTaxAmount(),
		//
		//TotalAmount: GetTotalAmount(),
		//
		//Status: GetStatus(),
		//
		//IssuedDate: GetIssuedDate(),
		//
		//DueDate: GetDueDate(),
		//
		// TODO
	}
}

// (GET /v1/payments).
func (h HttpAdapter) FindAllPayments(ctx context.Context, request FindAllPaymentsRequestObject) (FindAllPaymentsResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "payment.adapters.inbound.http.find_all_payments")
	defer span.End()

	payment, err := h.app.Queries.FindAllPayments.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responsePayments []Payment
	for _, e := range payment {
		responsePayments = append(responsePayments, queryPaymentToHttpPayment(e))
	}

	return FindAllPayments200JSONResponse(responsePayments), nil
}

// (GET /v1/payment/{paymentId}).
func (h HttpAdapter) FindOnePayment(ctx context.Context, request FindOnePaymentRequestObject) (FindOnePaymentResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_payment")
	defer span.End()

	payment, err := h.app.Queries.FindOnePayment.Handle(ctx, query.FindOnePayment{PaymentID: request.PaymentId})
	if err != nil {
		return nil, err
	}

	return FindOnePayment200JSONResponse(queryPaymentToHttpPayment(payment)), nil
}

func queryPaymentToHttpPayment(e query.Payment) Payment {
	return Payment{
		//
		//PaymentId: GetPaymentId(),
		//
		//InvoiceId: GetInvoiceId(),
		//
		//OwnerId: GetOwnerId(),
		//
		//AmountPaid: GetAmountPaid(),
		//
		//PaymentMethod: GetPaymentMethod(),
		//
		//PaymentDate: GetPaymentDate(),
		//
		//Status: GetStatus(),
		//
		// TODO
	}
}

// (GET /v1/refunds).
func (h HttpAdapter) FindAllRefunds(ctx context.Context, request FindAllRefundsRequestObject) (FindAllRefundsResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "refund.adapters.inbound.http.find_all_refunds")
	defer span.End()

	refund, err := h.app.Queries.FindAllRefunds.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseRefunds []Refund
	for _, e := range refund {
		responseRefunds = append(responseRefunds, queryRefundToHttpRefund(e))
	}

	return FindAllRefunds200JSONResponse(responseRefunds), nil
}

// (GET /v1/refund/{refundId}).
func (h HttpAdapter) FindOneRefund(ctx context.Context, request FindOneRefundRequestObject) (FindOneRefundResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_refund")
	defer span.End()

	refund, err := h.app.Queries.FindOneRefund.Handle(ctx, query.FindOneRefund{RefundID: request.RefundId})
	if err != nil {
		return nil, err
	}

	return FindOneRefund200JSONResponse(queryRefundToHttpRefund(refund)), nil
}

func queryRefundToHttpRefund(e query.Refund) Refund {
	return Refund{
		//
		//RefundId: GetRefundId(),
		//
		//PaymentId: GetPaymentId(),
		//
		//OwnerId: GetOwnerId(),
		//
		//RefundAmount: GetRefundAmount(),
		//
		//Reason: GetReason(),
		//
		//RefundDate: GetRefundDate(),
		//
		//Status: GetStatus(),
		//
		// TODO
	}
}
