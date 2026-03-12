//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllLoyaltyAccountsReadModel interface {
	FindAllLoyaltyAccounts(ctx context.Context) ([]LoyaltyAccount, error)
}

type FindAllLoyaltyAccountsHandler struct {
	readModel FindAllLoyaltyAccountsReadModel
}

func NewFindAllLoyaltyAccountsHandler(
	readModel FindAllLoyaltyAccountsReadModel,
) FindAllLoyaltyAccountsHandler {
	return FindAllLoyaltyAccountsHandler{readModel: readModel}
}

func (h FindAllLoyaltyAccountsHandler) Handle(ctx context.Context) ([]LoyaltyAccount, error) {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.query.find_all_loyalty_accounts.handle")
	defer span.End()

	exams, err := h.readModel.FindAllLoyaltyAccounts(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
