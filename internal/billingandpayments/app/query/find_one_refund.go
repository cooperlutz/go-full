//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneRefund struct {
	RefundID string
}

type FindOneRefundReadModel interface {
	FindOneRefund(ctx context.Context, refundId uuid.UUID) (Refund, error)
}

type FindOneRefundHandler struct {
	readModel FindOneRefundReadModel
}

func NewFindOneRefundHandler(
	readModel FindOneRefundReadModel,
) FindOneRefundHandler {
	return FindOneRefundHandler{readModel: readModel}
}

func (h FindOneRefundHandler) Handle(ctx context.Context, qry FindOneRefund) (Refund, error) {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.query.find_one_refund.handle")
	defer span.End()

	refund, err := h.readModel.FindOneRefund(ctx, uuid.MustParse(qry.RefundID))
	if err != nil {
		return Refund{}, err
	}

	return refund, nil
}
