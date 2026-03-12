package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/veterinarystaff/domain/veterinarystaff"
)

type DeactivateStaffMember struct {
	//
	//StaffId string,
	//
	//Reason string,
	//
	// TODO
}

type DeactivateStaffMemberHandler struct {
	VeterinarianRepo veterinarystaff.VeterinarianRepository

	StaffMemberRepo veterinarystaff.StaffMemberRepository

	AvailabilityScheduleRepo veterinarystaff.AvailabilityScheduleRepository
}

func NewDeactivateStaffMemberHandler(
	veterinarianRepo veterinarystaff.VeterinarianRepository,

	staffmemberRepo veterinarystaff.StaffMemberRepository,

	availabilityscheduleRepo veterinarystaff.AvailabilityScheduleRepository,
) DeactivateStaffMemberHandler {
	return DeactivateStaffMemberHandler{
		VeterinarianRepo: veterinarianRepo,

		StaffMemberRepo: staffmemberRepo,

		AvailabilityScheduleRepo: availabilityscheduleRepo,
	}
}

func (h DeactivateStaffMemberHandler) Handle(ctx context.Context, cmd DeactivateStaffMember) error {
	// ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.app.command.deactivate_staff_member.handle")
	// defer span.End()

	// TODO
	//err = h.VeterinarianRepo.UpdateVeterinarian(ctx, uuid.MustParse(cmd.VeterinarianId), func(v *veterinarystaff.Veterinarian) (*veterinarystaff.Veterinarian, error) {
	//
	//	 err := v.DeactivateStaffMember(
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
	//	 err := s.DeactivateStaffMember(
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
	//	 err := a.DeactivateStaffMember(
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
