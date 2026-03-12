
-- name: GetVeterinarian :one
SELECT * FROM veterinarystaff.veterinarians
WHERE veterinarian_id = $1;

-- name: AddVeterinarian :exec
INSERT INTO veterinarystaff.veterinarians (
    veterinarian_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
    --,
    --veterinarian_id
    --,
    --first_name
    --,
    --last_name
    --,
    --email
    --,
    --phone_number
    --,
    --license_number
    --,
    --specializations
    --,
    --status
    --
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateVeterinarian :exec
UPDATE veterinarystaff.veterinarians
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
    --,
    --veterinarian_id
    --,
    --first_name
    --,
    --last_name
    --,
    --email
    --,
    --phone_number
    --,
    --license_number
    --,
    --specializations
    --,
    --status
    --
    -- TODO
WHERE veterinarian_id = $1;

-- name: FindOneVeterinarian :one
SELECT * FROM veterinarystaff.veterinarians
WHERE veterinarian_id = $1;

-- name: FindAllVeterinarians :many
SELECT * FROM veterinarystaff.veterinarians;


-- name: GetStaffMember :one
SELECT * FROM veterinarystaff.staff_members
WHERE staff_member_id = $1;

-- name: AddStaffMember :exec
INSERT INTO veterinarystaff.staff_members (
    staff_member_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
    --,
    --staff_id
    --,
    --first_name
    --,
    --last_name
    --,
    --email
    --,
    --role
    --,
    --status
    --
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateStaffMember :exec
UPDATE veterinarystaff.staff_members
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
    --,
    --staff_id
    --,
    --first_name
    --,
    --last_name
    --,
    --email
    --,
    --role
    --,
    --status
    --
    -- TODO
WHERE staff_member_id = $1;

-- name: FindOneStaffMember :one
SELECT * FROM veterinarystaff.staff_members
WHERE staff_member_id = $1;

-- name: FindAllStaffMembers :many
SELECT * FROM veterinarystaff.staff_members;


-- name: GetAvailabilitySchedule :one
SELECT * FROM veterinarystaff.availability_schedules
WHERE availability_schedule_id = $1;

-- name: AddAvailabilitySchedule :exec
INSERT INTO veterinarystaff.availability_schedules (
    availability_schedule_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
    --,
    --schedule_id
    --,
    --staff_id
    --,
    --day_of_week
    --,
    --start_time
    --,
    --end_time
    --,
    --is_available
    --
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateAvailabilitySchedule :exec
UPDATE veterinarystaff.availability_schedules
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
    --,
    --schedule_id
    --,
    --staff_id
    --,
    --day_of_week
    --,
    --start_time
    --,
    --end_time
    --,
    --is_available
    --
    -- TODO
WHERE availability_schedule_id = $1;

-- name: FindOneAvailabilitySchedule :one
SELECT * FROM veterinarystaff.availability_schedules
WHERE availability_schedule_id = $1;

-- name: FindAllAvailabilitySchedules :many
SELECT * FROM veterinarystaff.availability_schedules;

