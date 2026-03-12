package outbound

import (
	"github.com/cooperlutz/go-full/internal/ownermanagement/app/query"
	"github.com/cooperlutz/go-full/internal/ownermanagement/domain/ownermanagement"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

// toDomain maps the OwnerOwner to the domain entity.
func (e OwnermanagementOwner) toDomain() (*ownermanagement.Owner, error) {
	return ownermanagement.MapToOwner(
		e.OwnerID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.OwnerId,
		//
		//e.FirstName,
		//
		//e.LastName,
		//
		//e.Email,
		//
		//e.PhoneNumber,
		//
		//e.Address,
		//
		//e.CommunicationPreference,
		//
		//e.LoyaltyMember,
		//
		//e.LoyaltyPoints,
		//
		//e.Status,
		//
		// TODO
	)
}

// toQueryOwner maps the ownerOwner to the query.Owner.
func (e OwnermanagementOwner) toQueryOwner() (query.Owner, error) {
	owner, err := e.toDomain()
	if err != nil {
		return query.Owner{}, err
	}

	return mapEntityOwnerToQuery(owner), nil
}

// ownerOwnersToQuery maps a slice of OwnerOwner to a slice of query.Owner entities.
func ownermanagementOwnersToQuery(owners []OwnermanagementOwner) ([]query.Owner, error) {
	var domainOwners []query.Owner

	for _, owner := range owners {
		queryOwner, err := owner.toQueryOwner()
		if err != nil {
			return nil, err
		}

		domainOwners = append(domainOwners, queryOwner)
	}

	return domainOwners, nil
}

// mapEntityOwnerToDB maps a domain Owner entity to the OwnerOwner database model.
func mapEntityOwnerToDB(owner *ownermanagement.Owner) OwnermanagementOwner {
	createdAt := owner.GetCreatedAtTime()
	updatedAt := owner.GetUpdatedAtTime()

	return OwnermanagementOwner{
		OwnerID:   pgxutil.UUIDToPgtypeUUID(owner.GetIdUUID()),
		CreatedAt: pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt: pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:   owner.IsDeleted(),
		DeletedAt: pgxutil.TimeToTimestampz(owner.GetDeletedAtTime()),
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

// mapEntityOwnerToQuery maps a domain Owner entity to a query.Owner.
func mapEntityOwnerToQuery(owner *ownermanagement.Owner) query.Owner {
	return query.Owner{
		// TODO
	}
}

// toDomain maps the LoyaltyaccountLoyaltyAccount to the domain entity.
func (e OwnermanagementLoyaltyAccount) toDomain() (*ownermanagement.LoyaltyAccount, error) {
	return ownermanagement.MapToLoyaltyAccount(
		e.LoyaltyAccountID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.LoyaltyAccountId,
		//
		//e.OwnerId,
		//
		//e.PointsBalance,
		//
		//e.Tier,
		//
		//e.EnrolledDate,
		//
		// TODO
	)
}

// toQueryLoyaltyAccount maps the loyaltyaccountLoyaltyAccount to the query.LoyaltyAccount.
func (e OwnermanagementLoyaltyAccount) toQueryLoyaltyAccount() (query.LoyaltyAccount, error) {
	loyaltyaccount, err := e.toDomain()
	if err != nil {
		return query.LoyaltyAccount{}, err
	}

	return mapEntityLoyaltyAccountToQuery(loyaltyaccount), nil
}

// loyaltyaccountLoyaltyAccountsToQuery maps a slice of LoyaltyAccountLoyaltyAccount to a slice of query.LoyaltyAccount entities.
func ownermanagementLoyaltyAccountsToQuery(loyaltyaccounts []OwnermanagementLoyaltyAccount) ([]query.LoyaltyAccount, error) {
	var domainLoyaltyAccounts []query.LoyaltyAccount

	for _, loyaltyaccount := range loyaltyaccounts {
		queryLoyaltyAccount, err := loyaltyaccount.toQueryLoyaltyAccount()
		if err != nil {
			return nil, err
		}

		domainLoyaltyAccounts = append(domainLoyaltyAccounts, queryLoyaltyAccount)
	}

	return domainLoyaltyAccounts, nil
}

// mapEntityLoyaltyAccountToDB maps a domain LoyaltyAccount entity to the LoyaltyAccountLoyaltyAccount database model.
func mapEntityLoyaltyAccountToDB(loyaltyaccount *ownermanagement.LoyaltyAccount) OwnermanagementLoyaltyAccount {
	createdAt := loyaltyaccount.GetCreatedAtTime()
	updatedAt := loyaltyaccount.GetUpdatedAtTime()

	return OwnermanagementLoyaltyAccount{
		LoyaltyAccountID: pgxutil.UUIDToPgtypeUUID(loyaltyaccount.GetIdUUID()),
		CreatedAt:        pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:        pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:          loyaltyaccount.IsDeleted(),
		DeletedAt:        pgxutil.TimeToTimestampz(loyaltyaccount.GetDeletedAtTime()),
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

// mapEntityLoyaltyAccountToQuery maps a domain LoyaltyAccount entity to a query.LoyaltyAccount.
func mapEntityLoyaltyAccountToQuery(loyaltyaccount *ownermanagement.LoyaltyAccount) query.LoyaltyAccount {
	return query.LoyaltyAccount{
		// TODO
	}
}
