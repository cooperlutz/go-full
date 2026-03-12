package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/insuranceandclaims/domain/insuranceandclaims"
)

type RejectInsuranceClaim struct {
	//
	//ClaimId string,
	//
	//Reason string,
	//
	//ResolutionDate string,
	//
	// TODO
}

type RejectInsuranceClaimHandler struct {
	InsuranceProviderRepo insuranceandclaims.InsuranceProviderRepository

	InsuranceClaimRepo insuranceandclaims.InsuranceClaimRepository
}

func NewRejectInsuranceClaimHandler(
	insuranceproviderRepo insuranceandclaims.InsuranceProviderRepository,

	insuranceclaimRepo insuranceandclaims.InsuranceClaimRepository,
) RejectInsuranceClaimHandler {
	return RejectInsuranceClaimHandler{
		InsuranceProviderRepo: insuranceproviderRepo,

		InsuranceClaimRepo: insuranceclaimRepo,
	}
}

func (h RejectInsuranceClaimHandler) Handle(ctx context.Context, cmd RejectInsuranceClaim) error {
	// ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.app.command.reject_insurance_claim.handle")
	// defer span.End()

	// TODO
	//err = h.InsuranceProviderRepo.UpdateInsuranceProvider(ctx, uuid.MustParse(cmd.InsuranceProviderId), func(i *insuranceandclaims.InsuranceProvider) (*insuranceandclaims.InsuranceProvider, error) {
	//
	//	 err := i.RejectInsuranceClaim(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return i, nil
	//})
	//if err != nil {
	//	return err
	//}

	// TODO
	//err = h.InsuranceClaimRepo.UpdateInsuranceClaim(ctx, uuid.MustParse(cmd.InsuranceClaimId), func(i *insuranceandclaims.InsuranceClaim) (*insuranceandclaims.InsuranceClaim, error) {
	//
	//	 err := i.RejectInsuranceClaim(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return i, nil
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}
