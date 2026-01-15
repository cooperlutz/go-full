package app

import (
	"github.com/cooperlutz/go-full/internal/examination/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/examination/app/query"
	"github.com/cooperlutz/go-full/pkg/deebee"
)

type Application struct {
	Queries Queries
}

type Queries struct {
	AvailableExams query.AvailableExamsHandler
}

// NewApplication initializes the Examination application with its dependencies.
func NewApplication(pgconn deebee.IDatabase) Application {
	examinationRepository := outbound.NewPostgresAdapter(pgconn)
	app := Application{
		Queries: Queries{
			AvailableExams: query.NewAvailableExamsHandler(examinationRepository),
		},
	}

	return app
}
