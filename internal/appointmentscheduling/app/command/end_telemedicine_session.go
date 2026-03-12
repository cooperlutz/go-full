package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/appointmentscheduling/domain/appointmentscheduling"
)

type EndTelemedicineSession struct {
	//
	//SessionId string,
	//
	// TODO
}

type EndTelemedicineSessionHandler struct {
	AppointmentRepo appointmentscheduling.AppointmentRepository

	TelemedicineSessionRepo appointmentscheduling.TelemedicineSessionRepository
}

func NewEndTelemedicineSessionHandler(
	appointmentRepo appointmentscheduling.AppointmentRepository,

	telemedicinesessionRepo appointmentscheduling.TelemedicineSessionRepository,
) EndTelemedicineSessionHandler {
	return EndTelemedicineSessionHandler{
		AppointmentRepo: appointmentRepo,

		TelemedicineSessionRepo: telemedicinesessionRepo,
	}
}

func (h EndTelemedicineSessionHandler) Handle(ctx context.Context, cmd EndTelemedicineSession) error {
	// ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.command.end_telemedicine_session.handle")
	// defer span.End()

	// TODO
	//err = h.AppointmentRepo.UpdateAppointment(ctx, uuid.MustParse(cmd.AppointmentId), func(a *appointmentscheduling.Appointment) (*appointmentscheduling.Appointment, error) {
	//
	//	 err := a.EndTelemedicineSession(
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
	//	 err := t.EndTelemedicineSession(
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
