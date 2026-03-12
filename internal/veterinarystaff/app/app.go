package app

import (
	"github.com/cooperlutz/go-full/internal/veterinarystaff/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/veterinarystaff/app/command"
	"github.com/cooperlutz/go-full/internal/veterinarystaff/app/event"
	"github.com/cooperlutz/go-full/internal/veterinarystaff/app/query"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Events   Events
}

type Commands struct {
	OnboardVeterinarian command.OnboardVeterinarianHandler

	UpdateVeterinarianProfile command.UpdateVeterinarianProfileHandler

	SetStaffAvailability command.SetStaffAvailabilityHandler

	DeactivateStaffMember command.DeactivateStaffMemberHandler
}

type Queries struct {
	FindAllVeterinarians query.FindAllVeterinariansHandler
	FindOneVeterinarian  query.FindOneVeterinarianHandler

	FindAllStaffMembers query.FindAllStaffMembersHandler
	FindOneStaffMember  query.FindOneStaffMemberHandler

	FindAllAvailabilitySchedules query.FindAllAvailabilitySchedulesHandler
	FindOneAvailabilitySchedule  query.FindOneAvailabilityScheduleHandler
}

type Events struct {
	VeterinarianOnboarded event.VeterinarianOnboardedHandler

	VeterinarianProfileUpdated event.VeterinarianProfileUpdatedHandler

	StaffAvailabilityUpdated event.StaffAvailabilityUpdatedHandler

	StaffMemberDeactivated event.StaffMemberDeactivatedHandler
}

// NewApplication initializes the VeterinaryStaff application with its dependencies.
func NewApplication( //nolint:funlen // it's fine
	pgconn deebee.IDatabase,
	pubSub eeventdriven.IPubSubEventProcessor,
) (Application, error) {
	veterinarianRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	staffmemberRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	availabilityscheduleRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	app := Application{
		Commands: Commands{
			OnboardVeterinarian: command.NewOnboardVeterinarianHandler(

				veterinarianRepository,

				staffmemberRepository,

				availabilityscheduleRepository,
			),
			UpdateVeterinarianProfile: command.NewUpdateVeterinarianProfileHandler(

				veterinarianRepository,

				staffmemberRepository,

				availabilityscheduleRepository,
			),
			SetStaffAvailability: command.NewSetStaffAvailabilityHandler(

				veterinarianRepository,

				staffmemberRepository,

				availabilityscheduleRepository,
			),
			DeactivateStaffMember: command.NewDeactivateStaffMemberHandler(

				veterinarianRepository,

				staffmemberRepository,

				availabilityscheduleRepository,
			),
		},
		Queries: Queries{
			FindAllVeterinarians: query.NewFindAllVeterinariansHandler(
				veterinarianRepository,
			),
			FindOneVeterinarian: query.NewFindOneVeterinarianHandler(
				veterinarianRepository,
			),

			FindAllStaffMembers: query.NewFindAllStaffMembersHandler(
				staffmemberRepository,
			),
			FindOneStaffMember: query.NewFindOneStaffMemberHandler(
				staffmemberRepository,
			),

			FindAllAvailabilitySchedules: query.NewFindAllAvailabilitySchedulesHandler(
				availabilityscheduleRepository,
			),
			FindOneAvailabilitySchedule: query.NewFindOneAvailabilityScheduleHandler(
				availabilityscheduleRepository,
			),
		},
		Events: Events{
			VeterinarianOnboarded: event.NewVeterinarianOnboardedHandler(
				pubSub,
			),

			VeterinarianProfileUpdated: event.NewVeterinarianProfileUpdatedHandler(
				pubSub,
			),

			StaffAvailabilityUpdated: event.NewStaffAvailabilityUpdatedHandler(
				pubSub,
			),

			StaffMemberDeactivated: event.NewStaffMemberDeactivatedHandler(
				pubSub,
			),
		},
	}

	return app, nil
}
