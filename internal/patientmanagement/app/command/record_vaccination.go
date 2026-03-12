package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/patientmanagement/domain/patientmanagement"
)

type RecordVaccination struct {
	//
	//PetId string,
	//
	//VaccineName string,
	//
	//AdministeredDate string,
	//
	//ExpiryDate *string,
	//
	//AdministeredBy string,
	//
	// TODO
}

type RecordVaccinationHandler struct {
	PetRepo patientmanagement.PetRepository

	MedicalRecordRepo patientmanagement.MedicalRecordRepository

	VaccinationRecordRepo patientmanagement.VaccinationRecordRepository
}

func NewRecordVaccinationHandler(
	petRepo patientmanagement.PetRepository,

	medicalrecordRepo patientmanagement.MedicalRecordRepository,

	vaccinationrecordRepo patientmanagement.VaccinationRecordRepository,
) RecordVaccinationHandler {
	return RecordVaccinationHandler{
		PetRepo: petRepo,

		MedicalRecordRepo: medicalrecordRepo,

		VaccinationRecordRepo: vaccinationrecordRepo,
	}
}

func (h RecordVaccinationHandler) Handle(ctx context.Context, cmd RecordVaccination) error {
	// ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.command.record_vaccination.handle")
	// defer span.End()

	// TODO
	//err = h.PetRepo.UpdatePet(ctx, uuid.MustParse(cmd.PetId), func(p *patientmanagement.Pet) (*patientmanagement.Pet, error) {
	//
	//	 err := p.RecordVaccination(
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
	//	 err := m.RecordVaccination(
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
	//	 err := v.RecordVaccination(
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
