//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneInvoice struct {
	InvoiceID string
}

type FindOneInvoiceReadModel interface {
	FindOneInvoice(ctx context.Context, invoiceId uuid.UUID) (Invoice, error)
}

type FindOneInvoiceHandler struct {
	readModel FindOneInvoiceReadModel
}

func NewFindOneInvoiceHandler(
	readModel FindOneInvoiceReadModel,
) FindOneInvoiceHandler {
	return FindOneInvoiceHandler{readModel: readModel}
}

func (h FindOneInvoiceHandler) Handle(ctx context.Context, qry FindOneInvoice) (Invoice, error) {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.query.find_one_invoice.handle")
	defer span.End()

	invoice, err := h.readModel.FindOneInvoice(ctx, uuid.MustParse(qry.InvoiceID))
	if err != nil {
		return Invoice{}, err
	}

	return invoice, nil
}
