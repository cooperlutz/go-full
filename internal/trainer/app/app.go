package app

import (
	"context"

	"github.com/cooperlutz/go-full/internal/trainer/adapters"
	"github.com/cooperlutz/go-full/internal/trainer/app/command"
	"github.com/cooperlutz/go-full/internal/trainer/app/event"
	"github.com/cooperlutz/go-full/internal/trainer/app/query"
	"github.com/cooperlutz/go-full/internal/trainer/domain/hour"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Events   Events
}

type Events struct {
	TrainerHourMadeAvailable   event.TrainerHourMadeAvailableHandler
	TrainerHourMadeUnavailable event.HourNoLongerAvailableHandler
}

type Commands struct {
	CancelTraining   command.CancelTrainingHandler
	ScheduleTraining command.ScheduleTrainingHandler

	MakeHoursAvailable   command.MakeHoursAvailableHandler
	MakeHoursUnavailable command.MakeHoursUnavailableHandler
}

type Queries struct {
	HourAvailability      query.HourAvailabilityHandler
	TrainerAvailableHours query.AvailableHoursHandler
}

func NewApplication(ctx context.Context, pgconn deebee.IDatabase, basePS *eeventdriven.BasePgsqlPubSubProcessor) Application {
	factoryConfig := hour.FactoryConfig{
		MaxWeeksInTheFutureToSet: 6,  //nolint:mnd // business rule
		MinUtcHour:               12, //nolint:mnd // 8am local time
		MaxUtcHour:               20, //nolint:mnd // 4pm local time
	}

	hourFactory, err := hour.NewFactory(factoryConfig)
	if err != nil {
		panic(err)
	}

	hourRepository := adapters.NewMemoryHourRepository(hourFactory)

	ps, err := adapters.NewHourPubSub(pgconn, nil)
	if err != nil {
		panic(err)
	}

	app := Application{
		Commands: Commands{
			CancelTraining:       command.NewCancelTrainingHandler(hourRepository),
			ScheduleTraining:     command.NewScheduleTrainingHandler(hourRepository),
			MakeHoursAvailable:   command.NewMakeHoursAvailableHandler(hourRepository),
			MakeHoursUnavailable: command.NewMakeHoursUnavailableHandler(hourRepository),
		},
		Queries: Queries{
			HourAvailability: query.NewHourAvailabilityHandler(hourRepository),
		},
		Events: Events{
			TrainerHourMadeAvailable:   event.NewTrainerHourMadeAvailableHandler(ps),
			TrainerHourMadeUnavailable: event.NewHourNoLongerAvailableHandler(ps),
		},
	}

	return app
}
