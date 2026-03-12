package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/appointmentscheduling/domain/appointmentscheduling"
)

type CancelAppointment struct {
	//
	//AppointmentId string,
	//
	//Reason string,
	//
	//CancelledBy string,
	//
	// TODO
}

type CancelAppointmentHandler struct {
	AppointmentRepo appointmentscheduling.AppointmentRepository

	TelemedicineSessionRepo appointmentscheduling.TelemedicineSessionRepository
}

func NewCancelAppointmentHandler(
	appointmentRepo appointmentscheduling.AppointmentRepository,

	telemedicinesessionRepo appointmentscheduling.TelemedicineSessionRepository,
) CancelAppointmentHandler {
	return CancelAppointmentHandler{
		AppointmentRepo: appointmentRepo,

		TelemedicineSessionRepo: telemedicinesessionRepo,
	}
}

func (h CancelAppointmentHandler) Handle(ctx context.Context, cmd CancelAppointment) error {
	// ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.command.cancel_appointment.handle")
	// defer span.End()

	// TODO
	//err = h.AppointmentRepo.UpdateAppointment(ctx, uuid.MustParse(cmd.AppointmentId), func(a *appointmentscheduling.Appointment) (*appointmentscheduling.Appointment, error) {
	//
	//	 err := a.CancelAppointment(
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
	//	 err := t.CancelAppointment(
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
