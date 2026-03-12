//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package outbound

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/cooperlutz/go-full/internal/ownermanagement/app/query"
	"github.com/cooperlutz/go-full/internal/ownermanagement/domain/ownermanagement"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// PostgresAdapter implements the examination repository using Postgres as the data store.
type PostgresAdapter struct {
	Handler IQuerierOwnerManagement
}

// NewPostgresAdapter creates a new instance of PostgresAdapter.
func NewPostgresAdapter(db deebee.IDatabase) PostgresAdapter {
	return PostgresAdapter{
		Handler: NewQueriesWrapper(db),
	}
}

func (p PostgresAdapter) FindAllOwners(ctx context.Context) ([]query.Owner, error) {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.adapters.outbound.postgres.find_all_owner")
	defer span.End()

	owners, err := p.Handler.FindAllOwners(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return ownermanagementOwnersToQuery(owners)
}

func (p PostgresAdapter) FindOneOwner(ctx context.Context, id uuid.UUID) (query.Owner, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_owner")
	defer span.End()

	owner, err := p.GetOwner(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.Owner{}, err
	}

	return mapEntityOwnerToQuery(owner), nil
}

// AddOwner adds a new exam to the database.
func (p PostgresAdapter) AddOwner(ctx context.Context, owner *ownermanagement.Owner) error {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.adapters.outbound.postgres.add_owner")
	defer span.End()

	dbOwner := mapEntityOwnerToDB(owner)

	err := p.Handler.AddOwner(ctx, AddOwnerParams(dbOwner))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetOwner(ctx context.Context, id uuid.UUID) (*ownermanagement.Owner, error) {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.adapters.outbound.postgres.get_owner")
	defer span.End()

	owner, err := p.Handler.GetOwner(
		ctx,
		GetOwnerParams{OwnerID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return owner.toDomain()
}

func (p PostgresAdapter) UpdateOwner(
	ctx context.Context,
	ownerId uuid.UUID,
	updateFn func(e *ownermanagement.Owner) (*ownermanagement.Owner, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.adapters.outbound.postgres.update_owner")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	owner, err := p.GetOwner(ctx, ownerId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedOwner, err := updateFn(owner)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbOwner := mapEntityOwnerToDB(updatedOwner)

	err = p.Handler.UpdateOwner(ctx, UpdateOwnerParams(dbOwner))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) FindAllLoyaltyAccounts(ctx context.Context) ([]query.LoyaltyAccount, error) {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.adapters.outbound.postgres.find_all_loyalty_account")
	defer span.End()

	loyaltyaccounts, err := p.Handler.FindAllLoyaltyAccounts(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return ownermanagementLoyaltyAccountsToQuery(loyaltyaccounts)
}

func (p PostgresAdapter) FindOneLoyaltyAccount(ctx context.Context, id uuid.UUID) (query.LoyaltyAccount, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_loyalty_account")
	defer span.End()

	loyaltyaccount, err := p.GetLoyaltyAccount(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.LoyaltyAccount{}, err
	}

	return mapEntityLoyaltyAccountToQuery(loyaltyaccount), nil
}

// AddLoyaltyAccount adds a new exam to the database.
func (p PostgresAdapter) AddLoyaltyAccount(ctx context.Context, loyaltyaccount *ownermanagement.LoyaltyAccount) error {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.adapters.outbound.postgres.add_loyalty_account")
	defer span.End()

	dbLoyaltyAccount := mapEntityLoyaltyAccountToDB(loyaltyaccount)

	err := p.Handler.AddLoyaltyAccount(ctx, AddLoyaltyAccountParams(dbLoyaltyAccount))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetLoyaltyAccount(ctx context.Context, id uuid.UUID) (*ownermanagement.LoyaltyAccount, error) {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.adapters.outbound.postgres.get_loyalty_account")
	defer span.End()

	loyaltyaccount, err := p.Handler.GetLoyaltyAccount(
		ctx,
		GetLoyaltyAccountParams{LoyaltyAccountID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return loyaltyaccount.toDomain()
}

func (p PostgresAdapter) UpdateLoyaltyAccount(
	ctx context.Context,
	loyaltyaccountId uuid.UUID,
	updateFn func(e *ownermanagement.LoyaltyAccount) (*ownermanagement.LoyaltyAccount, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.adapters.outbound.postgres.update_loyaltyaccount")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	loyaltyaccount, err := p.GetLoyaltyAccount(ctx, loyaltyaccountId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedLoyaltyAccount, err := updateFn(loyaltyaccount)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbLoyaltyAccount := mapEntityLoyaltyAccountToDB(updatedLoyaltyAccount)

	err = p.Handler.UpdateLoyaltyAccount(ctx, UpdateLoyaltyAccountParams(dbLoyaltyAccount))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

// finishTransaction commits or rolls back the transaction based on the error state.
func (p PostgresAdapter) finishTransaction(ctx context.Context, err error, tx pgx.Tx) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			telemetree.RecordError(ctx, rollbackErr, "failed to rollback tx")

			return rollbackErr
		}

		return err
	} else {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			telemetree.RecordError(ctx, commitErr, "failed to commit tx")

			return commitErr
		}

		return nil
	}
}
