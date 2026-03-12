package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/veterinarystaff/domain/veterinarystaff"
)

type OnboardVeterinarian struct {
	//
	//FirstName string,
	//
	//LastName string,
	//
	//Email string,
	//
	//PhoneNumber string,
	//
	//LicenseNumber string,
	//
	//Specializations *string,
	//
	// TODO
}

type OnboardVeterinarianHandler struct {
	VeterinarianRepo veterinarystaff.VeterinarianRepository

	StaffMemberRepo veterinarystaff.StaffMemberRepository

	AvailabilityScheduleRepo veterinarystaff.AvailabilityScheduleRepository
}

func NewOnboardVeterinarianHandler(
	veterinarianRepo veterinarystaff.VeterinarianRepository,

	staffmemberRepo veterinarystaff.StaffMemberRepository,

	availabilityscheduleRepo veterinarystaff.AvailabilityScheduleRepository,
) OnboardVeterinarianHandler {
	return OnboardVeterinarianHandler{
		VeterinarianRepo: veterinarianRepo,

		StaffMemberRepo: staffmemberRepo,

		AvailabilityScheduleRepo: availabilityscheduleRepo,
	}
}

func (h OnboardVeterinarianHandler) Handle(ctx context.Context, cmd OnboardVeterinarian) error {
	// ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.app.command.onboard_veterinarian.handle")
	// defer span.End()

	// TODO
	//err = h.VeterinarianRepo.UpdateVeterinarian(ctx, uuid.MustParse(cmd.VeterinarianId), func(v *veterinarystaff.Veterinarian) (*veterinarystaff.Veterinarian, error) {
	//
	//	 err := v.OnboardVeterinarian(
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
	//	 err := s.OnboardVeterinarian(
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
	//	 err := a.OnboardVeterinarian(
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
