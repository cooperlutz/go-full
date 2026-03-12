package outbound

import (
	"github.com/cooperlutz/go-full/internal/insuranceandclaims/app/query"
	"github.com/cooperlutz/go-full/internal/insuranceandclaims/domain/insuranceandclaims"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

// toDomain maps the InsuranceproviderInsuranceProvider to the domain entity.
func (e InsuranceandclaimsInsuranceProvider) toDomain() (*insuranceandclaims.InsuranceProvider, error) {
	return insuranceandclaims.MapToInsuranceProvider(
		e.InsuranceProviderID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.ProviderId,
		//
		//e.Name,
		//
		//e.ContactName,
		//
		//e.Email,
		//
		//e.PhoneNumber,
		//
		//e.ClaimSubmissionUrl,
		//
		//e.Status,
		//
		// TODO
	)
}

// toQueryInsuranceProvider maps the insuranceproviderInsuranceProvider to the query.InsuranceProvider.
func (e InsuranceandclaimsInsuranceProvider) toQueryInsuranceProvider() (query.InsuranceProvider, error) {
	insuranceprovider, err := e.toDomain()
	if err != nil {
		return query.InsuranceProvider{}, err
	}

	return mapEntityInsuranceProviderToQuery(insuranceprovider), nil
}

// insuranceproviderInsuranceProvidersToQuery maps a slice of InsuranceProviderInsuranceProvider to a slice of query.InsuranceProvider entities.
func insuranceandclaimsInsuranceProvidersToQuery(insuranceproviders []InsuranceandclaimsInsuranceProvider) ([]query.InsuranceProvider, error) {
	var domainInsuranceProviders []query.InsuranceProvider

	for _, insuranceprovider := range insuranceproviders {
		queryInsuranceProvider, err := insuranceprovider.toQueryInsuranceProvider()
		if err != nil {
			return nil, err
		}

		domainInsuranceProviders = append(domainInsuranceProviders, queryInsuranceProvider)
	}

	return domainInsuranceProviders, nil
}

// mapEntityInsuranceProviderToDB maps a domain InsuranceProvider entity to the InsuranceProviderInsuranceProvider database model.
func mapEntityInsuranceProviderToDB(insuranceprovider *insuranceandclaims.InsuranceProvider) InsuranceandclaimsInsuranceProvider {
	createdAt := insuranceprovider.GetCreatedAtTime()
	updatedAt := insuranceprovider.GetUpdatedAtTime()

	return InsuranceandclaimsInsuranceProvider{
		InsuranceProviderID: pgxutil.UUIDToPgtypeUUID(insuranceprovider.GetIdUUID()),
		CreatedAt:           pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:           pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:             insuranceprovider.IsDeleted(),
		DeletedAt:           pgxutil.TimeToTimestampz(insuranceprovider.GetDeletedAtTime()),
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

// mapEntityInsuranceProviderToQuery maps a domain InsuranceProvider entity to a query.InsuranceProvider.
func mapEntityInsuranceProviderToQuery(insuranceprovider *insuranceandclaims.InsuranceProvider) query.InsuranceProvider {
	return query.InsuranceProvider{
		// TODO
	}
}

// toDomain maps the InsuranceclaimInsuranceClaim to the domain entity.
func (e InsuranceandclaimsInsuranceClaim) toDomain() (*insuranceandclaims.InsuranceClaim, error) {
	return insuranceandclaims.MapToInsuranceClaim(
		e.InsuranceClaimID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.ClaimId,
		//
		//e.OwnerId,
		//
		//e.PetId,
		//
		//e.ProviderId,
		//
		//e.InvoiceId,
		//
		//e.PolicyNumber,
		//
		//e.ClaimAmount,
		//
		//e.ApprovedAmount,
		//
		//e.SubmissionDate,
		//
		//e.ResolutionDate,
		//
		//e.Status,
		//
		//e.Notes,
		//
		// TODO
	)
}

// toQueryInsuranceClaim maps the insuranceclaimInsuranceClaim to the query.InsuranceClaim.
func (e InsuranceandclaimsInsuranceClaim) toQueryInsuranceClaim() (query.InsuranceClaim, error) {
	insuranceclaim, err := e.toDomain()
	if err != nil {
		return query.InsuranceClaim{}, err
	}

	return mapEntityInsuranceClaimToQuery(insuranceclaim), nil
}

// insuranceclaimInsuranceClaimsToQuery maps a slice of InsuranceClaimInsuranceClaim to a slice of query.InsuranceClaim entities.
func insuranceandclaimsInsuranceClaimsToQuery(insuranceclaims []InsuranceandclaimsInsuranceClaim) ([]query.InsuranceClaim, error) {
	var domainInsuranceClaims []query.InsuranceClaim

	for _, insuranceclaim := range insuranceclaims {
		queryInsuranceClaim, err := insuranceclaim.toQueryInsuranceClaim()
		if err != nil {
			return nil, err
		}

		domainInsuranceClaims = append(domainInsuranceClaims, queryInsuranceClaim)
	}

	return domainInsuranceClaims, nil
}

// mapEntityInsuranceClaimToDB maps a domain InsuranceClaim entity to the InsuranceClaimInsuranceClaim database model.
func mapEntityInsuranceClaimToDB(insuranceclaim *insuranceandclaims.InsuranceClaim) InsuranceandclaimsInsuranceClaim {
	createdAt := insuranceclaim.GetCreatedAtTime()
	updatedAt := insuranceclaim.GetUpdatedAtTime()

	return InsuranceandclaimsInsuranceClaim{
		InsuranceClaimID: pgxutil.UUIDToPgtypeUUID(insuranceclaim.GetIdUUID()),
		CreatedAt:        pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:        pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:          insuranceclaim.IsDeleted(),
		DeletedAt:        pgxutil.TimeToTimestampz(insuranceclaim.GetDeletedAtTime()),
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

// mapEntityInsuranceClaimToQuery maps a domain InsuranceClaim entity to a query.InsuranceClaim.
func mapEntityInsuranceClaimToQuery(insuranceclaim *insuranceandclaims.InsuranceClaim) query.InsuranceClaim {
	return query.InsuranceClaim{
		// TODO
	}
}
