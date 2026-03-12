//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneLoyaltyAccount struct {
	LoyaltyAccountID string
}

type FindOneLoyaltyAccountReadModel interface {
	FindOneLoyaltyAccount(ctx context.Context, loyaltyaccountId uuid.UUID) (LoyaltyAccount, error)
}

type FindOneLoyaltyAccountHandler struct {
	readModel FindOneLoyaltyAccountReadModel
}

func NewFindOneLoyaltyAccountHandler(
	readModel FindOneLoyaltyAccountReadModel,
) FindOneLoyaltyAccountHandler {
	return FindOneLoyaltyAccountHandler{readModel: readModel}
}

func (h FindOneLoyaltyAccountHandler) Handle(ctx context.Context, qry FindOneLoyaltyAccount) (LoyaltyAccount, error) {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.query.find_one_loyalty_account.handle")
	defer span.End()

	loyaltyaccount, err := h.readModel.FindOneLoyaltyAccount(ctx, uuid.MustParse(qry.LoyaltyAccountID))
	if err != nil {
		return LoyaltyAccount{}, err
	}

	return loyaltyaccount, nil
}
