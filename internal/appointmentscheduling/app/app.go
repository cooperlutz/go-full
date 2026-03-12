package app

import (
	"github.com/cooperlutz/go-full/internal/appointmentscheduling/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/appointmentscheduling/app/command"
	"github.com/cooperlutz/go-full/internal/appointmentscheduling/app/event"
	"github.com/cooperlutz/go-full/internal/appointmentscheduling/app/query"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Events   Events
}

type Commands struct {
	ScheduleAppointment command.ScheduleAppointmentHandler

	RescheduleAppointment command.RescheduleAppointmentHandler

	CancelAppointment command.CancelAppointmentHandler

	ConfirmAppointment command.ConfirmAppointmentHandler

	CompleteAppointment command.CompleteAppointmentHandler

	StartTelemedicineSession command.StartTelemedicineSessionHandler

	EndTelemedicineSession command.EndTelemedicineSessionHandler
}

type Queries struct {
	FindAllAppointments query.FindAllAppointmentsHandler
	FindOneAppointment  query.FindOneAppointmentHandler

	FindAllTelemedicineSessions query.FindAllTelemedicineSessionsHandler
	FindOneTelemedicineSession  query.FindOneTelemedicineSessionHandler
}

type Events struct {
	AppointmentScheduled event.AppointmentScheduledHandler

	AppointmentRescheduled event.AppointmentRescheduledHandler

	AppointmentCancelled event.AppointmentCancelledHandler

	AppointmentConfirmed event.AppointmentConfirmedHandler

	AppointmentCompleted event.AppointmentCompletedHandler

	AppointmentReminderScheduled event.AppointmentReminderScheduledHandler

	TelemedicineSessionStarted event.TelemedicineSessionStartedHandler

	TelemedicineSessionEnded event.TelemedicineSessionEndedHandler
}

// NewApplication initializes the AppointmentScheduling application with its dependencies.
func NewApplication( //nolint:funlen // it's fine
	pgconn deebee.IDatabase,
	pubSub eeventdriven.IPubSubEventProcessor,
) (Application, error) {
	appointmentRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	telemedicinesessionRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	app := Application{
		Commands: Commands{
			ScheduleAppointment: command.NewScheduleAppointmentHandler(

				appointmentRepository,

				telemedicinesessionRepository,
			),
			RescheduleAppointment: command.NewRescheduleAppointmentHandler(

				appointmentRepository,

				telemedicinesessionRepository,
			),
			CancelAppointment: command.NewCancelAppointmentHandler(

				appointmentRepository,

				telemedicinesessionRepository,
			),
			ConfirmAppointment: command.NewConfirmAppointmentHandler(

				appointmentRepository,

				telemedicinesessionRepository,
			),
			CompleteAppointment: command.NewCompleteAppointmentHandler(

				appointmentRepository,

				telemedicinesessionRepository,
			),
			StartTelemedicineSession: command.NewStartTelemedicineSessionHandler(

				appointmentRepository,

				telemedicinesessionRepository,
			),
			EndTelemedicineSession: command.NewEndTelemedicineSessionHandler(

				appointmentRepository,

				telemedicinesessionRepository,
			),
		},
		Queries: Queries{
			FindAllAppointments: query.NewFindAllAppointmentsHandler(
				appointmentRepository,
			),
			FindOneAppointment: query.NewFindOneAppointmentHandler(
				appointmentRepository,
			),

			FindAllTelemedicineSessions: query.NewFindAllTelemedicineSessionsHandler(
				telemedicinesessionRepository,
			),
			FindOneTelemedicineSession: query.NewFindOneTelemedicineSessionHandler(
				telemedicinesessionRepository,
			),
		},
		Events: Events{
			AppointmentScheduled: event.NewAppointmentScheduledHandler(
				pubSub,
			),

			AppointmentRescheduled: event.NewAppointmentRescheduledHandler(
				pubSub,
			),

			AppointmentCancelled: event.NewAppointmentCancelledHandler(
				pubSub,
			),

			AppointmentConfirmed: event.NewAppointmentConfirmedHandler(
				pubSub,
			),

			AppointmentCompleted: event.NewAppointmentCompletedHandler(
				pubSub,
			),

			AppointmentReminderScheduled: event.NewAppointmentReminderScheduledHandler(
				pubSub,
			),

			TelemedicineSessionStarted: event.NewTelemedicineSessionStartedHandler(
				pubSub,
			),

			TelemedicineSessionEnded: event.NewTelemedicineSessionEndedHandler(
				pubSub,
			),
		},
	}

	return app, nil
}
