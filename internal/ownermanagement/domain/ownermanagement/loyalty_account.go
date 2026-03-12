package ownermanagement

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type LoyaltyAccount struct {
	*baseentitee.EntityMetadata
	//
	//loyaltyAccountId string
	//
	//ownerId string
	//
	//pointsBalance int32
	//
	//tier string
	//
	//enrolledDate string
	//
	// TODO
}

func NewLoyaltyAccount(
// loyaltyAccountId string,
//
// ownerId string,
//
// pointsBalance int32,
//
// tier string,
//
// enrolledDate string,
) *LoyaltyAccount {
	return &LoyaltyAccount{
		EntityMetadata: baseentitee.NewEntityMetadata(),
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
	}
}
