package outbound

import (
	"github.com/cooperlutz/go-full/internal/billingandpayments/app/query"
	"github.com/cooperlutz/go-full/internal/billingandpayments/domain/billingandpayments"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

// toDomain maps the InvoiceInvoice to the domain entity.
func (e BillingandpaymentsInvoice) toDomain() (*billingandpayments.Invoice, error) {
	return billingandpayments.MapToInvoice(
		e.InvoiceID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.InvoiceId,
		//
		//e.OwnerId,
		//
		//e.PetId,
		//
		//e.AppointmentId,
		//
		//e.LineItems,
		//
		//e.Subtotal,
		//
		//e.DiscountAmount,
		//
		//e.TaxAmount,
		//
		//e.TotalAmount,
		//
		//e.Status,
		//
		//e.IssuedDate,
		//
		//e.DueDate,
		//
		// TODO
	)
}

// toQueryInvoice maps the invoiceInvoice to the query.Invoice.
func (e BillingandpaymentsInvoice) toQueryInvoice() (query.Invoice, error) {
	invoice, err := e.toDomain()
	if err != nil {
		return query.Invoice{}, err
	}

	return mapEntityInvoiceToQuery(invoice), nil
}

// invoiceInvoicesToQuery maps a slice of InvoiceInvoice to a slice of query.Invoice entities.
func billingandpaymentsInvoicesToQuery(invoices []BillingandpaymentsInvoice) ([]query.Invoice, error) {
	var domainInvoices []query.Invoice

	for _, invoice := range invoices {
		queryInvoice, err := invoice.toQueryInvoice()
		if err != nil {
			return nil, err
		}

		domainInvoices = append(domainInvoices, queryInvoice)
	}

	return domainInvoices, nil
}

// mapEntityInvoiceToDB maps a domain Invoice entity to the InvoiceInvoice database model.
func mapEntityInvoiceToDB(invoice *billingandpayments.Invoice) BillingandpaymentsInvoice {
	createdAt := invoice.GetCreatedAtTime()
	updatedAt := invoice.GetUpdatedAtTime()

	return BillingandpaymentsInvoice{
		InvoiceID: pgxutil.UUIDToPgtypeUUID(invoice.GetIdUUID()),
		CreatedAt: pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt: pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:   invoice.IsDeleted(),
		DeletedAt: pgxutil.TimeToTimestampz(invoice.GetDeletedAtTime()),
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

// mapEntityInvoiceToQuery maps a domain Invoice entity to a query.Invoice.
func mapEntityInvoiceToQuery(invoice *billingandpayments.Invoice) query.Invoice {
	return query.Invoice{
		// TODO
	}
}

// toDomain maps the PaymentPayment to the domain entity.
func (e BillingandpaymentsPayment) toDomain() (*billingandpayments.Payment, error) {
	return billingandpayments.MapToPayment(
		e.PaymentID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.PaymentId,
		//
		//e.InvoiceId,
		//
		//e.OwnerId,
		//
		//e.AmountPaid,
		//
		//e.PaymentMethod,
		//
		//e.PaymentDate,
		//
		//e.Status,
		//
		// TODO
	)
}

// toQueryPayment maps the paymentPayment to the query.Payment.
func (e BillingandpaymentsPayment) toQueryPayment() (query.Payment, error) {
	payment, err := e.toDomain()
	if err != nil {
		return query.Payment{}, err
	}

	return mapEntityPaymentToQuery(payment), nil
}

// paymentPaymentsToQuery maps a slice of PaymentPayment to a slice of query.Payment entities.
func billingandpaymentsPaymentsToQuery(payments []BillingandpaymentsPayment) ([]query.Payment, error) {
	var domainPayments []query.Payment

	for _, payment := range payments {
		queryPayment, err := payment.toQueryPayment()
		if err != nil {
			return nil, err
		}

		domainPayments = append(domainPayments, queryPayment)
	}

	return domainPayments, nil
}

// mapEntityPaymentToDB maps a domain Payment entity to the PaymentPayment database model.
func mapEntityPaymentToDB(payment *billingandpayments.Payment) BillingandpaymentsPayment {
	createdAt := payment.GetCreatedAtTime()
	updatedAt := payment.GetUpdatedAtTime()

	return BillingandpaymentsPayment{
		PaymentID: pgxutil.UUIDToPgtypeUUID(payment.GetIdUUID()),
		CreatedAt: pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt: pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:   payment.IsDeleted(),
		DeletedAt: pgxutil.TimeToTimestampz(payment.GetDeletedAtTime()),
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

// mapEntityPaymentToQuery maps a domain Payment entity to a query.Payment.
func mapEntityPaymentToQuery(payment *billingandpayments.Payment) query.Payment {
	return query.Payment{
		// TODO
	}
}

// toDomain maps the RefundRefund to the domain entity.
func (e BillingandpaymentsRefund) toDomain() (*billingandpayments.Refund, error) {
	return billingandpayments.MapToRefund(
		e.RefundID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.RefundId,
		//
		//e.PaymentId,
		//
		//e.OwnerId,
		//
		//e.RefundAmount,
		//
		//e.Reason,
		//
		//e.RefundDate,
		//
		//e.Status,
		//
		// TODO
	)
}

// toQueryRefund maps the refundRefund to the query.Refund.
func (e BillingandpaymentsRefund) toQueryRefund() (query.Refund, error) {
	refund, err := e.toDomain()
	if err != nil {
		return query.Refund{}, err
	}

	return mapEntityRefundToQuery(refund), nil
}

// refundRefundsToQuery maps a slice of RefundRefund to a slice of query.Refund entities.
func billingandpaymentsRefundsToQuery(refunds []BillingandpaymentsRefund) ([]query.Refund, error) {
	var domainRefunds []query.Refund

	for _, refund := range refunds {
		queryRefund, err := refund.toQueryRefund()
		if err != nil {
			return nil, err
		}

		domainRefunds = append(domainRefunds, queryRefund)
	}

	return domainRefunds, nil
}

// mapEntityRefundToDB maps a domain Refund entity to the RefundRefund database model.
func mapEntityRefundToDB(refund *billingandpayments.Refund) BillingandpaymentsRefund {
	createdAt := refund.GetCreatedAtTime()
	updatedAt := refund.GetUpdatedAtTime()

	return BillingandpaymentsRefund{
		RefundID:  pgxutil.UUIDToPgtypeUUID(refund.GetIdUUID()),
		CreatedAt: pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt: pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:   refund.IsDeleted(),
		DeletedAt: pgxutil.TimeToTimestampz(refund.GetDeletedAtTime()),
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

// mapEntityRefundToQuery maps a domain Refund entity to a query.Refund.
func mapEntityRefundToQuery(refund *billingandpayments.Refund) query.Refund {
	return query.Refund{
		// TODO
	}
}
