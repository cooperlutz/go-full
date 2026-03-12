package veterinarystaff

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type StaffMember struct {
	*baseentitee.EntityMetadata
	//
	//staffId string
	//
	//firstName string
	//
	//lastName string
	//
	//email string
	//
	//role string
	//
	//status string
	//
	// TODO
}

func NewStaffMember(
// staffId string,
//
// firstName string,
//
// lastName string,
//
// email string,
//
// role string,
//
// status string,
) *StaffMember {
	return &StaffMember{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		//
		//staffId: staffId,
		//
		//firstName: firstName,
		//
		//lastName: lastName,
		//
		//email: email,
		//
		//role: role,
		//
		//status: status,
		//
	}
}
