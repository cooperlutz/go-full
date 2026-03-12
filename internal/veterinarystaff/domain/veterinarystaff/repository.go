package veterinarystaff

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type VeterinarianRepository interface {
	AddVeterinarian(ctx context.Context, veterinarian *Veterinarian) error

	GetVeterinarian(ctx context.Context, id uuid.UUID) (*Veterinarian, error)

	UpdateVeterinarian(
		ctx context.Context,
		veterinarianId uuid.UUID,
		updateFn func(e *Veterinarian) (*Veterinarian, error),
	) error
}

// MapToVeterinarian creates a Veterinarian domain object from the given parameters.
// This should ONLY BE USED when reconstructing an Veterinarian from its repository.
func MapToVeterinarian(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//veterinarianId string,
	//
	//firstName string,
	//
	//lastName string,
	//
	//email string,
	//
	//phoneNumber string,
	//
	//licenseNumber string,
	//
	//specializations *string,
	//
	//status string,
	//
) (*Veterinarian, error) {
	return &Veterinarian{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		//
		//veterinarianId: veterinarianId,
		//
		//firstName: firstName,
		//
		//lastName: lastName,
		//
		//email: email,
		//
		//phoneNumber: phoneNumber,
		//
		//licenseNumber: licenseNumber,
		//
		//specializations: specializations,
		//
		//status: status,
		//
		// TODO
	}, nil
}

type StaffMemberRepository interface {
	AddStaffMember(ctx context.Context, staffmember *StaffMember) error

	GetStaffMember(ctx context.Context, id uuid.UUID) (*StaffMember, error)

	UpdateStaffMember(
		ctx context.Context,
		staffmemberId uuid.UUID,
		updateFn func(e *StaffMember) (*StaffMember, error),
	) error
}

// MapToStaffMember creates a StaffMember domain object from the given parameters.
// This should ONLY BE USED when reconstructing an StaffMember from its repository.
func MapToStaffMember(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//staffId string,
	//
	//firstName string,
	//
	//lastName string,
	//
	//email string,
	//
	//role string,
	//
	//status string,
	//
) (*StaffMember, error) {
	return &StaffMember{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		//
		//staffId: staffId,
		//
		//firstName: firstName,
		//
		//lastName: lastName,
		//
		//email: email,
		//
		//role: role,
		//
		//status: status,
		//
		// TODO
	}, nil
}

type AvailabilityScheduleRepository interface {
	AddAvailabilitySchedule(ctx context.Context, availabilityschedule *AvailabilitySchedule) error

	GetAvailabilitySchedule(ctx context.Context, id uuid.UUID) (*AvailabilitySchedule, error)

	UpdateAvailabilitySchedule(
		ctx context.Context,
		availabilityscheduleId uuid.UUID,
		updateFn func(e *AvailabilitySchedule) (*AvailabilitySchedule, error),
	) error
}

// MapToAvailabilitySchedule creates a AvailabilitySchedule domain object from the given parameters.
// This should ONLY BE USED when reconstructing an AvailabilitySchedule from its repository.
func MapToAvailabilitySchedule(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	//
	//scheduleId string,
	//
	//staffId string,
	//
	//dayOfWeek string,
	//
	//startTime string,
	//
	//endTime string,
	//
	//isAvailable bool,
	//
) (*AvailabilitySchedule, error) {
	return &AvailabilitySchedule{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		//
		//scheduleId: scheduleId,
		//
		//staffId: staffId,
		//
		//dayOfWeek: dayOfWeek,
		//
		//startTime: startTime,
		//
		//endTime: endTime,
		//
		//isAvailable: isAvailable,
		//
		// TODO
	}, nil
}
