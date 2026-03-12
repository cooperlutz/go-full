package inbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/insuranceandclaims/app"
	"github.com/cooperlutz/go-full/internal/insuranceandclaims/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// HttpAdapter represents the HTTP server for the InsuranceAndClaims module.
type HttpAdapter struct {
	app app.Application
}

// NewHttpAdapter creates a new HttpAdapter instance with the provided InsuranceAndClaims application.
func NewHttpAdapter(application app.Application) HttpAdapter {
	return HttpAdapter{
		app: application,
	}
}

// StrictHandler returns a strict HTTP handler for the InsuranceAndClaims module.
func (h HttpAdapter) StrictHandler() ServerInterface {
	return NewStrictHandler(h, nil)
}

// (GET /v1/insuranceproviders).
func (h HttpAdapter) FindAllInsuranceProviders(ctx context.Context, request FindAllInsuranceProvidersRequestObject) (FindAllInsuranceProvidersResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "insuranceprovider.adapters.inbound.http.find_all_insuranceproviders")
	defer span.End()

	insuranceprovider, err := h.app.Queries.FindAllInsuranceProviders.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseInsuranceProviders []InsuranceProvider
	for _, e := range insuranceprovider {
		responseInsuranceProviders = append(responseInsuranceProviders, queryInsuranceProviderToHttpInsuranceProvider(e))
	}

	return FindAllInsuranceProviders200JSONResponse(responseInsuranceProviders), nil
}

// (GET /v1/insuranceprovider/{insurance_providerId}).
func (h HttpAdapter) FindOneInsuranceProvider(ctx context.Context, request FindOneInsuranceProviderRequestObject) (FindOneInsuranceProviderResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_insurance_provider")
	defer span.End()

	insuranceprovider, err := h.app.Queries.FindOneInsuranceProvider.Handle(ctx, query.FindOneInsuranceProvider{InsuranceProviderID: request.InsuranceProviderId})
	if err != nil {
		return nil, err
	}

	return FindOneInsuranceProvider200JSONResponse(queryInsuranceProviderToHttpInsuranceProvider(insuranceprovider)), nil
}

func queryInsuranceProviderToHttpInsuranceProvider(e query.InsuranceProvider) InsuranceProvider {
	return InsuranceProvider{
		//
		//ProviderId: GetProviderId(),
		//
		//Name: GetName(),
		//
		//ContactName: GetContactName(),
		//
		//Email: GetEmail(),
		//
		//PhoneNumber: GetPhoneNumber(),
		//
		//ClaimSubmissionUrl: GetClaimSubmissionUrl(),
		//
		//Status: GetStatus(),
		//
		// TODO
	}
}

// (GET /v1/insuranceclaims).
func (h HttpAdapter) FindAllInsuranceClaims(ctx context.Context, request FindAllInsuranceClaimsRequestObject) (FindAllInsuranceClaimsResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "insuranceclaim.adapters.inbound.http.find_all_insuranceclaims")
	defer span.End()

	insuranceclaim, err := h.app.Queries.FindAllInsuranceClaims.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseInsuranceClaims []InsuranceClaim
	for _, e := range insuranceclaim {
		responseInsuranceClaims = append(responseInsuranceClaims, queryInsuranceClaimToHttpInsuranceClaim(e))
	}

	return FindAllInsuranceClaims200JSONResponse(responseInsuranceClaims), nil
}

// (GET /v1/insuranceclaim/{insurance_claimId}).
func (h HttpAdapter) FindOneInsuranceClaim(ctx context.Context, request FindOneInsuranceClaimRequestObject) (FindOneInsuranceClaimResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_insurance_claim")
	defer span.End()

	insuranceclaim, err := h.app.Queries.FindOneInsuranceClaim.Handle(ctx, query.FindOneInsuranceClaim{InsuranceClaimID: request.InsuranceClaimId})
	if err != nil {
		return nil, err
	}

	return FindOneInsuranceClaim200JSONResponse(queryInsuranceClaimToHttpInsuranceClaim(insuranceclaim)), nil
}

func queryInsuranceClaimToHttpInsuranceClaim(e query.InsuranceClaim) InsuranceClaim {
	return InsuranceClaim{
		//
		//ClaimId: GetClaimId(),
		//
		//OwnerId: GetOwnerId(),
		//
		//PetId: GetPetId(),
		//
		//ProviderId: GetProviderId(),
		//
		//InvoiceId: GetInvoiceId(),
		//
		//PolicyNumber: GetPolicyNumber(),
		//
		//ClaimAmount: GetClaimAmount(),
		//
		//ApprovedAmount: GetApprovedAmount(),
		//
		//SubmissionDate: GetSubmissionDate(),
		//
		//ResolutionDate: GetResolutionDate(),
		//
		//Status: GetStatus(),
		//
		//Notes: GetNotes(),
		//
		// TODO
	}
}
