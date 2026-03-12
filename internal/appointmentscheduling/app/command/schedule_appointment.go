package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/appointmentscheduling/domain/appointmentscheduling"
)

type ScheduleAppointment struct {
	//
	//PetId string,
	//
	//OwnerId string,
	//
	//VeterinarianId string,
	//
	//AppointmentType string,
	//
	//ScheduledDate string,
	//
	//ScheduledTime string,
	//
	//DurationMinutes int32,
	//
	//Notes *string,
	//
	//IsTelemedicine bool,
	//
	// TODO
}

type ScheduleAppointmentHandler struct {
	AppointmentRepo appointmentscheduling.AppointmentRepository

	TelemedicineSessionRepo appointmentscheduling.TelemedicineSessionRepository
}

func NewScheduleAppointmentHandler(
	appointmentRepo appointmentscheduling.AppointmentRepository,

	telemedicinesessionRepo appointmentscheduling.TelemedicineSessionRepository,
) ScheduleAppointmentHandler {
	return ScheduleAppointmentHandler{
		AppointmentRepo: appointmentRepo,

		TelemedicineSessionRepo: telemedicinesessionRepo,
	}
}

func (h ScheduleAppointmentHandler) Handle(ctx context.Context, cmd ScheduleAppointment) error {
	// ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.command.schedule_appointment.handle")
	// defer span.End()

	// TODO
	//err = h.AppointmentRepo.UpdateAppointment(ctx, uuid.MustParse(cmd.AppointmentId), func(a *appointmentscheduling.Appointment) (*appointmentscheduling.Appointment, error) {
	//
	//	 err := a.ScheduleAppointment(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return a, nil
	//})
	//if err != nil {
	//	return err
	//}

	// TODO
	//err = h.TelemedicineSessionRepo.UpdateTelemedicineSession(ctx, uuid.MustParse(cmd.TelemedicineSessionId), func(t *appointmentscheduling.TelemedicineSession) (*appointmentscheduling.TelemedicineSession, error) {
	//
	//	 err := t.ScheduleAppointment(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return t, nil
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}
