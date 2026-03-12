//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package outbound

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/cooperlutz/go-full/internal/appointmentscheduling/app/query"
	"github.com/cooperlutz/go-full/internal/appointmentscheduling/domain/appointmentscheduling"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// PostgresAdapter implements the examination repository using Postgres as the data store.
type PostgresAdapter struct {
	Handler IQuerierAppointmentScheduling
}

// NewPostgresAdapter creates a new instance of PostgresAdapter.
func NewPostgresAdapter(db deebee.IDatabase) PostgresAdapter {
	return PostgresAdapter{
		Handler: NewQueriesWrapper(db),
	}
}

func (p PostgresAdapter) FindAllAppointments(ctx context.Context) ([]query.Appointment, error) {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.adapters.outbound.postgres.find_all_appointment")
	defer span.End()

	appointments, err := p.Handler.FindAllAppointments(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return appointmentschedulingAppointmentsToQuery(appointments)
}

func (p PostgresAdapter) FindOneAppointment(ctx context.Context, id uuid.UUID) (query.Appointment, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_appointment")
	defer span.End()

	appointment, err := p.GetAppointment(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.Appointment{}, err
	}

	return mapEntityAppointmentToQuery(appointment), nil
}

// AddAppointment adds a new exam to the database.
func (p PostgresAdapter) AddAppointment(ctx context.Context, appointment *appointmentscheduling.Appointment) error {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.adapters.outbound.postgres.add_appointment")
	defer span.End()

	dbAppointment := mapEntityAppointmentToDB(appointment)

	err := p.Handler.AddAppointment(ctx, AddAppointmentParams(dbAppointment))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetAppointment(ctx context.Context, id uuid.UUID) (*appointmentscheduling.Appointment, error) {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.adapters.outbound.postgres.get_appointment")
	defer span.End()

	appointment, err := p.Handler.GetAppointment(
		ctx,
		GetAppointmentParams{AppointmentID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return appointment.toDomain()
}

func (p PostgresAdapter) UpdateAppointment(
	ctx context.Context,
	appointmentId uuid.UUID,
	updateFn func(e *appointmentscheduling.Appointment) (*appointmentscheduling.Appointment, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.adapters.outbound.postgres.update_appointment")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	appointment, err := p.GetAppointment(ctx, appointmentId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedAppointment, err := updateFn(appointment)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbAppointment := mapEntityAppointmentToDB(updatedAppointment)

	err = p.Handler.UpdateAppointment(ctx, UpdateAppointmentParams(dbAppointment))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) FindAllTelemedicineSessions(ctx context.Context) ([]query.TelemedicineSession, error) {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.adapters.outbound.postgres.find_all_telemedicine_session")
	defer span.End()

	telemedicinesessions, err := p.Handler.FindAllTelemedicineSessions(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return appointmentschedulingTelemedicineSessionsToQuery(telemedicinesessions)
}

func (p PostgresAdapter) FindOneTelemedicineSession(ctx context.Context, id uuid.UUID) (query.TelemedicineSession, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_telemedicine_session")
	defer span.End()

	telemedicinesession, err := p.GetTelemedicineSession(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.TelemedicineSession{}, err
	}

	return mapEntityTelemedicineSessionToQuery(telemedicinesession), nil
}

// AddTelemedicineSession adds a new exam to the database.
func (p PostgresAdapter) AddTelemedicineSession(ctx context.Context, telemedicinesession *appointmentscheduling.TelemedicineSession) error {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.adapters.outbound.postgres.add_telemedicine_session")
	defer span.End()

	dbTelemedicineSession := mapEntityTelemedicineSessionToDB(telemedicinesession)

	err := p.Handler.AddTelemedicineSession(ctx, AddTelemedicineSessionParams(dbTelemedicineSession))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetTelemedicineSession(ctx context.Context, id uuid.UUID) (*appointmentscheduling.TelemedicineSession, error) {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.adapters.outbound.postgres.get_telemedicine_session")
	defer span.End()

	telemedicinesession, err := p.Handler.GetTelemedicineSession(
		ctx,
		GetTelemedicineSessionParams{TelemedicineSessionID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return telemedicinesession.toDomain()
}

func (p PostgresAdapter) UpdateTelemedicineSession(
	ctx context.Context,
	telemedicinesessionId uuid.UUID,
	updateFn func(e *appointmentscheduling.TelemedicineSession) (*appointmentscheduling.TelemedicineSession, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.adapters.outbound.postgres.update_telemedicinesession")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	telemedicinesession, err := p.GetTelemedicineSession(ctx, telemedicinesessionId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedTelemedicineSession, err := updateFn(telemedicinesession)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbTelemedicineSession := mapEntityTelemedicineSessionToDB(updatedTelemedicineSession)

	err = p.Handler.UpdateTelemedicineSession(ctx, UpdateTelemedicineSessionParams(dbTelemedicineSession))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

// finishTransaction commits or rolls back the transaction based on the error state.
func (p PostgresAdapter) finishTransaction(ctx context.Context, err error, tx pgx.Tx) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			telemetree.RecordError(ctx, rollbackErr, "failed to rollback tx")

			return rollbackErr
		}

		return err
	} else {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			telemetree.RecordError(ctx, commitErr, "failed to commit tx")

			return commitErr
		}

		return nil
	}
}
