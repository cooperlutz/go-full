package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/appointmentscheduling/domain/appointmentscheduling"
)

type RescheduleAppointment struct {
	//
	//AppointmentId string,
	//
	//NewScheduledDate string,
	//
	//NewScheduledTime string,
	//
	//Reason *string,
	//
	// TODO
}

type RescheduleAppointmentHandler struct {
	AppointmentRepo appointmentscheduling.AppointmentRepository

	TelemedicineSessionRepo appointmentscheduling.TelemedicineSessionRepository
}

func NewRescheduleAppointmentHandler(
	appointmentRepo appointmentscheduling.AppointmentRepository,

	telemedicinesessionRepo appointmentscheduling.TelemedicineSessionRepository,
) RescheduleAppointmentHandler {
	return RescheduleAppointmentHandler{
		AppointmentRepo: appointmentRepo,

		TelemedicineSessionRepo: telemedicinesessionRepo,
	}
}

func (h RescheduleAppointmentHandler) Handle(ctx context.Context, cmd RescheduleAppointment) error {
	// ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.command.reschedule_appointment.handle")
	// defer span.End()

	// TODO
	//err = h.AppointmentRepo.UpdateAppointment(ctx, uuid.MustParse(cmd.AppointmentId), func(a *appointmentscheduling.Appointment) (*appointmentscheduling.Appointment, error) {
	//
	//	 err := a.RescheduleAppointment(
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
	//	 err := t.RescheduleAppointment(
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
