package appointmentscheduling

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type AppointmentRepository interface {
	AddAppointment(ctx context.Context, appointment *Appointment) error

	GetAppointment(ctx context.Context, id uuid.UUID) (*Appointment, error)

	UpdateAppointment(
		ctx context.Context,
		appointmentId uuid.UUID,
		updateFn func(e *Appointment) (*Appointment, error),
	) error
}

// MapToAppointment creates a Appointment domain object from the given parameters.
// This should ONLY BE USED when reconstructing an Appointment from its repository.
func MapToAppointment(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//appointmentId string,
	//
	//petId string,
	//
	//ownerId string,
	//
	//veterinarianId string,
	//
	//appointmentType string,
	//
	//scheduledDate string,
	//
	//scheduledTime string,
	//
	//durationMinutes int32,
	//
	//status string,
	//
	//notes *string,
	//
	//isTelemedicine bool,
	//
) (*Appointment, error) {
	return &Appointment{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		//
		//appointmentId: appointmentId,
		//
		//petId: petId,
		//
		//ownerId: ownerId,
		//
		//veterinarianId: veterinarianId,
		//
		//appointmentType: appointmentType,
		//
		//scheduledDate: scheduledDate,
		//
		//scheduledTime: scheduledTime,
		//
		//durationMinutes: durationMinutes,
		//
		//status: status,
		//
		//notes: notes,
		//
		//isTelemedicine: isTelemedicine,
		//
		// TODO
	}, nil
}

type TelemedicineSessionRepository interface {
	AddTelemedicineSession(ctx context.Context, telemedicinesession *TelemedicineSession) error

	GetTelemedicineSession(ctx context.Context, id uuid.UUID) (*TelemedicineSession, error)

	UpdateTelemedicineSession(
		ctx context.Context,
		telemedicinesessionId uuid.UUID,
		updateFn func(e *TelemedicineSession) (*TelemedicineSession, error),
	) error
}

// MapToTelemedicineSession creates a TelemedicineSession domain object from the given parameters.
// This should ONLY BE USED when reconstructing an TelemedicineSession from its repository.
func MapToTelemedicineSession(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//sessionId string,
	//
	//appointmentId string,
	//
	//sessionUrl string,
	//
	//startedAt *string,
	//
	//endedAt *string,
	//
	//status string,
	//
) (*TelemedicineSession, error) {
	return &TelemedicineSession{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		//
		//sessionId: sessionId,
		//
		//appointmentId: appointmentId,
		//
		//sessionUrl: sessionUrl,
		//
		//startedAt: startedAt,
		//
		//endedAt: endedAt,
		//
		//status: status,
		//
		// TODO
	}, nil
}
