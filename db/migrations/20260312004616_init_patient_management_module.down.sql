
DROP TABLE IF EXISTS patientmanagement.pets;
DROP INDEX CONCURRENTLY idx_patient_management_pets_id;

DROP TABLE IF EXISTS patientmanagement.medical_records;
DROP INDEX CONCURRENTLY idx_patient_management_medical_records_id;

DROP TABLE IF EXISTS patientmanagement.vaccination_records;
DROP INDEX CONCURRENTLY idx_patient_management_vaccination_records_id;

DROP SCHEMA IF EXISTS patientmanagement;