package patientmanagement

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type VaccinationRecord struct {
	*baseentitee.EntityMetadata
	//
	//vaccinationId string
	//
	//petId string
	//
	//vaccineName string
	//
	//administeredDate string
	//
	//expiryDate *string
	//
	//administeredBy string
	//
	// TODO
}

func NewVaccinationRecord(
// vaccinationId string,
//
// petId string,
//
// vaccineName string,
//
// administeredDate string,
//
// expiryDate *string,
//
// administeredBy string,
) *VaccinationRecord {
	return &VaccinationRecord{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		//
		//vaccinationId: vaccinationId,
		//
		//petId: petId,
		//
		//vaccineName: vaccineName,
		//
		//administeredDate: administeredDate,
		//
		//expiryDate: expiryDate,
		//
		//administeredBy: administeredBy,
		//
	}
}
