package insuranceandclaims

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type InsuranceClaim struct {
	*baseentitee.EntityMetadata
	//
	//claimId string
	//
	//ownerId string
	//
	//petId string
	//
	//providerId string
	//
	//invoiceId string
	//
	//policyNumber string
	//
	//claimAmount float32
	//
	//approvedAmount *float32
	//
	//submissionDate string
	//
	//resolutionDate *string
	//
	//status string
	//
	//notes *string
	//
	// TODO
}

func NewInsuranceClaim(
// claimId string,
//
// ownerId string,
//
// petId string,
//
// providerId string,
//
// invoiceId string,
//
// policyNumber string,
//
// claimAmount float32,
//
// approvedAmount *float32,
//
// submissionDate string,
//
// resolutionDate *string,
//
// status string,
//
// notes *string,
) *InsuranceClaim {
	return &InsuranceClaim{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		//
		//claimId: claimId,
		//
		//ownerId: ownerId,
		//
		//petId: petId,
		//
		//providerId: providerId,
		//
		//invoiceId: invoiceId,
		//
		//policyNumber: policyNumber,
		//
		//claimAmount: claimAmount,
		//
		//approvedAmount: approvedAmount,
		//
		//submissionDate: submissionDate,
		//
		//resolutionDate: resolutionDate,
		//
		//status: status,
		//
		//notes: notes,
		//
	}
}
