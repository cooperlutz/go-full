package ownermanagement

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type Owner struct {
	*baseentitee.EntityMetadata
	//
	//ownerId string
	//
	//firstName string
	//
	//lastName string
	//
	//email string
	//
	//phoneNumber string
	//
	//address *string
	//
	//communicationPreference string
	//
	//loyaltyMember bool
	//
	//loyaltyPoints *int32
	//
	//status string
	//
	// TODO
}

func NewOwner(
// ownerId string,
//
// firstName string,
//
// lastName string,
//
// email string,
//
// phoneNumber string,
//
// address *string,
//
// communicationPreference string,
//
// loyaltyMember bool,
//
// loyaltyPoints *int32,
//
// status string,
) *Owner {
	return &Owner{
		EntityMetadata: baseentitee.NewEntityMetadata(),
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
	}
}
