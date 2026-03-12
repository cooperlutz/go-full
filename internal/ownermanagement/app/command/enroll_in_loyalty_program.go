package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/ownermanagement/domain/ownermanagement"
)

type EnrollInLoyaltyProgram struct {
	//
	//OwnerId string,
	//
	// TODO
}

type EnrollInLoyaltyProgramHandler struct {
	OwnerRepo ownermanagement.OwnerRepository

	LoyaltyAccountRepo ownermanagement.LoyaltyAccountRepository
}

func NewEnrollInLoyaltyProgramHandler(
	ownerRepo ownermanagement.OwnerRepository,

	loyaltyaccountRepo ownermanagement.LoyaltyAccountRepository,
) EnrollInLoyaltyProgramHandler {
	return EnrollInLoyaltyProgramHandler{
		OwnerRepo: ownerRepo,

		LoyaltyAccountRepo: loyaltyaccountRepo,
	}
}

func (h EnrollInLoyaltyProgramHandler) Handle(ctx context.Context, cmd EnrollInLoyaltyProgram) error {
	// ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.command.enroll_in_loyalty_program.handle")
	// defer span.End()

	// TODO
	//err = h.OwnerRepo.UpdateOwner(ctx, uuid.MustParse(cmd.OwnerId), func(o *ownermanagement.Owner) (*ownermanagement.Owner, error) {
	//
	//	 err := o.EnrollInLoyaltyProgram(
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
	//	 err := l.EnrollInLoyaltyProgram(
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
