//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllInvoicesReadModel interface {
	FindAllInvoices(ctx context.Context) ([]Invoice, error)
}

type FindAllInvoicesHandler struct {
	readModel FindAllInvoicesReadModel
}

func NewFindAllInvoicesHandler(
	readModel FindAllInvoicesReadModel,
) FindAllInvoicesHandler {
	return FindAllInvoicesHandler{readModel: readModel}
}

func (h FindAllInvoicesHandler) Handle(ctx context.Context) ([]Invoice, error) {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.query.find_all_invoices.handle")
	defer span.End()

	exams, err := h.readModel.FindAllInvoices(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
