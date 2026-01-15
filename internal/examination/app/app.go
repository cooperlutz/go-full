package app

import (
	"github.com/cooperlutz/go-full/internal/examination/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/examination/app/query"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type Application struct {
	Queries Queries
}

type Queries struct {
	AvailableExams query.AvailableExamsHandler
}

func NewApplication(pgconn deebee.IDatabase, basePS *eeventdriven.BasePgsqlPubSubProcessor) Application {
	examinationRepository := outbound.New(pgconn)
	app := Application{
		Queries: Queries{
			AvailableExams: query.NewAvailableExamsHandler(examinationRepository),
		},
	}

	return app
}
