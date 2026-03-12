package outbound

import (
	"github.com/jackc/pgx/v5"

	"github.com/cooperlutz/go-full/pkg/deebee"
)

// Ensure QueriesWrapper implements the IQuerierBillingAndPayments interface.
var _ IQuerierBillingAndPayments = (*QueriesWrapper)(nil)

// IQuerierBillingAndPayments wraps the sqlc Querier interface and adds the `WithTx` method.
type IQuerierBillingAndPayments interface {
	Querier
	WithTx(tx pgx.Tx) IQuerierBillingAndPayments
	deebee.IDatabase
}

// https://github.com/forkd-app/forkd/commit/2822d6ac1eac4b378e9ef99fb8a80041070c9f52
func NewQueriesWrapper(db deebee.IDatabase) IQuerierBillingAndPayments {
	return &QueriesWrapper{
		Queries:   New(db),
		IDatabase: db,
	}
}

// QueriesWrapper implements IQuerierBillingAndPayments.
type QueriesWrapper struct {
	*Queries
	deebee.IDatabase
}

func (q *QueriesWrapper) WithTx(tx pgx.Tx) IQuerierBillingAndPayments {
	return &QueriesWrapper{
		Queries:   q.Queries.WithTx(tx),
		IDatabase: q.IDatabase,
	}
}
