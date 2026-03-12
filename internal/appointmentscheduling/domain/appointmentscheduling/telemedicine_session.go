package appointmentscheduling

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type TelemedicineSession struct {
	*baseentitee.EntityMetadata
	//
	//sessionId string
	//
	//appointmentId string
	//
	//sessionUrl string
	//
	//startedAt *string
	//
	//endedAt *string
	//
	//status string
	//
	// TODO
}

func NewTelemedicineSession(
// sessionId string,
//
// appointmentId string,
//
// sessionUrl string,
//
// startedAt *string,
//
// endedAt *string,
//
// status string,
) *TelemedicineSession {
	return &TelemedicineSession{
		EntityMetadata: baseentitee.NewEntityMetadata(),
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
	}
}
