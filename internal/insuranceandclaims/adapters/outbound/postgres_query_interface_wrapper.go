package outbound

import (
	"github.com/jackc/pgx/v5"

	"github.com/cooperlutz/go-full/pkg/deebee"
)

// Ensure QueriesWrapper implements the IQuerierInsuranceAndClaims interface.
var _ IQuerierInsuranceAndClaims = (*QueriesWrapper)(nil)

// IQuerierInsuranceAndClaims wraps the sqlc Querier interface and adds the `WithTx` method.
type IQuerierInsuranceAndClaims interface {
	Querier
	WithTx(tx pgx.Tx) IQuerierInsuranceAndClaims
	deebee.IDatabase
}

// https://github.com/forkd-app/forkd/commit/2822d6ac1eac4b378e9ef99fb8a80041070c9f52
func NewQueriesWrapper(db deebee.IDatabase) IQuerierInsuranceAndClaims {
	return &QueriesWrapper{
		Queries:   New(db),
		IDatabase: db,
	}
}

// QueriesWrapper implements IQuerierInsuranceAndClaims.
type QueriesWrapper struct {
	*Queries
	deebee.IDatabase
}

func (q *QueriesWrapper) WithTx(tx pgx.Tx) IQuerierInsuranceAndClaims {
	return &QueriesWrapper{
		Queries:   q.Queries.WithTx(tx),
		IDatabase: q.IDatabase,
	}
}
