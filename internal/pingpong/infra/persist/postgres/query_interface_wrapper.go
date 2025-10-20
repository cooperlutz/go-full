package persist_postgres

import (
	"github.com/jackc/pgx/v5"
)

// Ensure QueriesWrapper implements the IQuerierPingPong interface.
var _ IQuerierPingPong = (*QueriesWrapper)(nil)

// IQuerierPingPong wraps the sqlc Querier interface and adds the `WithTx` method.
type IQuerierPingPong interface {
	Querier
	WithTx(tx pgx.Tx) IQuerierPingPong
}

// https://github.com/forkd-app/forkd/commit/2822d6ac1eac4b378e9ef99fb8a80041070c9f52
func NewQuerysWrapper(db DBTX) IQuerierPingPong {
	return &QueriesWrapper{
		Queries: New(db),
	}
}

// QueriesWrapper implements IQuerierPingPong.
type QueriesWrapper struct {
	*Queries
}

func (q *QueriesWrapper) WithTx(tx pgx.Tx) IQuerierPingPong {
	return &QueriesWrapper{
		Queries: q.Queries.WithTx(tx),
	}
}

// func GetQueriesWithConnection(connectionString string) (QueryWrapper, *pgxpool.Pool, error) {
// 	ctx := context.Background()
// 	pool, err := pgxpool.New(ctx, connectionString)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	if err := pool.Ping(ctx); err != nil {
// 		return nil, nil, err
// 	}

// 	queries := New(pool)
// 	// Wrap the sqlc Queries struct in our wrapper so we can have an interface with the `WithTx` method
// 	return &QueriesWrapper{
// 		Queries: queries,
// 	}, pool, nil
// }
