package patientmanagement

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type MedicalRecord struct {
	*baseentitee.EntityMetadata
	//
	//medicalRecordId string
	//
	//petId string
	//
	//veterinarianId string
	//
	//visitDate string
	//
	//diagnosis *string
	//
	//treatment *string
	//
	//notes *string
	//
	//followUpRequired bool
	//
	// TODO
}

func NewMedicalRecord(
// medicalRecordId string,
//
// petId string,
//
// veterinarianId string,
//
// visitDate string,
//
// diagnosis *string,
//
// treatment *string,
//
// notes *string,
//
// followUpRequired bool,
) *MedicalRecord {
	return &MedicalRecord{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		//
		//medicalRecordId: medicalRecordId,
		//
		//petId: petId,
		//
		//veterinarianId: veterinarianId,
		//
		//visitDate: visitDate,
		//
		//diagnosis: diagnosis,
		//
		//treatment: treatment,
		//
		//notes: notes,
		//
		//followUpRequired: followUpRequired,
		//
	}
}
