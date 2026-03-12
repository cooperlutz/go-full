package appointmentscheduling

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type Appointment struct {
	*baseentitee.EntityMetadata
	//
	//appointmentId string
	//
	//petId string
	//
	//ownerId string
	//
	//veterinarianId string
	//
	//appointmentType string
	//
	//scheduledDate string
	//
	//scheduledTime string
	//
	//durationMinutes int32
	//
	//status string
	//
	//notes *string
	//
	//isTelemedicine bool
	//
	// TODO
}

func NewAppointment(
// appointmentId string,
//
// petId string,
//
// ownerId string,
//
// veterinarianId string,
//
// appointmentType string,
//
// scheduledDate string,
//
// scheduledTime string,
//
// durationMinutes int32,
//
// status string,
//
// notes *string,
//
// isTelemedicine bool,
) *Appointment {
	return &Appointment{
		EntityMetadata: baseentitee.NewEntityMetadata(),
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
	}
}
