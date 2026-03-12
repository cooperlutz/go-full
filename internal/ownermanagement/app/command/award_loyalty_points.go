package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/ownermanagement/domain/ownermanagement"
)

type AwardLoyaltyPoints struct {
	//
	//OwnerId string,
	//
	//Points int32,
	//
	//Reason string,
	//
	// TODO
}

type AwardLoyaltyPointsHandler struct {
	OwnerRepo ownermanagement.OwnerRepository

	LoyaltyAccountRepo ownermanagement.LoyaltyAccountRepository
}

func NewAwardLoyaltyPointsHandler(
	ownerRepo ownermanagement.OwnerRepository,

	loyaltyaccountRepo ownermanagement.LoyaltyAccountRepository,
) AwardLoyaltyPointsHandler {
	return AwardLoyaltyPointsHandler{
		OwnerRepo: ownerRepo,

		LoyaltyAccountRepo: loyaltyaccountRepo,
	}
}

func (h AwardLoyaltyPointsHandler) Handle(ctx context.Context, cmd AwardLoyaltyPoints) error {
	// ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.command.award_loyalty_points.handle")
	// defer span.End()

	// TODO
	//err = h.OwnerRepo.UpdateOwner(ctx, uuid.MustParse(cmd.OwnerId), func(o *ownermanagement.Owner) (*ownermanagement.Owner, error) {
	//
	//	 err := o.AwardLoyaltyPoints(
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
	//	 err := l.AwardLoyaltyPoints(
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
