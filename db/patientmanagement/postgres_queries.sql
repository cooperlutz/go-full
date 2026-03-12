
-- name: GetPet :one
SELECT * FROM patientmanagement.pets
WHERE pet_id = $1;

-- name: AddPet :exec
INSERT INTO patientmanagement.pets (
    pet_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
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
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdatePet :exec
UPDATE patientmanagement.pets
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
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
    -- TODO
WHERE pet_id = $1;

-- name: FindOnePet :one
SELECT * FROM patientmanagement.pets
WHERE pet_id = $1;

-- name: FindAllPets :many
SELECT * FROM patientmanagement.pets;


-- name: GetMedicalRecord :one
SELECT * FROM patientmanagement.medical_records
WHERE medical_record_id = $1;

-- name: AddMedicalRecord :exec
INSERT INTO patientmanagement.medical_records (
    medical_record_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
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
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateMedicalRecord :exec
UPDATE patientmanagement.medical_records
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
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
    -- TODO
WHERE medical_record_id = $1;

-- name: FindOneMedicalRecord :one
SELECT * FROM patientmanagement.medical_records
WHERE medical_record_id = $1;

-- name: FindAllMedicalRecords :many
SELECT * FROM patientmanagement.medical_records;


-- name: GetVaccinationRecord :one
SELECT * FROM patientmanagement.vaccination_records
WHERE vaccination_record_id = $1;

-- name: AddVaccinationRecord :exec
INSERT INTO patientmanagement.vaccination_records (
    vaccination_record_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
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
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateVaccinationRecord :exec
UPDATE patientmanagement.vaccination_records
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
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
    -- TODO
WHERE vaccination_record_id = $1;

-- name: FindOneVaccinationRecord :one
SELECT * FROM patientmanagement.vaccination_records
WHERE vaccination_record_id = $1;

-- name: FindAllVaccinationRecords :many
SELECT * FROM patientmanagement.vaccination_records;

