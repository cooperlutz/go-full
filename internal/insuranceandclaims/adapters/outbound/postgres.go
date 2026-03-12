//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package outbound

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/cooperlutz/go-full/internal/insuranceandclaims/app/query"
	"github.com/cooperlutz/go-full/internal/insuranceandclaims/domain/insuranceandclaims"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// PostgresAdapter implements the examination repository using Postgres as the data store.
type PostgresAdapter struct {
	Handler IQuerierInsuranceAndClaims
}

// NewPostgresAdapter creates a new instance of PostgresAdapter.
func NewPostgresAdapter(db deebee.IDatabase) PostgresAdapter {
	return PostgresAdapter{
		Handler: NewQueriesWrapper(db),
	}
}

func (p PostgresAdapter) FindAllInsuranceProviders(ctx context.Context) ([]query.InsuranceProvider, error) {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.adapters.outbound.postgres.find_all_insurance_provider")
	defer span.End()

	insuranceproviders, err := p.Handler.FindAllInsuranceProviders(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return insuranceandclaimsInsuranceProvidersToQuery(insuranceproviders)
}

func (p PostgresAdapter) FindOneInsuranceProvider(ctx context.Context, id uuid.UUID) (query.InsuranceProvider, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_insurance_provider")
	defer span.End()

	insuranceprovider, err := p.GetInsuranceProvider(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.InsuranceProvider{}, err
	}

	return mapEntityInsuranceProviderToQuery(insuranceprovider), nil
}

// AddInsuranceProvider adds a new exam to the database.
func (p PostgresAdapter) AddInsuranceProvider(ctx context.Context, insuranceprovider *insuranceandclaims.InsuranceProvider) error {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.adapters.outbound.postgres.add_insurance_provider")
	defer span.End()

	dbInsuranceProvider := mapEntityInsuranceProviderToDB(insuranceprovider)

	err := p.Handler.AddInsuranceProvider(ctx, AddInsuranceProviderParams(dbInsuranceProvider))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetInsuranceProvider(ctx context.Context, id uuid.UUID) (*insuranceandclaims.InsuranceProvider, error) {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.adapters.outbound.postgres.get_insurance_provider")
	defer span.End()

	insuranceprovider, err := p.Handler.GetInsuranceProvider(
		ctx,
		GetInsuranceProviderParams{InsuranceProviderID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return insuranceprovider.toDomain()
}

func (p PostgresAdapter) UpdateInsuranceProvider(
	ctx context.Context,
	insuranceproviderId uuid.UUID,
	updateFn func(e *insuranceandclaims.InsuranceProvider) (*insuranceandclaims.InsuranceProvider, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.adapters.outbound.postgres.update_insuranceprovider")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	insuranceprovider, err := p.GetInsuranceProvider(ctx, insuranceproviderId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedInsuranceProvider, err := updateFn(insuranceprovider)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbInsuranceProvider := mapEntityInsuranceProviderToDB(updatedInsuranceProvider)

	err = p.Handler.UpdateInsuranceProvider(ctx, UpdateInsuranceProviderParams(dbInsuranceProvider))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) FindAllInsuranceClaims(ctx context.Context) ([]query.InsuranceClaim, error) {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.adapters.outbound.postgres.find_all_insurance_claim")
	defer span.End()

	insuranceclaims, err := p.Handler.FindAllInsuranceClaims(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return insuranceandclaimsInsuranceClaimsToQuery(insuranceclaims)
}

func (p PostgresAdapter) FindOneInsuranceClaim(ctx context.Context, id uuid.UUID) (query.InsuranceClaim, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_insurance_claim")
	defer span.End()

	insuranceclaim, err := p.GetInsuranceClaim(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.InsuranceClaim{}, err
	}

	return mapEntityInsuranceClaimToQuery(insuranceclaim), nil
}

// AddInsuranceClaim adds a new exam to the database.
func (p PostgresAdapter) AddInsuranceClaim(ctx context.Context, insuranceclaim *insuranceandclaims.InsuranceClaim) error {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.adapters.outbound.postgres.add_insurance_claim")
	defer span.End()

	dbInsuranceClaim := mapEntityInsuranceClaimToDB(insuranceclaim)

	err := p.Handler.AddInsuranceClaim(ctx, AddInsuranceClaimParams(dbInsuranceClaim))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetInsuranceClaim(ctx context.Context, id uuid.UUID) (*insuranceandclaims.InsuranceClaim, error) {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.adapters.outbound.postgres.get_insurance_claim")
	defer span.End()

	insuranceclaim, err := p.Handler.GetInsuranceClaim(
		ctx,
		GetInsuranceClaimParams{InsuranceClaimID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return insuranceclaim.toDomain()
}

func (p PostgresAdapter) UpdateInsuranceClaim(
	ctx context.Context,
	insuranceclaimId uuid.UUID,
	updateFn func(e *insuranceandclaims.InsuranceClaim) (*insuranceandclaims.InsuranceClaim, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.adapters.outbound.postgres.update_insuranceclaim")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	insuranceclaim, err := p.GetInsuranceClaim(ctx, insuranceclaimId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedInsuranceClaim, err := updateFn(insuranceclaim)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbInsuranceClaim := mapEntityInsuranceClaimToDB(updatedInsuranceClaim)

	err = p.Handler.UpdateInsuranceClaim(ctx, UpdateInsuranceClaimParams(dbInsuranceClaim))
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
