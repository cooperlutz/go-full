package persist

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/repository"
	persist_postgres "github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/postgres"
	"github.com/cooperlutz/go-full/pkg/deebee"
)

// Ensure examLibraryPersistPostgresRepository implements the IExamLibraryRepository interface.
var _ repository.IExamLibraryRepository = (*examLibraryPersistPostgresRepository)(nil)

// examLibraryPersistPostgresRepository is a repository for managing ExamLibrary entities in a PostgreSQL database.
//
// It implements the entity.ExamLibraryRepository interface.
type examLibraryPersistPostgresRepository struct {
	db    deebee.IDatabase
	query persist_postgres.IQuerierExamLibrary
}

// NewExamLibraryPostgresRepo creates a new instance of examLibraryPersistPostgresRepository.
func NewExamLibraryPostgresRepo(pgconn *pgxpool.Pool) *examLibraryPersistPostgresRepository {
	return &examLibraryPersistPostgresRepository{
		db:    pgconn,
		query: persist_postgres.NewQueriesWrapper(pgconn),
	}
}
