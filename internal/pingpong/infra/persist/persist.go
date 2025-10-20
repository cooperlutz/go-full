package persist

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/repository"
	persist_postgres "github.com/cooperlutz/go-full/internal/pingpong/infra/persist/postgres"
	"github.com/cooperlutz/go-full/pkg/deebee"
)

// Ensure pingPongPersistPostgresRepository implements the IPingPongRepository interface.
var _ repository.IPingPongRepository = (*pingPongPersistPostgresRepository)(nil)

// pingPongPersistPostgresRepository is a repository for managing PingPong entities in a PostgreSQL database.
//
// It implements the entity.PingPongRepository interface.
type pingPongPersistPostgresRepository struct {
	db    deebee.IDatabase
	query persist_postgres.IQuerierPingPong
}

// NewPingPongPostgresRepo creates a new instance of pingPongPersistPostgresRepository.
func NewPingPongPostgresRepo(pgconn *pgxpool.Pool) *pingPongPersistPostgresRepository {
	return &pingPongPersistPostgresRepository{
		db:    pgconn,
		query: persist_postgres.NewQuerysWrapper(pgconn),
	}
}
