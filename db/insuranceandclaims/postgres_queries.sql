
-- name: GetInsuranceProvider :one
SELECT * FROM insuranceandclaims.insurance_providers
WHERE insurance_provider_id = $1;

-- name: AddInsuranceProvider :exec
INSERT INTO insuranceandclaims.insurance_providers (
    insurance_provider_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
    --,
    --provider_id
    --,
    --name
    --,
    --contact_name
    --,
    --email
    --,
    --phone_number
    --,
    --claim_submission_url
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

-- name: UpdateInsuranceProvider :exec
UPDATE insuranceandclaims.insurance_providers
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
    --,
    --provider_id
    --,
    --name
    --,
    --contact_name
    --,
    --email
    --,
    --phone_number
    --,
    --claim_submission_url
    --,
    --status
    --
    -- TODO
WHERE insurance_provider_id = $1;

-- name: FindOneInsuranceProvider :one
SELECT * FROM insuranceandclaims.insurance_providers
WHERE insurance_provider_id = $1;

-- name: FindAllInsuranceProviders :many
SELECT * FROM insuranceandclaims.insurance_providers;


-- name: GetInsuranceClaim :one
SELECT * FROM insuranceandclaims.insurance_claims
WHERE insurance_claim_id = $1;

-- name: AddInsuranceClaim :exec
INSERT INTO insuranceandclaims.insurance_claims (
    insurance_claim_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
    --,
    --claim_id
    --,
    --owner_id
    --,
    --pet_id
    --,
    --provider_id
    --,
    --invoice_id
    --,
    --policy_number
    --,
    --claim_amount
    --,
    --approved_amount
    --,
    --submission_date
    --,
    --resolution_date
    --,
    --status
    --,
    --notes
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

-- name: UpdateInsuranceClaim :exec
UPDATE insuranceandclaims.insurance_claims
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
    --,
    --claim_id
    --,
    --owner_id
    --,
    --pet_id
    --,
    --provider_id
    --,
    --invoice_id
    --,
    --policy_number
    --,
    --claim_amount
    --,
    --approved_amount
    --,
    --submission_date
    --,
    --resolution_date
    --,
    --status
    --,
    --notes
    --
    -- TODO
WHERE insurance_claim_id = $1;

-- name: FindOneInsuranceClaim :one
SELECT * FROM insuranceandclaims.insurance_claims
WHERE insurance_claim_id = $1;

-- name: FindAllInsuranceClaims :many
SELECT * FROM insuranceandclaims.insurance_claims;

