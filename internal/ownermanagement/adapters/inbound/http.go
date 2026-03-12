package inbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/ownermanagement/app"
	"github.com/cooperlutz/go-full/internal/ownermanagement/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// HttpAdapter represents the HTTP server for the OwnerManagement module.
type HttpAdapter struct {
	app app.Application
}

// NewHttpAdapter creates a new HttpAdapter instance with the provided OwnerManagement application.
func NewHttpAdapter(application app.Application) HttpAdapter {
	return HttpAdapter{
		app: application,
	}
}

// StrictHandler returns a strict HTTP handler for the OwnerManagement module.
func (h HttpAdapter) StrictHandler() ServerInterface {
	return NewStrictHandler(h, nil)
}

// (GET /v1/owners).
func (h HttpAdapter) FindAllOwners(ctx context.Context, request FindAllOwnersRequestObject) (FindAllOwnersResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "owner.adapters.inbound.http.find_all_owners")
	defer span.End()

	owner, err := h.app.Queries.FindAllOwners.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseOwners []Owner
	for _, e := range owner {
		responseOwners = append(responseOwners, queryOwnerToHttpOwner(e))
	}

	return FindAllOwners200JSONResponse(responseOwners), nil
}

// (GET /v1/owner/{ownerId}).
func (h HttpAdapter) FindOneOwner(ctx context.Context, request FindOneOwnerRequestObject) (FindOneOwnerResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_owner")
	defer span.End()

	owner, err := h.app.Queries.FindOneOwner.Handle(ctx, query.FindOneOwner{OwnerID: request.OwnerId})
	if err != nil {
		return nil, err
	}

	return FindOneOwner200JSONResponse(queryOwnerToHttpOwner(owner)), nil
}

func queryOwnerToHttpOwner(e query.Owner) Owner {
	return Owner{
		//
		//OwnerId: GetOwnerId(),
		//
		//FirstName: GetFirstName(),
		//
		//LastName: GetLastName(),
		//
		//Email: GetEmail(),
		//
		//PhoneNumber: GetPhoneNumber(),
		//
		//Address: GetAddress(),
		//
		//CommunicationPreference: GetCommunicationPreference(),
		//
		//LoyaltyMember: GetLoyaltyMember(),
		//
		//LoyaltyPoints: GetLoyaltyPoints(),
		//
		//Status: GetStatus(),
		//
		// TODO
	}
}

// (GET /v1/loyaltyaccounts).
func (h HttpAdapter) FindAllLoyaltyAccounts(ctx context.Context, request FindAllLoyaltyAccountsRequestObject) (FindAllLoyaltyAccountsResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "loyaltyaccount.adapters.inbound.http.find_all_loyaltyaccounts")
	defer span.End()

	loyaltyaccount, err := h.app.Queries.FindAllLoyaltyAccounts.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseLoyaltyAccounts []LoyaltyAccount
	for _, e := range loyaltyaccount {
		responseLoyaltyAccounts = append(responseLoyaltyAccounts, queryLoyaltyAccountToHttpLoyaltyAccount(e))
	}

	return FindAllLoyaltyAccounts200JSONResponse(responseLoyaltyAccounts), nil
}

// (GET /v1/loyaltyaccount/{loyalty_accountId}).
func (h HttpAdapter) FindOneLoyaltyAccount(ctx context.Context, request FindOneLoyaltyAccountRequestObject) (FindOneLoyaltyAccountResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_loyalty_account")
	defer span.End()

	loyaltyaccount, err := h.app.Queries.FindOneLoyaltyAccount.Handle(ctx, query.FindOneLoyaltyAccount{LoyaltyAccountID: request.LoyaltyAccountId})
	if err != nil {
		return nil, err
	}

	return FindOneLoyaltyAccount200JSONResponse(queryLoyaltyAccountToHttpLoyaltyAccount(loyaltyaccount)), nil
}

func queryLoyaltyAccountToHttpLoyaltyAccount(e query.LoyaltyAccount) LoyaltyAccount {
	return LoyaltyAccount{
		//
		//LoyaltyAccountId: GetLoyaltyAccountId(),
		//
		//OwnerId: GetOwnerId(),
		//
		//PointsBalance: GetPointsBalance(),
		//
		//Tier: GetTier(),
		//
		//EnrolledDate: GetEnrolledDate(),
		//
		// TODO
	}
}
