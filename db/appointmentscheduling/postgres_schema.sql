CREATE SCHEMA IF NOT EXISTS appointmentscheduling;


CREATE TABLE IF NOT EXISTS appointmentscheduling.appointments (
    appointment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --appointment_id
    --,
    --pet_id
    --,
    --owner_id
    --,
    --veterinarian_id
    --,
    --appointment_type
    --,
    --scheduled_date
    --,
    --scheduled_time
    --,
    --duration_minutes
    --,
    --status
    --,
    --notes
    --,
    --is_telemedicine
    --
);

-- create index to optimize queries searching by appointment_id
CREATE INDEX IF NOT EXISTS idx_appointment_scheduling_appointments_id
ON appointmentscheduling.appointments (appointment_id);

CREATE TABLE IF NOT EXISTS appointmentscheduling.telemedicine_sessions (
    telemedicine_session_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --session_id
    --,
    --appointment_id
    --,
    --session_url
    --,
    --started_at
    --,
    --ended_at
    --,
    --status
    --
);

-- create index to optimize queries searching by telemedicine_session_id
CREATE INDEX IF NOT EXISTS idx_appointment_scheduling_telemedicine_sessions_id
ON appointmentscheduling.telemedicine_sessions (telemedicine_session_id);
