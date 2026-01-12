package examlibrary

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/repository"
	"github.com/cooperlutz/go-full/internal/examlibrary/infra/persist"
	"github.com/cooperlutz/go-full/internal/examlibrary/infra/pubsub"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

// ExamLibraryModule encapsulates the Exam Library module's components, dependencies, interfaces
// that is implemented within the core application to provide Exam Library functionality.
type ExamLibraryModule struct {
	PersistentRepo repository.IExamLibraryRepository
	UseCase        usecase.IExamLibraryUseCase
	RestApi        http.Handler
	PubSub         eeventdriven.IPubSubEventProcessor
}

// NewModule - Initializes the PingPong module with its needed dependencies.
func NewModule(pgconn *pgxpool.Pool) (*ExamLibraryModule, error) {
	repo := persist.NewExamLibraryPostgresRepo(pgconn)

	ps, err := pubsub.New(pgconn, repo)
	if err != nil {
		return nil, err
	}

	uc := usecase.NewExamLibraryUseCase(repo, ps)
	api := rest.NewExamLibraryAPIRouter(uc)

	module := &ExamLibraryModule{
		PersistentRepo: repo,
		UseCase:        uc,
		RestApi:        api,
		PubSub:         ps,
	}

	err = ps.RegisterSubscriberHandlers()
	if err != nil {
		return nil, err
	}

	return module, nil
}
