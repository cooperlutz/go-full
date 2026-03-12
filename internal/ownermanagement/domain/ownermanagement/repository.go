package ownermanagement

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type OwnerRepository interface {
	AddOwner(ctx context.Context, owner *Owner) error

	GetOwner(ctx context.Context, id uuid.UUID) (*Owner, error)

	UpdateOwner(
		ctx context.Context,
		ownerId uuid.UUID,
		updateFn func(e *Owner) (*Owner, error),
	) error
}

// MapToOwner creates a Owner domain object from the given parameters.
// This should ONLY BE USED when reconstructing an Owner from its repository.
func MapToOwner(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//ownerId string,
	//
	//firstName string,
	//
	//lastName string,
	//
	//email string,
	//
	//phoneNumber string,
	//
	//address *string,
	//
	//communicationPreference string,
	//
	//loyaltyMember bool,
	//
	//loyaltyPoints *int32,
	//
	//status string,
	//
) (*Owner, error) {
	return &Owner{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		//
		//ownerId: ownerId,
		//
		//firstName: firstName,
		//
		//lastName: lastName,
		//
		//email: email,
		//
		//phoneNumber: phoneNumber,
		//
		//address: address,
		//
		//communicationPreference: communicationPreference,
		//
		//loyaltyMember: loyaltyMember,
		//
		//loyaltyPoints: loyaltyPoints,
		//
		//status: status,
		//
		// TODO
	}, nil
}

type LoyaltyAccountRepository interface {
	AddLoyaltyAccount(ctx context.Context, loyaltyaccount *LoyaltyAccount) error

	GetLoyaltyAccount(ctx context.Context, id uuid.UUID) (*LoyaltyAccount, error)

	UpdateLoyaltyAccount(
		ctx context.Context,
		loyaltyaccountId uuid.UUID,
		updateFn func(e *LoyaltyAccount) (*LoyaltyAccount, error),
	) error
}

// MapToLoyaltyAccount creates a LoyaltyAccount domain object from the given parameters.
// This should ONLY BE USED when reconstructing an LoyaltyAccount from its repository.
func MapToLoyaltyAccount(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//loyaltyAccountId string,
	//
	//ownerId string,
	//
	//pointsBalance int32,
	//
	//tier string,
	//
	//enrolledDate string,
	//
) (*LoyaltyAccount, error) {
	return &LoyaltyAccount{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		//
		//loyaltyAccountId: loyaltyAccountId,
		//
		//ownerId: ownerId,
		//
		//pointsBalance: pointsBalance,
		//
		//tier: tier,
		//
		//enrolledDate: enrolledDate,
		//
		// TODO
	}, nil
}
