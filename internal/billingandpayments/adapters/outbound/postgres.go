//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package outbound

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/cooperlutz/go-full/internal/billingandpayments/app/query"
	"github.com/cooperlutz/go-full/internal/billingandpayments/domain/billingandpayments"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// PostgresAdapter implements the examination repository using Postgres as the data store.
type PostgresAdapter struct {
	Handler IQuerierBillingAndPayments
}

// NewPostgresAdapter creates a new instance of PostgresAdapter.
func NewPostgresAdapter(db deebee.IDatabase) PostgresAdapter {
	return PostgresAdapter{
		Handler: NewQueriesWrapper(db),
	}
}

func (p PostgresAdapter) FindAllInvoices(ctx context.Context) ([]query.Invoice, error) {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.adapters.outbound.postgres.find_all_invoice")
	defer span.End()

	invoices, err := p.Handler.FindAllInvoices(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return billingandpaymentsInvoicesToQuery(invoices)
}

func (p PostgresAdapter) FindOneInvoice(ctx context.Context, id uuid.UUID) (query.Invoice, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_invoice")
	defer span.End()

	invoice, err := p.GetInvoice(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.Invoice{}, err
	}

	return mapEntityInvoiceToQuery(invoice), nil
}

// AddInvoice adds a new exam to the database.
func (p PostgresAdapter) AddInvoice(ctx context.Context, invoice *billingandpayments.Invoice) error {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.adapters.outbound.postgres.add_invoice")
	defer span.End()

	dbInvoice := mapEntityInvoiceToDB(invoice)

	err := p.Handler.AddInvoice(ctx, AddInvoiceParams(dbInvoice))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetInvoice(ctx context.Context, id uuid.UUID) (*billingandpayments.Invoice, error) {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.adapters.outbound.postgres.get_invoice")
	defer span.End()

	invoice, err := p.Handler.GetInvoice(
		ctx,
		GetInvoiceParams{InvoiceID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return invoice.toDomain()
}

func (p PostgresAdapter) UpdateInvoice(
	ctx context.Context,
	invoiceId uuid.UUID,
	updateFn func(e *billingandpayments.Invoice) (*billingandpayments.Invoice, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.adapters.outbound.postgres.update_invoice")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	invoice, err := p.GetInvoice(ctx, invoiceId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedInvoice, err := updateFn(invoice)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbInvoice := mapEntityInvoiceToDB(updatedInvoice)

	err = p.Handler.UpdateInvoice(ctx, UpdateInvoiceParams(dbInvoice))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) FindAllPayments(ctx context.Context) ([]query.Payment, error) {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.adapters.outbound.postgres.find_all_payment")
	defer span.End()

	payments, err := p.Handler.FindAllPayments(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return billingandpaymentsPaymentsToQuery(payments)
}

func (p PostgresAdapter) FindOnePayment(ctx context.Context, id uuid.UUID) (query.Payment, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_payment")
	defer span.End()

	payment, err := p.GetPayment(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.Payment{}, err
	}

	return mapEntityPaymentToQuery(payment), nil
}

// AddPayment adds a new exam to the database.
func (p PostgresAdapter) AddPayment(ctx context.Context, payment *billingandpayments.Payment) error {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.adapters.outbound.postgres.add_payment")
	defer span.End()

	dbPayment := mapEntityPaymentToDB(payment)

	err := p.Handler.AddPayment(ctx, AddPaymentParams(dbPayment))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetPayment(ctx context.Context, id uuid.UUID) (*billingandpayments.Payment, error) {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.adapters.outbound.postgres.get_payment")
	defer span.End()

	payment, err := p.Handler.GetPayment(
		ctx,
		GetPaymentParams{PaymentID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return payment.toDomain()
}

func (p PostgresAdapter) UpdatePayment(
	ctx context.Context,
	paymentId uuid.UUID,
	updateFn func(e *billingandpayments.Payment) (*billingandpayments.Payment, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.adapters.outbound.postgres.update_payment")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	payment, err := p.GetPayment(ctx, paymentId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedPayment, err := updateFn(payment)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbPayment := mapEntityPaymentToDB(updatedPayment)

	err = p.Handler.UpdatePayment(ctx, UpdatePaymentParams(dbPayment))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) FindAllRefunds(ctx context.Context) ([]query.Refund, error) {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.adapters.outbound.postgres.find_all_refund")
	defer span.End()

	refunds, err := p.Handler.FindAllRefunds(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return billingandpaymentsRefundsToQuery(refunds)
}

func (p PostgresAdapter) FindOneRefund(ctx context.Context, id uuid.UUID) (query.Refund, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_refund")
	defer span.End()

	refund, err := p.GetRefund(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.Refund{}, err
	}

	return mapEntityRefundToQuery(refund), nil
}

// AddRefund adds a new exam to the database.
func (p PostgresAdapter) AddRefund(ctx context.Context, refund *billingandpayments.Refund) error {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.adapters.outbound.postgres.add_refund")
	defer span.End()

	dbRefund := mapEntityRefundToDB(refund)

	err := p.Handler.AddRefund(ctx, AddRefundParams(dbRefund))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetRefund(ctx context.Context, id uuid.UUID) (*billingandpayments.Refund, error) {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.adapters.outbound.postgres.get_refund")
	defer span.End()

	refund, err := p.Handler.GetRefund(
		ctx,
		GetRefundParams{RefundID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return refund.toDomain()
}

func (p PostgresAdapter) UpdateRefund(
	ctx context.Context,
	refundId uuid.UUID,
	updateFn func(e *billingandpayments.Refund) (*billingandpayments.Refund, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.adapters.outbound.postgres.update_refund")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	refund, err := p.GetRefund(ctx, refundId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedRefund, err := updateFn(refund)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbRefund := mapEntityRefundToDB(updatedRefund)

	err = p.Handler.UpdateRefund(ctx, UpdateRefundParams(dbRefund))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

// finishTransaction commits or rolls back the transaction based on the error state.
func (p PostgresAdapter) finishTransaction(ctx context.Context, err error, tx pgx.Tx) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			telemetree.RecordError(ctx, rollbackErr, "failed to rollback tx")

			return rollbackErr
		}

		return err
	} else {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			telemetree.RecordError(ctx, commitErr, "failed to commit tx")

			return commitErr
		}

		return nil
	}
}
