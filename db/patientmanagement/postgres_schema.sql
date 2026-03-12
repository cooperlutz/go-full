CREATE SCHEMA IF NOT EXISTS patientmanagement;


CREATE TABLE IF NOT EXISTS patientmanagement.pets (
    pet_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --pet_id
    --,
    --owner_id
    --,
    --name
    --,
    --species
    --,
    --breed
    --,
    --date_of_birth
    --,
    --gender
    --,
    --weight
    --,
    --microchip_number
    --,
    --status
    --
);

-- create index to optimize queries searching by pet_id
CREATE INDEX IF NOT EXISTS idx_patient_management_pets_id
ON patientmanagement.pets (pet_id);

CREATE TABLE IF NOT EXISTS patientmanagement.medical_records (
    medical_record_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --medical_record_id
    --,
    --pet_id
    --,
    --veterinarian_id
    --,
    --visit_date
    --,
    --diagnosis
    --,
    --treatment
    --,
    --notes
    --,
    --follow_up_required
    --
);

-- create index to optimize queries searching by medical_record_id
CREATE INDEX IF NOT EXISTS idx_patient_management_medical_records_id
ON patientmanagement.medical_records (medical_record_id);

CREATE TABLE IF NOT EXISTS patientmanagement.vaccination_records (
    vaccination_record_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --vaccination_id
    --,
    --pet_id
    --,
    --vaccine_name
    --,
    --administered_date
    --,
    --expiry_date
    --,
    --administered_by
    --
);

-- create index to optimize queries searching by vaccination_record_id
CREATE INDEX IF NOT EXISTS idx_patient_management_vaccination_records_id
ON patientmanagement.vaccination_records (vaccination_record_id);
