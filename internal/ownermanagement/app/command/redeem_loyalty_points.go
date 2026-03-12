package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/ownermanagement/domain/ownermanagement"
)

type RedeemLoyaltyPoints struct {
	//
	//OwnerId string,
	//
	//Points int32,
	//
	// TODO
}

type RedeemLoyaltyPointsHandler struct {
	OwnerRepo ownermanagement.OwnerRepository

	LoyaltyAccountRepo ownermanagement.LoyaltyAccountRepository
}

func NewRedeemLoyaltyPointsHandler(
	ownerRepo ownermanagement.OwnerRepository,

	loyaltyaccountRepo ownermanagement.LoyaltyAccountRepository,
) RedeemLoyaltyPointsHandler {
	return RedeemLoyaltyPointsHandler{
		OwnerRepo: ownerRepo,

		LoyaltyAccountRepo: loyaltyaccountRepo,
	}
}

func (h RedeemLoyaltyPointsHandler) Handle(ctx context.Context, cmd RedeemLoyaltyPoints) error {
	// ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.command.redeem_loyalty_points.handle")
	// defer span.End()

	// TODO
	//err = h.OwnerRepo.UpdateOwner(ctx, uuid.MustParse(cmd.OwnerId), func(o *ownermanagement.Owner) (*ownermanagement.Owner, error) {
	//
	//	 err := o.RedeemLoyaltyPoints(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return o, nil
	//})
	//if err != nil {
	//	return err
	//}

	// TODO
	//err = h.LoyaltyAccountRepo.UpdateLoyaltyAccount(ctx, uuid.MustParse(cmd.LoyaltyAccountId), func(l *ownermanagement.LoyaltyAccount) (*ownermanagement.LoyaltyAccount, error) {
	//
	//	 err := l.RedeemLoyaltyPoints(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return l, nil
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}
