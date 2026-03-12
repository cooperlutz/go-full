package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/insuranceandclaims/domain/insuranceandclaims"
)

type RegisterInsuranceProvider struct {
	//
	//Name string,
	//
	//ContactName *string,
	//
	//Email string,
	//
	//PhoneNumber string,
	//
	//ClaimSubmissionUrl *string,
	//
	// TODO
}

type RegisterInsuranceProviderHandler struct {
	InsuranceProviderRepo insuranceandclaims.InsuranceProviderRepository

	InsuranceClaimRepo insuranceandclaims.InsuranceClaimRepository
}

func NewRegisterInsuranceProviderHandler(
	insuranceproviderRepo insuranceandclaims.InsuranceProviderRepository,

	insuranceclaimRepo insuranceandclaims.InsuranceClaimRepository,
) RegisterInsuranceProviderHandler {
	return RegisterInsuranceProviderHandler{
		InsuranceProviderRepo: insuranceproviderRepo,

		InsuranceClaimRepo: insuranceclaimRepo,
	}
}

func (h RegisterInsuranceProviderHandler) Handle(ctx context.Context, cmd RegisterInsuranceProvider) error {
	// ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.app.command.register_insurance_provider.handle")
	// defer span.End()

	// TODO
	//err = h.InsuranceProviderRepo.UpdateInsuranceProvider(ctx, uuid.MustParse(cmd.InsuranceProviderId), func(i *insuranceandclaims.InsuranceProvider) (*insuranceandclaims.InsuranceProvider, error) {
	//
	//	 err := i.RegisterInsuranceProvider(
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
	//	 err := i.RegisterInsuranceProvider(
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
