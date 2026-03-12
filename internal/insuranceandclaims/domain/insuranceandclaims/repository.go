package insuranceandclaims

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type InsuranceProviderRepository interface {
	AddInsuranceProvider(ctx context.Context, insuranceprovider *InsuranceProvider) error

	GetInsuranceProvider(ctx context.Context, id uuid.UUID) (*InsuranceProvider, error)

	UpdateInsuranceProvider(
		ctx context.Context,
		insuranceproviderId uuid.UUID,
		updateFn func(e *InsuranceProvider) (*InsuranceProvider, error),
	) error
}

// MapToInsuranceProvider creates a InsuranceProvider domain object from the given parameters.
// This should ONLY BE USED when reconstructing an InsuranceProvider from its repository.
func MapToInsuranceProvider(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//providerId string,
	//
	//name string,
	//
	//contactName *string,
	//
	//email string,
	//
	//phoneNumber string,
	//
	//claimSubmissionUrl *string,
	//
	//status string,
	//
) (*InsuranceProvider, error) {
	return &InsuranceProvider{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
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
		// TODO
	}, nil
}

type InsuranceClaimRepository interface {
	AddInsuranceClaim(ctx context.Context, insuranceclaim *InsuranceClaim) error

	GetInsuranceClaim(ctx context.Context, id uuid.UUID) (*InsuranceClaim, error)

	UpdateInsuranceClaim(
		ctx context.Context,
		insuranceclaimId uuid.UUID,
		updateFn func(e *InsuranceClaim) (*InsuranceClaim, error),
	) error
}

// MapToInsuranceClaim creates a InsuranceClaim domain object from the given parameters.
// This should ONLY BE USED when reconstructing an InsuranceClaim from its repository.
func MapToInsuranceClaim(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//claimId string,
	//
	//ownerId string,
	//
	//petId string,
	//
	//providerId string,
	//
	//invoiceId string,
	//
	//policyNumber string,
	//
	//claimAmount float32,
	//
	//approvedAmount *float32,
	//
	//submissionDate string,
	//
	//resolutionDate *string,
	//
	//status string,
	//
	//notes *string,
	//
) (*InsuranceClaim, error) {
	return &InsuranceClaim{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
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
		// TODO
	}, nil
}
