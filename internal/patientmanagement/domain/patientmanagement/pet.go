package patientmanagement

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type Pet struct {
	*baseentitee.EntityMetadata
	//
	//petId string
	//
	//ownerId string
	//
	//name string
	//
	//species string
	//
	//breed *string
	//
	//dateOfBirth *string
	//
	//gender string
	//
	//weight *float32
	//
	//microchipNumber *string
	//
	//status string
	//
	// TODO
}

func NewPet(
// petId string,
//
// ownerId string,
//
// name string,
//
// species string,
//
// breed *string,
//
// dateOfBirth *string,
//
// gender string,
//
// weight *float32,
//
// microchipNumber *string,
//
// status string,
) *Pet {
	return &Pet{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		//
		//petId: petId,
		//
		//ownerId: ownerId,
		//
		//name: name,
		//
		//species: species,
		//
		//breed: breed,
		//
		//dateOfBirth: dateOfBirth,
		//
		//gender: gender,
		//
		//weight: weight,
		//
		//microchipNumber: microchipNumber,
		//
		//status: status,
		//
	}
}
