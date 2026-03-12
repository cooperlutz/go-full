package inbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/appointmentscheduling/app"
	"github.com/cooperlutz/go-full/internal/appointmentscheduling/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// HttpAdapter represents the HTTP server for the AppointmentScheduling module.
type HttpAdapter struct {
	app app.Application
}

// NewHttpAdapter creates a new HttpAdapter instance with the provided AppointmentScheduling application.
func NewHttpAdapter(application app.Application) HttpAdapter {
	return HttpAdapter{
		app: application,
	}
}

// StrictHandler returns a strict HTTP handler for the AppointmentScheduling module.
func (h HttpAdapter) StrictHandler() ServerInterface {
	return NewStrictHandler(h, nil)
}

// (GET /v1/appointments).
func (h HttpAdapter) FindAllAppointments(ctx context.Context, request FindAllAppointmentsRequestObject) (FindAllAppointmentsResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "appointment.adapters.inbound.http.find_all_appointments")
	defer span.End()

	appointment, err := h.app.Queries.FindAllAppointments.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseAppointments []Appointment
	for _, e := range appointment {
		responseAppointments = append(responseAppointments, queryAppointmentToHttpAppointment(e))
	}

	return FindAllAppointments200JSONResponse(responseAppointments), nil
}

// (GET /v1/appointment/{appointmentId}).
func (h HttpAdapter) FindOneAppointment(ctx context.Context, request FindOneAppointmentRequestObject) (FindOneAppointmentResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_appointment")
	defer span.End()

	appointment, err := h.app.Queries.FindOneAppointment.Handle(ctx, query.FindOneAppointment{AppointmentID: request.AppointmentId})
	if err != nil {
		return nil, err
	}

	return FindOneAppointment200JSONResponse(queryAppointmentToHttpAppointment(appointment)), nil
}

func queryAppointmentToHttpAppointment(e query.Appointment) Appointment {
	return Appointment{
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

// (GET /v1/telemedicinesessions).
func (h HttpAdapter) FindAllTelemedicineSessions(ctx context.Context, request FindAllTelemedicineSessionsRequestObject) (FindAllTelemedicineSessionsResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "telemedicinesession.adapters.inbound.http.find_all_telemedicinesessions")
	defer span.End()

	telemedicinesession, err := h.app.Queries.FindAllTelemedicineSessions.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseTelemedicineSessions []TelemedicineSession
	for _, e := range telemedicinesession {
		responseTelemedicineSessions = append(responseTelemedicineSessions, queryTelemedicineSessionToHttpTelemedicineSession(e))
	}

	return FindAllTelemedicineSessions200JSONResponse(responseTelemedicineSessions), nil
}

// (GET /v1/telemedicinesession/{telemedicine_sessionId}).
func (h HttpAdapter) FindOneTelemedicineSession(ctx context.Context, request FindOneTelemedicineSessionRequestObject) (FindOneTelemedicineSessionResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_telemedicine_session")
	defer span.End()

	telemedicinesession, err := h.app.Queries.FindOneTelemedicineSession.Handle(ctx, query.FindOneTelemedicineSession{TelemedicineSessionID: request.TelemedicineSessionId})
	if err != nil {
		return nil, err
	}

	return FindOneTelemedicineSession200JSONResponse(queryTelemedicineSessionToHttpTelemedicineSession(telemedicinesession)), nil
}

func queryTelemedicineSessionToHttpTelemedicineSession(e query.TelemedicineSession) TelemedicineSession {
	return TelemedicineSession{
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
