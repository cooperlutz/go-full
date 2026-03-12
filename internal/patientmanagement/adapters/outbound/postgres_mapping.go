package outbound

import (
	"github.com/cooperlutz/go-full/internal/patientmanagement/app/query"
	"github.com/cooperlutz/go-full/internal/patientmanagement/domain/patientmanagement"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

// toDomain maps the PetPet to the domain entity.
func (e PatientmanagementPet) toDomain() (*patientmanagement.Pet, error) {
	return patientmanagement.MapToPet(
		e.PetID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.PetId,
		//
		//e.OwnerId,
		//
		//e.Name,
		//
		//e.Species,
		//
		//e.Breed,
		//
		//e.DateOfBirth,
		//
		//e.Gender,
		//
		//e.Weight,
		//
		//e.MicrochipNumber,
		//
		//e.Status,
		//
		// TODO
	)
}

// toQueryPet maps the petPet to the query.Pet.
func (e PatientmanagementPet) toQueryPet() (query.Pet, error) {
	pet, err := e.toDomain()
	if err != nil {
		return query.Pet{}, err
	}

	return mapEntityPetToQuery(pet), nil
}

// petPetsToQuery maps a slice of PetPet to a slice of query.Pet entities.
func patientmanagementPetsToQuery(pets []PatientmanagementPet) ([]query.Pet, error) {
	var domainPets []query.Pet

	for _, pet := range pets {
		queryPet, err := pet.toQueryPet()
		if err != nil {
			return nil, err
		}

		domainPets = append(domainPets, queryPet)
	}

	return domainPets, nil
}

// mapEntityPetToDB maps a domain Pet entity to the PetPet database model.
func mapEntityPetToDB(pet *patientmanagement.Pet) PatientmanagementPet {
	createdAt := pet.GetCreatedAtTime()
	updatedAt := pet.GetUpdatedAtTime()

	return PatientmanagementPet{
		PetID:     pgxutil.UUIDToPgtypeUUID(pet.GetIdUUID()),
		CreatedAt: pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt: pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:   pet.IsDeleted(),
		DeletedAt: pgxutil.TimeToTimestampz(pet.GetDeletedAtTime()),
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

// mapEntityPetToQuery maps a domain Pet entity to a query.Pet.
func mapEntityPetToQuery(pet *patientmanagement.Pet) query.Pet {
	return query.Pet{
		// TODO
	}
}

// toDomain maps the MedicalrecordMedicalRecord to the domain entity.
func (e PatientmanagementMedicalRecord) toDomain() (*patientmanagement.MedicalRecord, error) {
	return patientmanagement.MapToMedicalRecord(
		e.MedicalRecordID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.MedicalRecordId,
		//
		//e.PetId,
		//
		//e.VeterinarianId,
		//
		//e.VisitDate,
		//
		//e.Diagnosis,
		//
		//e.Treatment,
		//
		//e.Notes,
		//
		//e.FollowUpRequired,
		//
		// TODO
	)
}

// toQueryMedicalRecord maps the medicalrecordMedicalRecord to the query.MedicalRecord.
func (e PatientmanagementMedicalRecord) toQueryMedicalRecord() (query.MedicalRecord, error) {
	medicalrecord, err := e.toDomain()
	if err != nil {
		return query.MedicalRecord{}, err
	}

	return mapEntityMedicalRecordToQuery(medicalrecord), nil
}

// medicalrecordMedicalRecordsToQuery maps a slice of MedicalRecordMedicalRecord to a slice of query.MedicalRecord entities.
func patientmanagementMedicalRecordsToQuery(medicalrecords []PatientmanagementMedicalRecord) ([]query.MedicalRecord, error) {
	var domainMedicalRecords []query.MedicalRecord

	for _, medicalrecord := range medicalrecords {
		queryMedicalRecord, err := medicalrecord.toQueryMedicalRecord()
		if err != nil {
			return nil, err
		}

		domainMedicalRecords = append(domainMedicalRecords, queryMedicalRecord)
	}

	return domainMedicalRecords, nil
}

// mapEntityMedicalRecordToDB maps a domain MedicalRecord entity to the MedicalRecordMedicalRecord database model.
func mapEntityMedicalRecordToDB(medicalrecord *patientmanagement.MedicalRecord) PatientmanagementMedicalRecord {
	createdAt := medicalrecord.GetCreatedAtTime()
	updatedAt := medicalrecord.GetUpdatedAtTime()

	return PatientmanagementMedicalRecord{
		MedicalRecordID: pgxutil.UUIDToPgtypeUUID(medicalrecord.GetIdUUID()),
		CreatedAt:       pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:       pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:         medicalrecord.IsDeleted(),
		DeletedAt:       pgxutil.TimeToTimestampz(medicalrecord.GetDeletedAtTime()),
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

// mapEntityMedicalRecordToQuery maps a domain MedicalRecord entity to a query.MedicalRecord.
func mapEntityMedicalRecordToQuery(medicalrecord *patientmanagement.MedicalRecord) query.MedicalRecord {
	return query.MedicalRecord{
		// TODO
	}
}

// toDomain maps the VaccinationrecordVaccinationRecord to the domain entity.
func (e PatientmanagementVaccinationRecord) toDomain() (*patientmanagement.VaccinationRecord, error) {
	return patientmanagement.MapToVaccinationRecord(
		e.VaccinationRecordID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.VaccinationId,
		//
		//e.PetId,
		//
		//e.VaccineName,
		//
		//e.AdministeredDate,
		//
		//e.ExpiryDate,
		//
		//e.AdministeredBy,
		//
		// TODO
	)
}

// toQueryVaccinationRecord maps the vaccinationrecordVaccinationRecord to the query.VaccinationRecord.
func (e PatientmanagementVaccinationRecord) toQueryVaccinationRecord() (query.VaccinationRecord, error) {
	vaccinationrecord, err := e.toDomain()
	if err != nil {
		return query.VaccinationRecord{}, err
	}

	return mapEntityVaccinationRecordToQuery(vaccinationrecord), nil
}

// vaccinationrecordVaccinationRecordsToQuery maps a slice of VaccinationRecordVaccinationRecord to a slice of query.VaccinationRecord entities.
func patientmanagementVaccinationRecordsToQuery(vaccinationrecords []PatientmanagementVaccinationRecord) ([]query.VaccinationRecord, error) {
	var domainVaccinationRecords []query.VaccinationRecord

	for _, vaccinationrecord := range vaccinationrecords {
		queryVaccinationRecord, err := vaccinationrecord.toQueryVaccinationRecord()
		if err != nil {
			return nil, err
		}

		domainVaccinationRecords = append(domainVaccinationRecords, queryVaccinationRecord)
	}

	return domainVaccinationRecords, nil
}

// mapEntityVaccinationRecordToDB maps a domain VaccinationRecord entity to the VaccinationRecordVaccinationRecord database model.
func mapEntityVaccinationRecordToDB(vaccinationrecord *patientmanagement.VaccinationRecord) PatientmanagementVaccinationRecord {
	createdAt := vaccinationrecord.GetCreatedAtTime()
	updatedAt := vaccinationrecord.GetUpdatedAtTime()

	return PatientmanagementVaccinationRecord{
		VaccinationRecordID: pgxutil.UUIDToPgtypeUUID(vaccinationrecord.GetIdUUID()),
		CreatedAt:           pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:           pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:             vaccinationrecord.IsDeleted(),
		DeletedAt:           pgxutil.TimeToTimestampz(vaccinationrecord.GetDeletedAtTime()),
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

// mapEntityVaccinationRecordToQuery maps a domain VaccinationRecord entity to a query.VaccinationRecord.
func mapEntityVaccinationRecordToQuery(vaccinationrecord *patientmanagement.VaccinationRecord) query.VaccinationRecord {
	return query.VaccinationRecord{
		// TODO
	}
}
