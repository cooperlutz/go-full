package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/veterinarystaff/domain/veterinarystaff"
)

type UpdateVeterinarianProfile struct {
	//
	//VeterinarianId string,
	//
	//Specializations *string,
	//
	//PhoneNumber *string,
	//
	//Email *string,
	//
	// TODO
}

type UpdateVeterinarianProfileHandler struct {
	VeterinarianRepo veterinarystaff.VeterinarianRepository

	StaffMemberRepo veterinarystaff.StaffMemberRepository

	AvailabilityScheduleRepo veterinarystaff.AvailabilityScheduleRepository
}

func NewUpdateVeterinarianProfileHandler(
	veterinarianRepo veterinarystaff.VeterinarianRepository,

	staffmemberRepo veterinarystaff.StaffMemberRepository,

	availabilityscheduleRepo veterinarystaff.AvailabilityScheduleRepository,
) UpdateVeterinarianProfileHandler {
	return UpdateVeterinarianProfileHandler{
		VeterinarianRepo: veterinarianRepo,

		StaffMemberRepo: staffmemberRepo,

		AvailabilityScheduleRepo: availabilityscheduleRepo,
	}
}

func (h UpdateVeterinarianProfileHandler) Handle(ctx context.Context, cmd UpdateVeterinarianProfile) error {
	// ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.app.command.update_veterinarian_profile.handle")
	// defer span.End()

	// TODO
	//err = h.VeterinarianRepo.UpdateVeterinarian(ctx, uuid.MustParse(cmd.VeterinarianId), func(v *veterinarystaff.Veterinarian) (*veterinarystaff.Veterinarian, error) {
	//
	//	 err := v.UpdateVeterinarianProfile(
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
	//	 err := s.UpdateVeterinarianProfile(
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
	//	 err := a.UpdateVeterinarianProfile(
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
