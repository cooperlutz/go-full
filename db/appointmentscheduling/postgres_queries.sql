
-- name: GetAppointment :one
SELECT * FROM appointmentscheduling.appointments
WHERE appointment_id = $1;

-- name: AddAppointment :exec
INSERT INTO appointmentscheduling.appointments (
    appointment_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
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
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateAppointment :exec
UPDATE appointmentscheduling.appointments
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
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
    -- TODO
WHERE appointment_id = $1;

-- name: FindOneAppointment :one
SELECT * FROM appointmentscheduling.appointments
WHERE appointment_id = $1;

-- name: FindAllAppointments :many
SELECT * FROM appointmentscheduling.appointments;


-- name: GetTelemedicineSession :one
SELECT * FROM appointmentscheduling.telemedicine_sessions
WHERE telemedicine_session_id = $1;

-- name: AddTelemedicineSession :exec
INSERT INTO appointmentscheduling.telemedicine_sessions (
    telemedicine_session_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
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
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateTelemedicineSession :exec
UPDATE appointmentscheduling.telemedicine_sessions
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
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
    -- TODO
WHERE telemedicine_session_id = $1;

-- name: FindOneTelemedicineSession :one
SELECT * FROM appointmentscheduling.telemedicine_sessions
WHERE telemedicine_session_id = $1;

-- name: FindAllTelemedicineSessions :many
SELECT * FROM appointmentscheduling.telemedicine_sessions;

