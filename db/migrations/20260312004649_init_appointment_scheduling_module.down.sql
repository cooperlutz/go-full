
DROP TABLE IF EXISTS appointmentscheduling.appointments;
DROP INDEX CONCURRENTLY idx_appointment_scheduling_appointments_id;

DROP TABLE IF EXISTS appointmentscheduling.telemedicine_sessions;
DROP INDEX CONCURRENTLY idx_appointment_scheduling_telemedicine_sessions_id;

DROP SCHEMA IF EXISTS appointmentscheduling;