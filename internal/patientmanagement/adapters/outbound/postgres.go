//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package outbound

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/cooperlutz/go-full/internal/patientmanagement/app/query"
	"github.com/cooperlutz/go-full/internal/patientmanagement/domain/patientmanagement"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// PostgresAdapter implements the examination repository using Postgres as the data store.
type PostgresAdapter struct {
	Handler IQuerierPatientManagement
}

// NewPostgresAdapter creates a new instance of PostgresAdapter.
func NewPostgresAdapter(db deebee.IDatabase) PostgresAdapter {
	return PostgresAdapter{
		Handler: NewQueriesWrapper(db),
	}
}

func (p PostgresAdapter) FindAllPets(ctx context.Context) ([]query.Pet, error) {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.adapters.outbound.postgres.find_all_pet")
	defer span.End()

	pets, err := p.Handler.FindAllPets(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return patientmanagementPetsToQuery(pets)
}

func (p PostgresAdapter) FindOnePet(ctx context.Context, id uuid.UUID) (query.Pet, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_pet")
	defer span.End()

	pet, err := p.GetPet(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.Pet{}, err
	}

	return mapEntityPetToQuery(pet), nil
}

// AddPet adds a new exam to the database.
func (p PostgresAdapter) AddPet(ctx context.Context, pet *patientmanagement.Pet) error {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.adapters.outbound.postgres.add_pet")
	defer span.End()

	dbPet := mapEntityPetToDB(pet)

	err := p.Handler.AddPet(ctx, AddPetParams(dbPet))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetPet(ctx context.Context, id uuid.UUID) (*patientmanagement.Pet, error) {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.adapters.outbound.postgres.get_pet")
	defer span.End()

	pet, err := p.Handler.GetPet(
		ctx,
		GetPetParams{PetID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return pet.toDomain()
}

func (p PostgresAdapter) UpdatePet(
	ctx context.Context,
	petId uuid.UUID,
	updateFn func(e *patientmanagement.Pet) (*patientmanagement.Pet, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.adapters.outbound.postgres.update_pet")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	pet, err := p.GetPet(ctx, petId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedPet, err := updateFn(pet)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbPet := mapEntityPetToDB(updatedPet)

	err = p.Handler.UpdatePet(ctx, UpdatePetParams(dbPet))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) FindAllMedicalRecords(ctx context.Context) ([]query.MedicalRecord, error) {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.adapters.outbound.postgres.find_all_medical_record")
	defer span.End()

	medicalrecords, err := p.Handler.FindAllMedicalRecords(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return patientmanagementMedicalRecordsToQuery(medicalrecords)
}

func (p PostgresAdapter) FindOneMedicalRecord(ctx context.Context, id uuid.UUID) (query.MedicalRecord, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_medical_record")
	defer span.End()

	medicalrecord, err := p.GetMedicalRecord(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.MedicalRecord{}, err
	}

	return mapEntityMedicalRecordToQuery(medicalrecord), nil
}

// AddMedicalRecord adds a new exam to the database.
func (p PostgresAdapter) AddMedicalRecord(ctx context.Context, medicalrecord *patientmanagement.MedicalRecord) error {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.adapters.outbound.postgres.add_medical_record")
	defer span.End()

	dbMedicalRecord := mapEntityMedicalRecordToDB(medicalrecord)

	err := p.Handler.AddMedicalRecord(ctx, AddMedicalRecordParams(dbMedicalRecord))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetMedicalRecord(ctx context.Context, id uuid.UUID) (*patientmanagement.MedicalRecord, error) {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.adapters.outbound.postgres.get_medical_record")
	defer span.End()

	medicalrecord, err := p.Handler.GetMedicalRecord(
		ctx,
		GetMedicalRecordParams{MedicalRecordID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return medicalrecord.toDomain()
}

func (p PostgresAdapter) UpdateMedicalRecord(
	ctx context.Context,
	medicalrecordId uuid.UUID,
	updateFn func(e *patientmanagement.MedicalRecord) (*patientmanagement.MedicalRecord, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.adapters.outbound.postgres.update_medicalrecord")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	medicalrecord, err := p.GetMedicalRecord(ctx, medicalrecordId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedMedicalRecord, err := updateFn(medicalrecord)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbMedicalRecord := mapEntityMedicalRecordToDB(updatedMedicalRecord)

	err = p.Handler.UpdateMedicalRecord(ctx, UpdateMedicalRecordParams(dbMedicalRecord))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) FindAllVaccinationRecords(ctx context.Context) ([]query.VaccinationRecord, error) {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.adapters.outbound.postgres.find_all_vaccination_record")
	defer span.End()

	vaccinationrecords, err := p.Handler.FindAllVaccinationRecords(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return patientmanagementVaccinationRecordsToQuery(vaccinationrecords)
}

func (p PostgresAdapter) FindOneVaccinationRecord(ctx context.Context, id uuid.UUID) (query.VaccinationRecord, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_vaccination_record")
	defer span.End()

	vaccinationrecord, err := p.GetVaccinationRecord(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.VaccinationRecord{}, err
	}

	return mapEntityVaccinationRecordToQuery(vaccinationrecord), nil
}

// AddVaccinationRecord adds a new exam to the database.
func (p PostgresAdapter) AddVaccinationRecord(ctx context.Context, vaccinationrecord *patientmanagement.VaccinationRecord) error {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.adapters.outbound.postgres.add_vaccination_record")
	defer span.End()

	dbVaccinationRecord := mapEntityVaccinationRecordToDB(vaccinationrecord)

	err := p.Handler.AddVaccinationRecord(ctx, AddVaccinationRecordParams(dbVaccinationRecord))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetVaccinationRecord(ctx context.Context, id uuid.UUID) (*patientmanagement.VaccinationRecord, error) {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.adapters.outbound.postgres.get_vaccination_record")
	defer span.End()

	vaccinationrecord, err := p.Handler.GetVaccinationRecord(
		ctx,
		GetVaccinationRecordParams{VaccinationRecordID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return vaccinationrecord.toDomain()
}

func (p PostgresAdapter) UpdateVaccinationRecord(
	ctx context.Context,
	vaccinationrecordId uuid.UUID,
	updateFn func(e *patientmanagement.VaccinationRecord) (*patientmanagement.VaccinationRecord, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.adapters.outbound.postgres.update_vaccinationrecord")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	vaccinationrecord, err := p.GetVaccinationRecord(ctx, vaccinationrecordId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedVaccinationRecord, err := updateFn(vaccinationrecord)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbVaccinationRecord := mapEntityVaccinationRecordToDB(updatedVaccinationRecord)

	err = p.Handler.UpdateVaccinationRecord(ctx, UpdateVaccinationRecordParams(dbVaccinationRecord))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

// finishTransaction commits or rolls back the transaction based on the error state.
func (p PostgresAdapter) finishTransaction(ctx context.Context, err error, tx pgx.Tx) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			telemetree.RecordError(ctx, rollbackErr, "failed to rollback tx")

			return rollbackErr
		}

		return err
	} else {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			telemetree.RecordError(ctx, commitErr, "failed to commit tx")

			return commitErr
		}

		return nil
	}
}
