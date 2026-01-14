/*
The Exam Library Module encapuslates the functionality related to the Exam Library bounded context.

The Architecture follows a layered / Clean approach, separating concerns into distinct layers:

1. API Layer: Contains HTTP handlers and routing logic.

2. Domain Layer: Contains the core business logic and domain entities.

3. Application Layer: Contains use cases that orchestrate domain entities to fulfill application requirements.

4. Infrastructure Layer: Contains implementations for data persistence and Pub Sub integrations.

Each layer interacts only with the layer directly below it, ensuring a clear separation of concerns and maintainability.
Dependencies are inverted using interfaces to promote loose coupling between layers.
*/
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

// NewModule - Initializes the Exam Library module with its needed dependencies.
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
