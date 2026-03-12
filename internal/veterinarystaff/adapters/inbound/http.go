package inbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/veterinarystaff/app"
	"github.com/cooperlutz/go-full/internal/veterinarystaff/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// HttpAdapter represents the HTTP server for the VeterinaryStaff module.
type HttpAdapter struct {
	app app.Application
}

// NewHttpAdapter creates a new HttpAdapter instance with the provided VeterinaryStaff application.
func NewHttpAdapter(application app.Application) HttpAdapter {
	return HttpAdapter{
		app: application,
	}
}

// StrictHandler returns a strict HTTP handler for the VeterinaryStaff module.
func (h HttpAdapter) StrictHandler() ServerInterface {
	return NewStrictHandler(h, nil)
}

// (GET /v1/veterinarians).
func (h HttpAdapter) FindAllVeterinarians(ctx context.Context, request FindAllVeterinariansRequestObject) (FindAllVeterinariansResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "veterinarian.adapters.inbound.http.find_all_veterinarians")
	defer span.End()

	veterinarian, err := h.app.Queries.FindAllVeterinarians.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseVeterinarians []Veterinarian
	for _, e := range veterinarian {
		responseVeterinarians = append(responseVeterinarians, queryVeterinarianToHttpVeterinarian(e))
	}

	return FindAllVeterinarians200JSONResponse(responseVeterinarians), nil
}

// (GET /v1/veterinarian/{veterinarianId}).
func (h HttpAdapter) FindOneVeterinarian(ctx context.Context, request FindOneVeterinarianRequestObject) (FindOneVeterinarianResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_veterinarian")
	defer span.End()

	veterinarian, err := h.app.Queries.FindOneVeterinarian.Handle(ctx, query.FindOneVeterinarian{VeterinarianID: request.VeterinarianId})
	if err != nil {
		return nil, err
	}

	return FindOneVeterinarian200JSONResponse(queryVeterinarianToHttpVeterinarian(veterinarian)), nil
}

func queryVeterinarianToHttpVeterinarian(e query.Veterinarian) Veterinarian {
	return Veterinarian{
		//
		//VeterinarianId: GetVeterinarianId(),
		//
		//FirstName: GetFirstName(),
		//
		//LastName: GetLastName(),
		//
		//Email: GetEmail(),
		//
		//PhoneNumber: GetPhoneNumber(),
		//
		//LicenseNumber: GetLicenseNumber(),
		//
		//Specializations: GetSpecializations(),
		//
		//Status: GetStatus(),
		//
		// TODO
	}
}

// (GET /v1/staffmembers).
func (h HttpAdapter) FindAllStaffMembers(ctx context.Context, request FindAllStaffMembersRequestObject) (FindAllStaffMembersResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "staffmember.adapters.inbound.http.find_all_staffmembers")
	defer span.End()

	staffmember, err := h.app.Queries.FindAllStaffMembers.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseStaffMembers []StaffMember
	for _, e := range staffmember {
		responseStaffMembers = append(responseStaffMembers, queryStaffMemberToHttpStaffMember(e))
	}

	return FindAllStaffMembers200JSONResponse(responseStaffMembers), nil
}

// (GET /v1/staffmember/{staff_memberId}).
func (h HttpAdapter) FindOneStaffMember(ctx context.Context, request FindOneStaffMemberRequestObject) (FindOneStaffMemberResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_staff_member")
	defer span.End()

	staffmember, err := h.app.Queries.FindOneStaffMember.Handle(ctx, query.FindOneStaffMember{StaffMemberID: request.StaffMemberId})
	if err != nil {
		return nil, err
	}

	return FindOneStaffMember200JSONResponse(queryStaffMemberToHttpStaffMember(staffmember)), nil
}

func queryStaffMemberToHttpStaffMember(e query.StaffMember) StaffMember {
	return StaffMember{
		//
		//StaffId: GetStaffId(),
		//
		//FirstName: GetFirstName(),
		//
		//LastName: GetLastName(),
		//
		//Email: GetEmail(),
		//
		//Role: GetRole(),
		//
		//Status: GetStatus(),
		//
		// TODO
	}
}

// (GET /v1/availabilityschedules).
func (h HttpAdapter) FindAllAvailabilitySchedules(ctx context.Context, request FindAllAvailabilitySchedulesRequestObject) (FindAllAvailabilitySchedulesResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "availabilityschedule.adapters.inbound.http.find_all_availabilityschedules")
	defer span.End()

	availabilityschedule, err := h.app.Queries.FindAllAvailabilitySchedules.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseAvailabilitySchedules []AvailabilitySchedule
	for _, e := range availabilityschedule {
		responseAvailabilitySchedules = append(responseAvailabilitySchedules, queryAvailabilityScheduleToHttpAvailabilitySchedule(e))
	}

	return FindAllAvailabilitySchedules200JSONResponse(responseAvailabilitySchedules), nil
}

// (GET /v1/availabilityschedule/{availability_scheduleId}).
func (h HttpAdapter) FindOneAvailabilitySchedule(ctx context.Context, request FindOneAvailabilityScheduleRequestObject) (FindOneAvailabilityScheduleResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "work.adapters.inbound.http.find_one_availability_schedule")
	defer span.End()

	availabilityschedule, err := h.app.Queries.FindOneAvailabilitySchedule.Handle(ctx, query.FindOneAvailabilitySchedule{AvailabilityScheduleID: request.AvailabilityScheduleId})
	if err != nil {
		return nil, err
	}

	return FindOneAvailabilitySchedule200JSONResponse(queryAvailabilityScheduleToHttpAvailabilitySchedule(availabilityschedule)), nil
}

func queryAvailabilityScheduleToHttpAvailabilitySchedule(e query.AvailabilitySchedule) AvailabilitySchedule {
	return AvailabilitySchedule{
		//
		//ScheduleId: GetScheduleId(),
		//
		//StaffId: GetStaffId(),
		//
		//DayOfWeek: GetDayOfWeek(),
		//
		//StartTime: GetStartTime(),
		//
		//EndTime: GetEndTime(),
		//
		//IsAvailable: GetIsAvailable(),
		//
		// TODO
	}
}
