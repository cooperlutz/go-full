package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/ownermanagement/domain/ownermanagement"
)

type RegisterOwner struct {
	//
	//FirstName string,
	//
	//LastName string,
	//
	//Email string,
	//
	//PhoneNumber string,
	//
	//Address *string,
	//
	//CommunicationPreference string,
	//
	// TODO
}

type RegisterOwnerHandler struct {
	OwnerRepo ownermanagement.OwnerRepository

	LoyaltyAccountRepo ownermanagement.LoyaltyAccountRepository
}

func NewRegisterOwnerHandler(
	ownerRepo ownermanagement.OwnerRepository,

	loyaltyaccountRepo ownermanagement.LoyaltyAccountRepository,
) RegisterOwnerHandler {
	return RegisterOwnerHandler{
		OwnerRepo: ownerRepo,

		LoyaltyAccountRepo: loyaltyaccountRepo,
	}
}

func (h RegisterOwnerHandler) Handle(ctx context.Context, cmd RegisterOwner) error {
	// ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.command.register_owner.handle")
	// defer span.End()

	// TODO
	//err = h.OwnerRepo.UpdateOwner(ctx, uuid.MustParse(cmd.OwnerId), func(o *ownermanagement.Owner) (*ownermanagement.Owner, error) {
	//
	//	 err := o.RegisterOwner(
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
	//	 err := l.RegisterOwner(
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
