package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/patientmanagement/domain/patientmanagement"
)

type RegisterPet struct {
	//
	//OwnerId string,
	//
	//Name string,
	//
	//Species string,
	//
	//Breed *string,
	//
	//DateOfBirth *string,
	//
	//Gender string,
	//
	//Weight *float32,
	//
	//MicrochipNumber *string,
	//
	// TODO
}

type RegisterPetHandler struct {
	PetRepo patientmanagement.PetRepository

	MedicalRecordRepo patientmanagement.MedicalRecordRepository

	VaccinationRecordRepo patientmanagement.VaccinationRecordRepository
}

func NewRegisterPetHandler(
	petRepo patientmanagement.PetRepository,

	medicalrecordRepo patientmanagement.MedicalRecordRepository,

	vaccinationrecordRepo patientmanagement.VaccinationRecordRepository,
) RegisterPetHandler {
	return RegisterPetHandler{
		PetRepo: petRepo,

		MedicalRecordRepo: medicalrecordRepo,

		VaccinationRecordRepo: vaccinationrecordRepo,
	}
}

func (h RegisterPetHandler) Handle(ctx context.Context, cmd RegisterPet) error {
	// ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.command.register_pet.handle")
	// defer span.End()

	// TODO
	//err = h.PetRepo.UpdatePet(ctx, uuid.MustParse(cmd.PetId), func(p *patientmanagement.Pet) (*patientmanagement.Pet, error) {
	//
	//	 err := p.RegisterPet(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return p, nil
	//})
	//if err != nil {
	//	return err
	//}

	// TODO
	//err = h.MedicalRecordRepo.UpdateMedicalRecord(ctx, uuid.MustParse(cmd.MedicalRecordId), func(m *patientmanagement.MedicalRecord) (*patientmanagement.MedicalRecord, error) {
	//
	//	 err := m.RegisterPet(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return m, nil
	//})
	//if err != nil {
	//	return err
	//}

	// TODO
	//err = h.VaccinationRecordRepo.UpdateVaccinationRecord(ctx, uuid.MustParse(cmd.VaccinationRecordId), func(v *patientmanagement.VaccinationRecord) (*patientmanagement.VaccinationRecord, error) {
	//
	//	 err := v.RegisterPet(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return v, nil
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}
