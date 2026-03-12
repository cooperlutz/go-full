package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/appointmentscheduling/domain/appointmentscheduling"
)

type StartTelemedicineSession struct {
	//
	//AppointmentId string,
	//
	// TODO
}

type StartTelemedicineSessionHandler struct {
	AppointmentRepo appointmentscheduling.AppointmentRepository

	TelemedicineSessionRepo appointmentscheduling.TelemedicineSessionRepository
}

func NewStartTelemedicineSessionHandler(
	appointmentRepo appointmentscheduling.AppointmentRepository,

	telemedicinesessionRepo appointmentscheduling.TelemedicineSessionRepository,
) StartTelemedicineSessionHandler {
	return StartTelemedicineSessionHandler{
		AppointmentRepo: appointmentRepo,

		TelemedicineSessionRepo: telemedicinesessionRepo,
	}
}

func (h StartTelemedicineSessionHandler) Handle(ctx context.Context, cmd StartTelemedicineSession) error {
	// ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.command.start_telemedicine_session.handle")
	// defer span.End()

	// TODO
	//err = h.AppointmentRepo.UpdateAppointment(ctx, uuid.MustParse(cmd.AppointmentId), func(a *appointmentscheduling.Appointment) (*appointmentscheduling.Appointment, error) {
	//
	//	 err := a.StartTelemedicineSession(
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
	//	 err := t.StartTelemedicineSession(
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
