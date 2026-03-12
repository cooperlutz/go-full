package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/insuranceandclaims/domain/insuranceandclaims"
)

type ApproveInsuranceClaim struct {
	//
	//ClaimId string,
	//
	//ApprovedAmount float32,
	//
	//ResolutionDate string,
	//
	// TODO
}

type ApproveInsuranceClaimHandler struct {
	InsuranceProviderRepo insuranceandclaims.InsuranceProviderRepository

	InsuranceClaimRepo insuranceandclaims.InsuranceClaimRepository
}

func NewApproveInsuranceClaimHandler(
	insuranceproviderRepo insuranceandclaims.InsuranceProviderRepository,

	insuranceclaimRepo insuranceandclaims.InsuranceClaimRepository,
) ApproveInsuranceClaimHandler {
	return ApproveInsuranceClaimHandler{
		InsuranceProviderRepo: insuranceproviderRepo,

		InsuranceClaimRepo: insuranceclaimRepo,
	}
}

func (h ApproveInsuranceClaimHandler) Handle(ctx context.Context, cmd ApproveInsuranceClaim) error {
	// ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.app.command.approve_insurance_claim.handle")
	// defer span.End()

	// TODO
	//err = h.InsuranceProviderRepo.UpdateInsuranceProvider(ctx, uuid.MustParse(cmd.InsuranceProviderId), func(i *insuranceandclaims.InsuranceProvider) (*insuranceandclaims.InsuranceProvider, error) {
	//
	//	 err := i.ApproveInsuranceClaim(
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
	//	 err := i.ApproveInsuranceClaim(
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
