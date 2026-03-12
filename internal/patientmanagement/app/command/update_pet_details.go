package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/patientmanagement/domain/patientmanagement"
)

type UpdatePetDetails struct {
	//
	//PetId string,
	//
	//Name *string,
	//
	//Weight *float32,
	//
	//MicrochipNumber *string,
	//
	// TODO
}

type UpdatePetDetailsHandler struct {
	PetRepo patientmanagement.PetRepository

	MedicalRecordRepo patientmanagement.MedicalRecordRepository

	VaccinationRecordRepo patientmanagement.VaccinationRecordRepository
}

func NewUpdatePetDetailsHandler(
	petRepo patientmanagement.PetRepository,

	medicalrecordRepo patientmanagement.MedicalRecordRepository,

	vaccinationrecordRepo patientmanagement.VaccinationRecordRepository,
) UpdatePetDetailsHandler {
	return UpdatePetDetailsHandler{
		PetRepo: petRepo,

		MedicalRecordRepo: medicalrecordRepo,

		VaccinationRecordRepo: vaccinationrecordRepo,
	}
}

func (h UpdatePetDetailsHandler) Handle(ctx context.Context, cmd UpdatePetDetails) error {
	// ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.command.update_pet_details.handle")
	// defer span.End()

	// TODO
	//err = h.PetRepo.UpdatePet(ctx, uuid.MustParse(cmd.PetId), func(p *patientmanagement.Pet) (*patientmanagement.Pet, error) {
	//
	//	 err := p.UpdatePetDetails(
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
	//	 err := m.UpdatePetDetails(
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
	//	 err := v.UpdatePetDetails(
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
