package veterinarystaff

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type Veterinarian struct {
	*baseentitee.EntityMetadata
	//
	//veterinarianId string
	//
	//firstName string
	//
	//lastName string
	//
	//email string
	//
	//phoneNumber string
	//
	//licenseNumber string
	//
	//specializations *string
	//
	//status string
	//
	// TODO
}

func NewVeterinarian(
// veterinarianId string,
//
// firstName string,
//
// lastName string,
//
// email string,
//
// phoneNumber string,
//
// licenseNumber string,
//
// specializations *string,
//
// status string,
) *Veterinarian {
	return &Veterinarian{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		//
		//veterinarianId: veterinarianId,
		//
		//firstName: firstName,
		//
		//lastName: lastName,
		//
		//email: email,
		//
		//phoneNumber: phoneNumber,
		//
		//licenseNumber: licenseNumber,
		//
		//specializations: specializations,
		//
		//status: status,
		//
	}
}
