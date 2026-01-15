package outbound

import (
	"github.com/jackc/pgx/v5"
)

// Ensure QueriesWrapper implements the IQuerierExamination interface.
var _ IQuerierExamination = (*QueriesWrapper)(nil)

// IQuerierExamination wraps the sqlc Querier interface and adds the `WithTx` method.
type IQuerierExamination interface {
	Querier
	WithTx(tx pgx.Tx) IQuerierExamination
}

// https://github.com/forkd-app/forkd/commit/2822d6ac1eac4b378e9ef99fb8a80041070c9f52
func NewQueriesWrapper(db DBTX) IQuerierExamination {
	return &QueriesWrapper{
		Queries: New(db),
	}
}

// QueriesWrapper implements IQuerierPingPong.
type QueriesWrapper struct {
	*Queries
}

func (q *QueriesWrapper) WithTx(tx pgx.Tx) IQuerierExamination {
	return &QueriesWrapper{
		Queries: q.Queries.WithTx(tx),
	}
}
