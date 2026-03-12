//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOnePayment struct {
	PaymentID string
}

type FindOnePaymentReadModel interface {
	FindOnePayment(ctx context.Context, paymentId uuid.UUID) (Payment, error)
}

type FindOnePaymentHandler struct {
	readModel FindOnePaymentReadModel
}

func NewFindOnePaymentHandler(
	readModel FindOnePaymentReadModel,
) FindOnePaymentHandler {
	return FindOnePaymentHandler{readModel: readModel}
}

func (h FindOnePaymentHandler) Handle(ctx context.Context, qry FindOnePayment) (Payment, error) {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.query.find_one_payment.handle")
	defer span.End()

	payment, err := h.readModel.FindOnePayment(ctx, uuid.MustParse(qry.PaymentID))
	if err != nil {
		return Payment{}, err
	}

	return payment, nil
}
