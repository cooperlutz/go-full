package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/veterinarystaff/domain/veterinarystaff"
)

type SetStaffAvailability struct {
	//
	//StaffId string,
	//
	//DayOfWeek string,
	//
	//StartTime string,
	//
	//EndTime string,
	//
	//IsAvailable bool,
	//
	// TODO
}

type SetStaffAvailabilityHandler struct {
	VeterinarianRepo veterinarystaff.VeterinarianRepository

	StaffMemberRepo veterinarystaff.StaffMemberRepository

	AvailabilityScheduleRepo veterinarystaff.AvailabilityScheduleRepository
}

func NewSetStaffAvailabilityHandler(
	veterinarianRepo veterinarystaff.VeterinarianRepository,

	staffmemberRepo veterinarystaff.StaffMemberRepository,

	availabilityscheduleRepo veterinarystaff.AvailabilityScheduleRepository,
) SetStaffAvailabilityHandler {
	return SetStaffAvailabilityHandler{
		VeterinarianRepo: veterinarianRepo,

		StaffMemberRepo: staffmemberRepo,

		AvailabilityScheduleRepo: availabilityscheduleRepo,
	}
}

func (h SetStaffAvailabilityHandler) Handle(ctx context.Context, cmd SetStaffAvailability) error {
	// ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.app.command.set_staff_availability.handle")
	// defer span.End()

	// TODO
	//err = h.VeterinarianRepo.UpdateVeterinarian(ctx, uuid.MustParse(cmd.VeterinarianId), func(v *veterinarystaff.Veterinarian) (*veterinarystaff.Veterinarian, error) {
	//
	//	 err := v.SetStaffAvailability(
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

	// TODO
	//err = h.StaffMemberRepo.UpdateStaffMember(ctx, uuid.MustParse(cmd.StaffMemberId), func(s *veterinarystaff.StaffMember) (*veterinarystaff.StaffMember, error) {
	//
	//	 err := s.SetStaffAvailability(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return s, nil
	//})
	//if err != nil {
	//	return err
	//}

	// TODO
	//err = h.AvailabilityScheduleRepo.UpdateAvailabilitySchedule(ctx, uuid.MustParse(cmd.AvailabilityScheduleId), func(a *veterinarystaff.AvailabilitySchedule) (*veterinarystaff.AvailabilitySchedule, error) {
	//
	//	 err := a.SetStaffAvailability(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return a, nil
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}
