package outbound

import (
	"github.com/cooperlutz/go-full/internal/appointmentscheduling/app/query"
	"github.com/cooperlutz/go-full/internal/appointmentscheduling/domain/appointmentscheduling"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

// toDomain maps the AppointmentAppointment to the domain entity.
func (e AppointmentschedulingAppointment) toDomain() (*appointmentscheduling.Appointment, error) {
	return appointmentscheduling.MapToAppointment(
		e.AppointmentID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.AppointmentId,
		//
		//e.PetId,
		//
		//e.OwnerId,
		//
		//e.VeterinarianId,
		//
		//e.AppointmentType,
		//
		//e.ScheduledDate,
		//
		//e.ScheduledTime,
		//
		//e.DurationMinutes,
		//
		//e.Status,
		//
		//e.Notes,
		//
		//e.IsTelemedicine,
		//
		// TODO
	)
}

// toQueryAppointment maps the appointmentAppointment to the query.Appointment.
func (e AppointmentschedulingAppointment) toQueryAppointment() (query.Appointment, error) {
	appointment, err := e.toDomain()
	if err != nil {
		return query.Appointment{}, err
	}

	return mapEntityAppointmentToQuery(appointment), nil
}

// appointmentAppointmentsToQuery maps a slice of AppointmentAppointment to a slice of query.Appointment entities.
func appointmentschedulingAppointmentsToQuery(appointments []AppointmentschedulingAppointment) ([]query.Appointment, error) {
	var domainAppointments []query.Appointment

	for _, appointment := range appointments {
		queryAppointment, err := appointment.toQueryAppointment()
		if err != nil {
			return nil, err
		}

		domainAppointments = append(domainAppointments, queryAppointment)
	}

	return domainAppointments, nil
}

// mapEntityAppointmentToDB maps a domain Appointment entity to the AppointmentAppointment database model.
func mapEntityAppointmentToDB(appointment *appointmentscheduling.Appointment) AppointmentschedulingAppointment {
	createdAt := appointment.GetCreatedAtTime()
	updatedAt := appointment.GetUpdatedAtTime()

	return AppointmentschedulingAppointment{
		AppointmentID: pgxutil.UUIDToPgtypeUUID(appointment.GetIdUUID()),
		CreatedAt:     pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:     pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:       appointment.IsDeleted(),
		DeletedAt:     pgxutil.TimeToTimestampz(appointment.GetDeletedAtTime()),
		//
		//AppointmentId: GetAppointmentId(),
		//
		//PetId: GetPetId(),
		//
		//OwnerId: GetOwnerId(),
		//
		//VeterinarianId: GetVeterinarianId(),
		//
		//AppointmentType: GetAppointmentType(),
		//
		//ScheduledDate: GetScheduledDate(),
		//
		//ScheduledTime: GetScheduledTime(),
		//
		//DurationMinutes: GetDurationMinutes(),
		//
		//Status: GetStatus(),
		//
		//Notes: GetNotes(),
		//
		//IsTelemedicine: GetIsTelemedicine(),
		//
		// TODO
	}
}

// mapEntityAppointmentToQuery maps a domain Appointment entity to a query.Appointment.
func mapEntityAppointmentToQuery(appointment *appointmentscheduling.Appointment) query.Appointment {
	return query.Appointment{
		// TODO
	}
}

// toDomain maps the TelemedicinesessionTelemedicineSession to the domain entity.
func (e AppointmentschedulingTelemedicineSession) toDomain() (*appointmentscheduling.TelemedicineSession, error) {
	return appointmentscheduling.MapToTelemedicineSession(
		e.TelemedicineSessionID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.SessionId,
		//
		//e.AppointmentId,
		//
		//e.SessionUrl,
		//
		//e.StartedAt,
		//
		//e.EndedAt,
		//
		//e.Status,
		//
		// TODO
	)
}

// toQueryTelemedicineSession maps the telemedicinesessionTelemedicineSession to the query.TelemedicineSession.
func (e AppointmentschedulingTelemedicineSession) toQueryTelemedicineSession() (query.TelemedicineSession, error) {
	telemedicinesession, err := e.toDomain()
	if err != nil {
		return query.TelemedicineSession{}, err
	}

	return mapEntityTelemedicineSessionToQuery(telemedicinesession), nil
}

// telemedicinesessionTelemedicineSessionsToQuery maps a slice of TelemedicineSessionTelemedicineSession to a slice of query.TelemedicineSession entities.
func appointmentschedulingTelemedicineSessionsToQuery(telemedicinesessions []AppointmentschedulingTelemedicineSession) ([]query.TelemedicineSession, error) {
	var domainTelemedicineSessions []query.TelemedicineSession

	for _, telemedicinesession := range telemedicinesessions {
		queryTelemedicineSession, err := telemedicinesession.toQueryTelemedicineSession()
		if err != nil {
			return nil, err
		}

		domainTelemedicineSessions = append(domainTelemedicineSessions, queryTelemedicineSession)
	}

	return domainTelemedicineSessions, nil
}

// mapEntityTelemedicineSessionToDB maps a domain TelemedicineSession entity to the TelemedicineSessionTelemedicineSession database model.
func mapEntityTelemedicineSessionToDB(telemedicinesession *appointmentscheduling.TelemedicineSession) AppointmentschedulingTelemedicineSession {
	createdAt := telemedicinesession.GetCreatedAtTime()
	updatedAt := telemedicinesession.GetUpdatedAtTime()

	return AppointmentschedulingTelemedicineSession{
		TelemedicineSessionID: pgxutil.UUIDToPgtypeUUID(telemedicinesession.GetIdUUID()),
		CreatedAt:             pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:             pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:               telemedicinesession.IsDeleted(),
		DeletedAt:             pgxutil.TimeToTimestampz(telemedicinesession.GetDeletedAtTime()),
		//
		//SessionId: GetSessionId(),
		//
		//AppointmentId: GetAppointmentId(),
		//
		//SessionUrl: GetSessionUrl(),
		//
		//StartedAt: GetStartedAt(),
		//
		//EndedAt: GetEndedAt(),
		//
		//Status: GetStatus(),
		//
		// TODO
	}
}

// mapEntityTelemedicineSessionToQuery maps a domain TelemedicineSession entity to a query.TelemedicineSession.
func mapEntityTelemedicineSessionToQuery(telemedicinesession *appointmentscheduling.TelemedicineSession) query.TelemedicineSession {
	return query.TelemedicineSession{
		// TODO
	}
}
