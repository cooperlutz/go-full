package app

import (
	"github.com/cooperlutz/go-full/internal/patientmanagement/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/patientmanagement/app/command"
	"github.com/cooperlutz/go-full/internal/patientmanagement/app/event"
	"github.com/cooperlutz/go-full/internal/patientmanagement/app/query"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Events   Events
}

type Commands struct {
	RegisterPet command.RegisterPetHandler

	UpdatePetDetails command.UpdatePetDetailsHandler

	AddMedicalRecord command.AddMedicalRecordHandler

	RecordVaccination command.RecordVaccinationHandler

	DeactivatePet command.DeactivatePetHandler
}

type Queries struct {
	FindAllPets query.FindAllPetsHandler
	FindOnePet  query.FindOnePetHandler

	FindAllMedicalRecords query.FindAllMedicalRecordsHandler
	FindOneMedicalRecord  query.FindOneMedicalRecordHandler

	FindAllVaccinationRecords query.FindAllVaccinationRecordsHandler
	FindOneVaccinationRecord  query.FindOneVaccinationRecordHandler
}

type Events struct {
	PetRegistered event.PetRegisteredHandler

	PetDetailsUpdated event.PetDetailsUpdatedHandler

	MedicalRecordAdded event.MedicalRecordAddedHandler

	VaccinationRecorded event.VaccinationRecordedHandler

	VaccinationDueReminderScheduled event.VaccinationDueReminderScheduledHandler

	PetDeactivated event.PetDeactivatedHandler

	AppointmentCompleted event.AppointmentCompletedHandler
}

// NewApplication initializes the PatientManagement application with its dependencies.
func NewApplication( //nolint:funlen // it's fine
	pgconn deebee.IDatabase,
	pubSub eeventdriven.IPubSubEventProcessor,
) (Application, error) {
	petRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	medicalrecordRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	vaccinationrecordRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	app := Application{
		Commands: Commands{
			RegisterPet: command.NewRegisterPetHandler(

				petRepository,

				medicalrecordRepository,

				vaccinationrecordRepository,
			),
			UpdatePetDetails: command.NewUpdatePetDetailsHandler(

				petRepository,

				medicalrecordRepository,

				vaccinationrecordRepository,
			),
			AddMedicalRecord: command.NewAddMedicalRecordHandler(

				petRepository,

				medicalrecordRepository,

				vaccinationrecordRepository,
			),
			RecordVaccination: command.NewRecordVaccinationHandler(

				petRepository,

				medicalrecordRepository,

				vaccinationrecordRepository,
			),
			DeactivatePet: command.NewDeactivatePetHandler(

				petRepository,

				medicalrecordRepository,

				vaccinationrecordRepository,
			),
		},
		Queries: Queries{
			FindAllPets: query.NewFindAllPetsHandler(
				petRepository,
			),
			FindOnePet: query.NewFindOnePetHandler(
				petRepository,
			),

			FindAllMedicalRecords: query.NewFindAllMedicalRecordsHandler(
				medicalrecordRepository,
			),
			FindOneMedicalRecord: query.NewFindOneMedicalRecordHandler(
				medicalrecordRepository,
			),

			FindAllVaccinationRecords: query.NewFindAllVaccinationRecordsHandler(
				vaccinationrecordRepository,
			),
			FindOneVaccinationRecord: query.NewFindOneVaccinationRecordHandler(
				vaccinationrecordRepository,
			),
		},
		Events: Events{
			PetRegistered: event.NewPetRegisteredHandler(
				pubSub,
			),

			PetDetailsUpdated: event.NewPetDetailsUpdatedHandler(
				pubSub,
			),

			MedicalRecordAdded: event.NewMedicalRecordAddedHandler(
				pubSub,
			),

			VaccinationRecorded: event.NewVaccinationRecordedHandler(
				pubSub,
			),

			VaccinationDueReminderScheduled: event.NewVaccinationDueReminderScheduledHandler(
				pubSub,
			),

			PetDeactivated: event.NewPetDeactivatedHandler(
				pubSub,
			),

			AppointmentCompleted: event.NewAppointmentCompletedHandler(
				pubSub,
			),
		},
	}

	return app, nil
}
