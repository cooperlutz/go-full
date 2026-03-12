CREATE SCHEMA IF NOT EXISTS veterinarystaff;


CREATE TABLE IF NOT EXISTS veterinarystaff.veterinarians (
    veterinarian_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
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
);

-- create index to optimize queries searching by veterinarian_id
CREATE INDEX IF NOT EXISTS idx_veterinary_staff_veterinarians_id
ON veterinarystaff.veterinarians (veterinarian_id);

CREATE TABLE IF NOT EXISTS veterinarystaff.staff_members (
    staff_member_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
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
);

-- create index to optimize queries searching by staff_member_id
CREATE INDEX IF NOT EXISTS idx_veterinary_staff_staff_members_id
ON veterinarystaff.staff_members (staff_member_id);

CREATE TABLE IF NOT EXISTS veterinarystaff.availability_schedules (
    availability_schedule_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
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
);

-- create index to optimize queries searching by availability_schedule_id
CREATE INDEX IF NOT EXISTS idx_veterinary_staff_availability_schedules_id
ON veterinarystaff.availability_schedules (availability_schedule_id);
