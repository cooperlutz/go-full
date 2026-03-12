package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/patientmanagement/domain/patientmanagement"
)

type DeactivatePet struct {
	//
	//PetId string,
	//
	//Reason string,
	//
	// TODO
}

type DeactivatePetHandler struct {
	PetRepo patientmanagement.PetRepository

	MedicalRecordRepo patientmanagement.MedicalRecordRepository

	VaccinationRecordRepo patientmanagement.VaccinationRecordRepository
}

func NewDeactivatePetHandler(
	petRepo patientmanagement.PetRepository,

	medicalrecordRepo patientmanagement.MedicalRecordRepository,

	vaccinationrecordRepo patientmanagement.VaccinationRecordRepository,
) DeactivatePetHandler {
	return DeactivatePetHandler{
		PetRepo: petRepo,

		MedicalRecordRepo: medicalrecordRepo,

		VaccinationRecordRepo: vaccinationrecordRepo,
	}
}

func (h DeactivatePetHandler) Handle(ctx context.Context, cmd DeactivatePet) error {
	// ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.command.deactivate_pet.handle")
	// defer span.End()

	// TODO
	//err = h.PetRepo.UpdatePet(ctx, uuid.MustParse(cmd.PetId), func(p *patientmanagement.Pet) (*patientmanagement.Pet, error) {
	//
	//	 err := p.DeactivatePet(
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
	//	 err := m.DeactivatePet(
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
	//	 err := v.DeactivatePet(
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
