//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllRefundsReadModel interface {
	FindAllRefunds(ctx context.Context) ([]Refund, error)
}

type FindAllRefundsHandler struct {
	readModel FindAllRefundsReadModel
}

func NewFindAllRefundsHandler(
	readModel FindAllRefundsReadModel,
) FindAllRefundsHandler {
	return FindAllRefundsHandler{readModel: readModel}
}

func (h FindAllRefundsHandler) Handle(ctx context.Context) ([]Refund, error) {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.query.find_all_refunds.handle")
	defer span.End()

	exams, err := h.readModel.FindAllRefunds(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
