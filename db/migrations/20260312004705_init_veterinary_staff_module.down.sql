
DROP TABLE IF EXISTS veterinarystaff.veterinarians;
DROP INDEX CONCURRENTLY idx_veterinary_staff_veterinarians_id;

DROP TABLE IF EXISTS veterinarystaff.staff_members;
DROP INDEX CONCURRENTLY idx_veterinary_staff_staff_members_id;

DROP TABLE IF EXISTS veterinarystaff.availability_schedules;
DROP INDEX CONCURRENTLY idx_veterinary_staff_availability_schedules_id;

DROP SCHEMA IF EXISTS veterinarystaff;