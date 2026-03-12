package veterinarystaff

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type AvailabilitySchedule struct {
	*baseentitee.EntityMetadata
	//
	//scheduleId string
	//
	//staffId string
	//
	//dayOfWeek string
	//
	//startTime string
	//
	//endTime string
	//
	//isAvailable bool
	//
	// TODO
}

func NewAvailabilitySchedule(
// scheduleId string,
//
// staffId string,
//
// dayOfWeek string,
//
// startTime string,
//
// endTime string,
//
// isAvailable bool,
) *AvailabilitySchedule {
	return &AvailabilitySchedule{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		//
		//scheduleId: scheduleId,
		//
		//staffId: staffId,
		//
		//dayOfWeek: dayOfWeek,
		//
		//startTime: startTime,
		//
		//endTime: endTime,
		//
		//isAvailable: isAvailable,
		//
	}
}
