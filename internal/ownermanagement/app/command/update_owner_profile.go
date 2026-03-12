package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/ownermanagement/domain/ownermanagement"
)

type UpdateOwnerProfile struct {
	//
	//OwnerId string,
	//
	//Email *string,
	//
	//PhoneNumber *string,
	//
	//Address *string,
	//
	//CommunicationPreference *string,
	//
	// TODO
}

type UpdateOwnerProfileHandler struct {
	OwnerRepo ownermanagement.OwnerRepository

	LoyaltyAccountRepo ownermanagement.LoyaltyAccountRepository
}

func NewUpdateOwnerProfileHandler(
	ownerRepo ownermanagement.OwnerRepository,

	loyaltyaccountRepo ownermanagement.LoyaltyAccountRepository,
) UpdateOwnerProfileHandler {
	return UpdateOwnerProfileHandler{
		OwnerRepo: ownerRepo,

		LoyaltyAccountRepo: loyaltyaccountRepo,
	}
}

func (h UpdateOwnerProfileHandler) Handle(ctx context.Context, cmd UpdateOwnerProfile) error {
	// ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.command.update_owner_profile.handle")
	// defer span.End()

	// TODO
	//err = h.OwnerRepo.UpdateOwner(ctx, uuid.MustParse(cmd.OwnerId), func(o *ownermanagement.Owner) (*ownermanagement.Owner, error) {
	//
	//	 err := o.UpdateOwnerProfile(
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
	//	 err := l.UpdateOwnerProfile(
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
