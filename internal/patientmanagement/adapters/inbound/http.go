package inbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/patientmanagement/app"
	"github.com/cooperlutz/go-full/internal/patientmanagement/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// HttpAdapter represents the HTTP server for the PatientManagement module.
type HttpAdapter struct {
	app app.Application
}

// NewHttpAdapter creates a new HttpAdapter instance with the provided PatientManagement application.
func NewHttpAdapter(application app.Application) HttpAdapter {
	return HttpAdapter{
		app: application,
	}
}

// StrictHandler returns a strict HTTP handler for the PatientManagement module.
func (h HttpAdapter) StrictHandler() ServerInterface {
	return NewStrictHandler(h, nil)
}

// (GET /v1/pets).
func (h HttpAdapter) FindAllPets(ctx context.Context, request FindAllPetsRequestObject) (FindAllPetsResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "pet.adapters.inbound.http.find_all_pets")
	defer span.End()

	pet, err := h.app.Queries.FindAllPets.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responsePets []Pet
	for _, e := range pet {
		responsePets = append(responsePets, queryPetToHttpPet(e))
	}

	return FindAllPets200JSONResponse(responsePets), nil
}

// (GET /v1/pet/{petId}).
func (h HttpAdapter) FindOnePet(ctx context.Context, request FindOnePetRequestObject) (FindOnePetResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_pet")
	defer span.End()

	pet, err := h.app.Queries.FindOnePet.Handle(ctx, query.FindOnePet{PetID: request.PetId})
	if err != nil {
		return nil, err
	}

	return FindOnePet200JSONResponse(queryPetToHttpPet(pet)), nil
}

func queryPetToHttpPet(e query.Pet) Pet {
	return Pet{
		//
		//PetId: GetPetId(),
		//
		//OwnerId: GetOwnerId(),
		//
		//Name: GetName(),
		//
		//Species: GetSpecies(),
		//
		//Breed: GetBreed(),
		//
		//DateOfBirth: GetDateOfBirth(),
		//
		//Gender: GetGender(),
		//
		//Weight: GetWeight(),
		//
		//MicrochipNumber: GetMicrochipNumber(),
		//
		//Status: GetStatus(),
		//
		// TODO
	}
}

// (GET /v1/medicalrecords).
func (h HttpAdapter) FindAllMedicalRecords(ctx context.Context, request FindAllMedicalRecordsRequestObject) (FindAllMedicalRecordsResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "medicalrecord.adapters.inbound.http.find_all_medicalrecords")
	defer span.End()

	medicalrecord, err := h.app.Queries.FindAllMedicalRecords.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseMedicalRecords []MedicalRecord
	for _, e := range medicalrecord {
		responseMedicalRecords = append(responseMedicalRecords, queryMedicalRecordToHttpMedicalRecord(e))
	}

	return FindAllMedicalRecords200JSONResponse(responseMedicalRecords), nil
}

// (GET /v1/medicalrecord/{medical_recordId}).
func (h HttpAdapter) FindOneMedicalRecord(ctx context.Context, request FindOneMedicalRecordRequestObject) (FindOneMedicalRecordResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_medical_record")
	defer span.End()

	medicalrecord, err := h.app.Queries.FindOneMedicalRecord.Handle(ctx, query.FindOneMedicalRecord{MedicalRecordID: request.MedicalRecordId})
	if err != nil {
		return nil, err
	}

	return FindOneMedicalRecord200JSONResponse(queryMedicalRecordToHttpMedicalRecord(medicalrecord)), nil
}

func queryMedicalRecordToHttpMedicalRecord(e query.MedicalRecord) MedicalRecord {
	return MedicalRecord{
		//
		//MedicalRecordId: GetMedicalRecordId(),
		//
		//PetId: GetPetId(),
		//
		//VeterinarianId: GetVeterinarianId(),
		//
		//VisitDate: GetVisitDate(),
		//
		//Diagnosis: GetDiagnosis(),
		//
		//Treatment: GetTreatment(),
		//
		//Notes: GetNotes(),
		//
		//FollowUpRequired: GetFollowUpRequired(),
		//
		// TODO
	}
}

// (GET /v1/vaccinationrecords).
func (h HttpAdapter) FindAllVaccinationRecords(ctx context.Context, request FindAllVaccinationRecordsRequestObject) (FindAllVaccinationRecordsResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "vaccinationrecord.adapters.inbound.http.find_all_vaccinationrecords")
	defer span.End()

	vaccinationrecord, err := h.app.Queries.FindAllVaccinationRecords.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseVaccinationRecords []VaccinationRecord
	for _, e := range vaccinationrecord {
		responseVaccinationRecords = append(responseVaccinationRecords, queryVaccinationRecordToHttpVaccinationRecord(e))
	}

	return FindAllVaccinationRecords200JSONResponse(responseVaccinationRecords), nil
}

// (GET /v1/vaccinationrecord/{vaccination_recordId}).
func (h HttpAdapter) FindOneVaccinationRecord(ctx context.Context, request FindOneVaccinationRecordRequestObject) (FindOneVaccinationRecordResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_vaccination_record")
	defer span.End()

	vaccinationrecord, err := h.app.Queries.FindOneVaccinationRecord.Handle(ctx, query.FindOneVaccinationRecord{VaccinationRecordID: request.VaccinationRecordId})
	if err != nil {
		return nil, err
	}

	return FindOneVaccinationRecord200JSONResponse(queryVaccinationRecordToHttpVaccinationRecord(vaccinationrecord)), nil
}

func queryVaccinationRecordToHttpVaccinationRecord(e query.VaccinationRecord) VaccinationRecord {
	return VaccinationRecord{
		//
		//VaccinationId: GetVaccinationId(),
		//
		//PetId: GetPetId(),
		//
		//VaccineName: GetVaccineName(),
		//
		//AdministeredDate: GetAdministeredDate(),
		//
		//ExpiryDate: GetExpiryDate(),
		//
		//AdministeredBy: GetAdministeredBy(),
		//
		// TODO
	}
}
