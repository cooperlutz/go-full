package persist

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/repository"
	persist_postgres "github.com/cooperlutz/go-full/internal/pingpong/infra/persist/postgres"
	"github.com/cooperlutz/go-full/pkg/deebee"
)

// Ensure PingPongPersistPostgresRepository implements the IPingPongRepository interface.
var _ repository.IPingPongRepository = (*PingPongPersistPostgresRepository)(nil)

// PingPongPersistPostgresRepository is a repository for managing PingPong entities in a PostgreSQL database.
//
// It implements the entity.PingPongRepository interface.
type PingPongPersistPostgresRepository struct {
	db    deebee.IDatabase
	query persist_postgres.IQuerierPingPong
}

// NewPingPongPostgresRepo creates a new instance of PingPongPersistPostgresRepository.
func NewPingPongPostgresRepo(pgconn *pgxpool.Pool) *PingPongPersistPostgresRepository {
	return &PingPongPersistPostgresRepository{
		db:    pgconn,
		query: persist_postgres.NewQuerysWrapper(pgconn),
	}
}
