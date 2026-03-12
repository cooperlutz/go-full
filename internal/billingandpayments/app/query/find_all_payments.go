//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllPaymentsReadModel interface {
	FindAllPayments(ctx context.Context) ([]Payment, error)
}

type FindAllPaymentsHandler struct {
	readModel FindAllPaymentsReadModel
}

func NewFindAllPaymentsHandler(
	readModel FindAllPaymentsReadModel,
) FindAllPaymentsHandler {
	return FindAllPaymentsHandler{readModel: readModel}
}

func (h FindAllPaymentsHandler) Handle(ctx context.Context) ([]Payment, error) {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.query.find_all_payments.handle")
	defer span.End()

	exams, err := h.readModel.FindAllPayments(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
