//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package outbound

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/cooperlutz/go-full/internal/veterinarystaff/app/query"
	"github.com/cooperlutz/go-full/internal/veterinarystaff/domain/veterinarystaff"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// PostgresAdapter implements the examination repository using Postgres as the data store.
type PostgresAdapter struct {
	Handler IQuerierVeterinaryStaff
}

// NewPostgresAdapter creates a new instance of PostgresAdapter.
func NewPostgresAdapter(db deebee.IDatabase) PostgresAdapter {
	return PostgresAdapter{
		Handler: NewQueriesWrapper(db),
	}
}

func (p PostgresAdapter) FindAllVeterinarians(ctx context.Context) ([]query.Veterinarian, error) {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.adapters.outbound.postgres.find_all_veterinarian")
	defer span.End()

	veterinarians, err := p.Handler.FindAllVeterinarians(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return veterinarystaffVeterinariansToQuery(veterinarians)
}

func (p PostgresAdapter) FindOneVeterinarian(ctx context.Context, id uuid.UUID) (query.Veterinarian, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_veterinarian")
	defer span.End()

	veterinarian, err := p.GetVeterinarian(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.Veterinarian{}, err
	}

	return mapEntityVeterinarianToQuery(veterinarian), nil
}

// AddVeterinarian adds a new exam to the database.
func (p PostgresAdapter) AddVeterinarian(ctx context.Context, veterinarian *veterinarystaff.Veterinarian) error {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.adapters.outbound.postgres.add_veterinarian")
	defer span.End()

	dbVeterinarian := mapEntityVeterinarianToDB(veterinarian)

	err := p.Handler.AddVeterinarian(ctx, AddVeterinarianParams(dbVeterinarian))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetVeterinarian(ctx context.Context, id uuid.UUID) (*veterinarystaff.Veterinarian, error) {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.adapters.outbound.postgres.get_veterinarian")
	defer span.End()

	veterinarian, err := p.Handler.GetVeterinarian(
		ctx,
		GetVeterinarianParams{VeterinarianID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return veterinarian.toDomain()
}

func (p PostgresAdapter) UpdateVeterinarian(
	ctx context.Context,
	veterinarianId uuid.UUID,
	updateFn func(e *veterinarystaff.Veterinarian) (*veterinarystaff.Veterinarian, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.adapters.outbound.postgres.update_veterinarian")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	veterinarian, err := p.GetVeterinarian(ctx, veterinarianId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedVeterinarian, err := updateFn(veterinarian)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbVeterinarian := mapEntityVeterinarianToDB(updatedVeterinarian)

	err = p.Handler.UpdateVeterinarian(ctx, UpdateVeterinarianParams(dbVeterinarian))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) FindAllStaffMembers(ctx context.Context) ([]query.StaffMember, error) {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.adapters.outbound.postgres.find_all_staff_member")
	defer span.End()

	staffmembers, err := p.Handler.FindAllStaffMembers(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return veterinarystaffStaffMembersToQuery(staffmembers)
}

func (p PostgresAdapter) FindOneStaffMember(ctx context.Context, id uuid.UUID) (query.StaffMember, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_staff_member")
	defer span.End()

	staffmember, err := p.GetStaffMember(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.StaffMember{}, err
	}

	return mapEntityStaffMemberToQuery(staffmember), nil
}

// AddStaffMember adds a new exam to the database.
func (p PostgresAdapter) AddStaffMember(ctx context.Context, staffmember *veterinarystaff.StaffMember) error {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.adapters.outbound.postgres.add_staff_member")
	defer span.End()

	dbStaffMember := mapEntityStaffMemberToDB(staffmember)

	err := p.Handler.AddStaffMember(ctx, AddStaffMemberParams(dbStaffMember))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetStaffMember(ctx context.Context, id uuid.UUID) (*veterinarystaff.StaffMember, error) {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.adapters.outbound.postgres.get_staff_member")
	defer span.End()

	staffmember, err := p.Handler.GetStaffMember(
		ctx,
		GetStaffMemberParams{StaffMemberID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return staffmember.toDomain()
}

func (p PostgresAdapter) UpdateStaffMember(
	ctx context.Context,
	staffmemberId uuid.UUID,
	updateFn func(e *veterinarystaff.StaffMember) (*veterinarystaff.StaffMember, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.adapters.outbound.postgres.update_staffmember")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	staffmember, err := p.GetStaffMember(ctx, staffmemberId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedStaffMember, err := updateFn(staffmember)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbStaffMember := mapEntityStaffMemberToDB(updatedStaffMember)

	err = p.Handler.UpdateStaffMember(ctx, UpdateStaffMemberParams(dbStaffMember))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) FindAllAvailabilitySchedules(ctx context.Context) ([]query.AvailabilitySchedule, error) {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.adapters.outbound.postgres.find_all_availability_schedule")
	defer span.End()

	availabilityschedules, err := p.Handler.FindAllAvailabilitySchedules(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return veterinarystaffAvailabilitySchedulesToQuery(availabilityschedules)
}

func (p PostgresAdapter) FindOneAvailabilitySchedule(ctx context.Context, id uuid.UUID) (query.AvailabilitySchedule, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_availability_schedule")
	defer span.End()

	availabilityschedule, err := p.GetAvailabilitySchedule(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.AvailabilitySchedule{}, err
	}

	return mapEntityAvailabilityScheduleToQuery(availabilityschedule), nil
}

// AddAvailabilitySchedule adds a new exam to the database.
func (p PostgresAdapter) AddAvailabilitySchedule(ctx context.Context, availabilityschedule *veterinarystaff.AvailabilitySchedule) error {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.adapters.outbound.postgres.add_availability_schedule")
	defer span.End()

	dbAvailabilitySchedule := mapEntityAvailabilityScheduleToDB(availabilityschedule)

	err := p.Handler.AddAvailabilitySchedule(ctx, AddAvailabilityScheduleParams(dbAvailabilitySchedule))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetAvailabilitySchedule(ctx context.Context, id uuid.UUID) (*veterinarystaff.AvailabilitySchedule, error) {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.adapters.outbound.postgres.get_availability_schedule")
	defer span.End()

	availabilityschedule, err := p.Handler.GetAvailabilitySchedule(
		ctx,
		GetAvailabilityScheduleParams{AvailabilityScheduleID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return availabilityschedule.toDomain()
}

func (p PostgresAdapter) UpdateAvailabilitySchedule(
	ctx context.Context,
	availabilityscheduleId uuid.UUID,
	updateFn func(e *veterinarystaff.AvailabilitySchedule) (*veterinarystaff.AvailabilitySchedule, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.adapters.outbound.postgres.update_availabilityschedule")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	availabilityschedule, err := p.GetAvailabilitySchedule(ctx, availabilityscheduleId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedAvailabilitySchedule, err := updateFn(availabilityschedule)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbAvailabilitySchedule := mapEntityAvailabilityScheduleToDB(updatedAvailabilitySchedule)

	err = p.Handler.UpdateAvailabilitySchedule(ctx, UpdateAvailabilityScheduleParams(dbAvailabilitySchedule))
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
