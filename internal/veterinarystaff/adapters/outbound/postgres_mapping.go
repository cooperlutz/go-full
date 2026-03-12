package outbound

import (
	"github.com/cooperlutz/go-full/internal/veterinarystaff/app/query"
	"github.com/cooperlutz/go-full/internal/veterinarystaff/domain/veterinarystaff"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

// toDomain maps the VeterinarianVeterinarian to the domain entity.
func (e VeterinarystaffVeterinarian) toDomain() (*veterinarystaff.Veterinarian, error) {
	return veterinarystaff.MapToVeterinarian(
		e.VeterinarianID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.VeterinarianId,
		//
		//e.FirstName,
		//
		//e.LastName,
		//
		//e.Email,
		//
		//e.PhoneNumber,
		//
		//e.LicenseNumber,
		//
		//e.Specializations,
		//
		//e.Status,
		//
		// TODO
	)
}

// toQueryVeterinarian maps the veterinarianVeterinarian to the query.Veterinarian.
func (e VeterinarystaffVeterinarian) toQueryVeterinarian() (query.Veterinarian, error) {
	veterinarian, err := e.toDomain()
	if err != nil {
		return query.Veterinarian{}, err
	}

	return mapEntityVeterinarianToQuery(veterinarian), nil
}

// veterinarianVeterinariansToQuery maps a slice of VeterinarianVeterinarian to a slice of query.Veterinarian entities.
func veterinarystaffVeterinariansToQuery(veterinarians []VeterinarystaffVeterinarian) ([]query.Veterinarian, error) {
	var domainVeterinarians []query.Veterinarian

	for _, veterinarian := range veterinarians {
		queryVeterinarian, err := veterinarian.toQueryVeterinarian()
		if err != nil {
			return nil, err
		}

		domainVeterinarians = append(domainVeterinarians, queryVeterinarian)
	}

	return domainVeterinarians, nil
}

// mapEntityVeterinarianToDB maps a domain Veterinarian entity to the VeterinarianVeterinarian database model.
func mapEntityVeterinarianToDB(veterinarian *veterinarystaff.Veterinarian) VeterinarystaffVeterinarian {
	createdAt := veterinarian.GetCreatedAtTime()
	updatedAt := veterinarian.GetUpdatedAtTime()

	return VeterinarystaffVeterinarian{
		VeterinarianID: pgxutil.UUIDToPgtypeUUID(veterinarian.GetIdUUID()),
		CreatedAt:      pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:      pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:        veterinarian.IsDeleted(),
		DeletedAt:      pgxutil.TimeToTimestampz(veterinarian.GetDeletedAtTime()),
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

// mapEntityVeterinarianToQuery maps a domain Veterinarian entity to a query.Veterinarian.
func mapEntityVeterinarianToQuery(veterinarian *veterinarystaff.Veterinarian) query.Veterinarian {
	return query.Veterinarian{
		// TODO
	}
}

// toDomain maps the StaffmemberStaffMember to the domain entity.
func (e VeterinarystaffStaffMember) toDomain() (*veterinarystaff.StaffMember, error) {
	return veterinarystaff.MapToStaffMember(
		e.StaffMemberID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.StaffId,
		//
		//e.FirstName,
		//
		//e.LastName,
		//
		//e.Email,
		//
		//e.Role,
		//
		//e.Status,
		//
		// TODO
	)
}

// toQueryStaffMember maps the staffmemberStaffMember to the query.StaffMember.
func (e VeterinarystaffStaffMember) toQueryStaffMember() (query.StaffMember, error) {
	staffmember, err := e.toDomain()
	if err != nil {
		return query.StaffMember{}, err
	}

	return mapEntityStaffMemberToQuery(staffmember), nil
}

// staffmemberStaffMembersToQuery maps a slice of StaffMemberStaffMember to a slice of query.StaffMember entities.
func veterinarystaffStaffMembersToQuery(staffmembers []VeterinarystaffStaffMember) ([]query.StaffMember, error) {
	var domainStaffMembers []query.StaffMember

	for _, staffmember := range staffmembers {
		queryStaffMember, err := staffmember.toQueryStaffMember()
		if err != nil {
			return nil, err
		}

		domainStaffMembers = append(domainStaffMembers, queryStaffMember)
	}

	return domainStaffMembers, nil
}

// mapEntityStaffMemberToDB maps a domain StaffMember entity to the StaffMemberStaffMember database model.
func mapEntityStaffMemberToDB(staffmember *veterinarystaff.StaffMember) VeterinarystaffStaffMember {
	createdAt := staffmember.GetCreatedAtTime()
	updatedAt := staffmember.GetUpdatedAtTime()

	return VeterinarystaffStaffMember{
		StaffMemberID: pgxutil.UUIDToPgtypeUUID(staffmember.GetIdUUID()),
		CreatedAt:     pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:     pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:       staffmember.IsDeleted(),
		DeletedAt:     pgxutil.TimeToTimestampz(staffmember.GetDeletedAtTime()),
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

// mapEntityStaffMemberToQuery maps a domain StaffMember entity to a query.StaffMember.
func mapEntityStaffMemberToQuery(staffmember *veterinarystaff.StaffMember) query.StaffMember {
	return query.StaffMember{
		// TODO
	}
}

// toDomain maps the AvailabilityscheduleAvailabilitySchedule to the domain entity.
func (e VeterinarystaffAvailabilitySchedule) toDomain() (*veterinarystaff.AvailabilitySchedule, error) {
	return veterinarystaff.MapToAvailabilitySchedule(
		e.AvailabilityScheduleID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		//
		//e.ScheduleId,
		//
		//e.StaffId,
		//
		//e.DayOfWeek,
		//
		//e.StartTime,
		//
		//e.EndTime,
		//
		//e.IsAvailable,
		//
		// TODO
	)
}

// toQueryAvailabilitySchedule maps the availabilityscheduleAvailabilitySchedule to the query.AvailabilitySchedule.
func (e VeterinarystaffAvailabilitySchedule) toQueryAvailabilitySchedule() (query.AvailabilitySchedule, error) {
	availabilityschedule, err := e.toDomain()
	if err != nil {
		return query.AvailabilitySchedule{}, err
	}

	return mapEntityAvailabilityScheduleToQuery(availabilityschedule), nil
}

// availabilityscheduleAvailabilitySchedulesToQuery maps a slice of AvailabilityScheduleAvailabilitySchedule to a slice of query.AvailabilitySchedule entities.
func veterinarystaffAvailabilitySchedulesToQuery(availabilityschedules []VeterinarystaffAvailabilitySchedule) ([]query.AvailabilitySchedule, error) {
	var domainAvailabilitySchedules []query.AvailabilitySchedule

	for _, availabilityschedule := range availabilityschedules {
		queryAvailabilitySchedule, err := availabilityschedule.toQueryAvailabilitySchedule()
		if err != nil {
			return nil, err
		}

		domainAvailabilitySchedules = append(domainAvailabilitySchedules, queryAvailabilitySchedule)
	}

	return domainAvailabilitySchedules, nil
}

// mapEntityAvailabilityScheduleToDB maps a domain AvailabilitySchedule entity to the AvailabilityScheduleAvailabilitySchedule database model.
func mapEntityAvailabilityScheduleToDB(availabilityschedule *veterinarystaff.AvailabilitySchedule) VeterinarystaffAvailabilitySchedule {
	createdAt := availabilityschedule.GetCreatedAtTime()
	updatedAt := availabilityschedule.GetUpdatedAtTime()

	return VeterinarystaffAvailabilitySchedule{
		AvailabilityScheduleID: pgxutil.UUIDToPgtypeUUID(availabilityschedule.GetIdUUID()),
		CreatedAt:              pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:              pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:                availabilityschedule.IsDeleted(),
		DeletedAt:              pgxutil.TimeToTimestampz(availabilityschedule.GetDeletedAtTime()),
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

// mapEntityAvailabilityScheduleToQuery maps a domain AvailabilitySchedule entity to a query.AvailabilitySchedule.
func mapEntityAvailabilityScheduleToQuery(availabilityschedule *veterinarystaff.AvailabilitySchedule) query.AvailabilitySchedule {
	return query.AvailabilitySchedule{
		// TODO
	}
}
