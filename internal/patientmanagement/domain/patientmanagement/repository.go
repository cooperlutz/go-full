package patientmanagement

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type PetRepository interface {
	AddPet(ctx context.Context, pet *Pet) error

	GetPet(ctx context.Context, id uuid.UUID) (*Pet, error)

	UpdatePet(
		ctx context.Context,
		petId uuid.UUID,
		updateFn func(e *Pet) (*Pet, error),
	) error
}

// MapToPet creates a Pet domain object from the given parameters.
// This should ONLY BE USED when reconstructing an Pet from its repository.
func MapToPet(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//petId string,
	//
	//ownerId string,
	//
	//name string,
	//
	//species string,
	//
	//breed *string,
	//
	//dateOfBirth *string,
	//
	//gender string,
	//
	//weight *float32,
	//
	//microchipNumber *string,
	//
	//status string,
	//
) (*Pet, error) {
	return &Pet{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
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
		// TODO
	}, nil
}

type MedicalRecordRepository interface {
	AddMedicalRecord(ctx context.Context, medicalrecord *MedicalRecord) error

	GetMedicalRecord(ctx context.Context, id uuid.UUID) (*MedicalRecord, error)

	UpdateMedicalRecord(
		ctx context.Context,
		medicalrecordId uuid.UUID,
		updateFn func(e *MedicalRecord) (*MedicalRecord, error),
	) error
}

// MapToMedicalRecord creates a MedicalRecord domain object from the given parameters.
// This should ONLY BE USED when reconstructing an MedicalRecord from its repository.
func MapToMedicalRecord(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//medicalRecordId string,
	//
	//petId string,
	//
	//veterinarianId string,
	//
	//visitDate string,
	//
	//diagnosis *string,
	//
	//treatment *string,
	//
	//notes *string,
	//
	//followUpRequired bool,
	//
) (*MedicalRecord, error) {
	return &MedicalRecord{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
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
		// TODO
	}, nil
}

type VaccinationRecordRepository interface {
	AddVaccinationRecord(ctx context.Context, vaccinationrecord *VaccinationRecord) error

	GetVaccinationRecord(ctx context.Context, id uuid.UUID) (*VaccinationRecord, error)

	UpdateVaccinationRecord(
		ctx context.Context,
		vaccinationrecordId uuid.UUID,
		updateFn func(e *VaccinationRecord) (*VaccinationRecord, error),
	) error
}

// MapToVaccinationRecord creates a VaccinationRecord domain object from the given parameters.
// This should ONLY BE USED when reconstructing an VaccinationRecord from its repository.
func MapToVaccinationRecord(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//vaccinationId string,
	//
	//petId string,
	//
	//vaccineName string,
	//
	//administeredDate string,
	//
	//expiryDate *string,
	//
	//administeredBy string,
	//
) (*VaccinationRecord, error) {
	return &VaccinationRecord{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
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
		// TODO
	}, nil
}
