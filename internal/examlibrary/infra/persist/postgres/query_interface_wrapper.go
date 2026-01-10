package persist_postgres

import (
	"github.com/jackc/pgx/v5"
)

// Ensure QueriesWrapper implements the IQuerierExamLibrary interface.
var _ IQuerierExamLibrary = (*QueriesWrapper)(nil)

// IQuerierExamLibrary wraps the sqlc Querier interface and adds the `WithTx` method.
type IQuerierExamLibrary interface {
	Querier
	WithTx(tx pgx.Tx) IQuerierExamLibrary
}

// https://github.com/forkd-app/forkd/commit/2822d6ac1eac4b378e9ef99fb8a80041070c9f52
func NewQueriesWrapper(db DBTX) IQuerierExamLibrary {
	return &QueriesWrapper{
		Queries: New(db),
	}
}

type QueriesWrapper struct {
	*Queries
}

func (q *QueriesWrapper) WithTx(tx pgx.Tx) IQuerierExamLibrary {
	return &QueriesWrapper{
		Queries: q.Queries.WithTx(tx),
	}
}
