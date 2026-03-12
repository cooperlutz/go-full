package insuranceandclaims

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type InsuranceProvider struct {
	*baseentitee.EntityMetadata
	//
	//providerId string
	//
	//name string
	//
	//contactName *string
	//
	//email string
	//
	//phoneNumber string
	//
	//claimSubmissionUrl *string
	//
	//status string
	//
	// TODO
}

func NewInsuranceProvider(
// providerId string,
//
// name string,
//
// contactName *string,
//
// email string,
//
// phoneNumber string,
//
// claimSubmissionUrl *string,
//
// status string,
) *InsuranceProvider {
	return &InsuranceProvider{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		//
		//providerId: providerId,
		//
		//name: name,
		//
		//contactName: contactName,
		//
		//email: email,
		//
		//phoneNumber: phoneNumber,
		//
		//claimSubmissionUrl: claimSubmissionUrl,
		//
		//status: status,
		//
	}
}
