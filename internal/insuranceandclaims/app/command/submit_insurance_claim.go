package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/insuranceandclaims/domain/insuranceandclaims"
)

type SubmitInsuranceClaim struct {
	//
	//OwnerId string,
	//
	//PetId string,
	//
	//ProviderId string,
	//
	//InvoiceId string,
	//
	//PolicyNumber string,
	//
	//ClaimAmount float32,
	//
	// TODO
}

type SubmitInsuranceClaimHandler struct {
	InsuranceProviderRepo insuranceandclaims.InsuranceProviderRepository

	InsuranceClaimRepo insuranceandclaims.InsuranceClaimRepository
}

func NewSubmitInsuranceClaimHandler(
	insuranceproviderRepo insuranceandclaims.InsuranceProviderRepository,

	insuranceclaimRepo insuranceandclaims.InsuranceClaimRepository,
) SubmitInsuranceClaimHandler {
	return SubmitInsuranceClaimHandler{
		InsuranceProviderRepo: insuranceproviderRepo,

		InsuranceClaimRepo: insuranceclaimRepo,
	}
}

func (h SubmitInsuranceClaimHandler) Handle(ctx context.Context, cmd SubmitInsuranceClaim) error {
	// ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.app.command.submit_insurance_claim.handle")
	// defer span.End()

	// TODO
	//err = h.InsuranceProviderRepo.UpdateInsuranceProvider(ctx, uuid.MustParse(cmd.InsuranceProviderId), func(i *insuranceandclaims.InsuranceProvider) (*insuranceandclaims.InsuranceProvider, error) {
	//
	//	 err := i.SubmitInsuranceClaim(
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
	//	 err := i.SubmitInsuranceClaim(
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
