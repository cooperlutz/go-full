package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/patientmanagement/domain/patientmanagement"
)

type AddMedicalRecord struct {
	//
	//PetId string,
	//
	//VeterinarianId string,
	//
	//VisitDate string,
	//
	//Diagnosis *string,
	//
	//Treatment *string,
	//
	//Notes *string,
	//
	//FollowUpRequired bool,
	//
	// TODO
}

type AddMedicalRecordHandler struct {
	PetRepo patientmanagement.PetRepository

	MedicalRecordRepo patientmanagement.MedicalRecordRepository

	VaccinationRecordRepo patientmanagement.VaccinationRecordRepository
}

func NewAddMedicalRecordHandler(
	petRepo patientmanagement.PetRepository,

	medicalrecordRepo patientmanagement.MedicalRecordRepository,

	vaccinationrecordRepo patientmanagement.VaccinationRecordRepository,
) AddMedicalRecordHandler {
	return AddMedicalRecordHandler{
		PetRepo: petRepo,

		MedicalRecordRepo: medicalrecordRepo,

		VaccinationRecordRepo: vaccinationrecordRepo,
	}
}

func (h AddMedicalRecordHandler) Handle(ctx context.Context, cmd AddMedicalRecord) error {
	// ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.command.add_medical_record.handle")
	// defer span.End()

	// TODO
	//err = h.PetRepo.UpdatePet(ctx, uuid.MustParse(cmd.PetId), func(p *patientmanagement.Pet) (*patientmanagement.Pet, error) {
	//
	//	 err := p.AddMedicalRecord(
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
	//	 err := m.AddMedicalRecord(
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
	//	 err := v.AddMedicalRecord(
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
