CREATE SCHEMA IF NOT EXISTS insuranceandclaims;


CREATE TABLE IF NOT EXISTS insuranceandclaims.insurance_providers (
    insurance_provider_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
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
);

-- create index to optimize queries searching by insurance_provider_id
CREATE INDEX IF NOT EXISTS idx_insurance_and_claims_insurance_providers_id
ON insuranceandclaims.insurance_providers (insurance_provider_id);

CREATE TABLE IF NOT EXISTS insuranceandclaims.insurance_claims (
    insurance_claim_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
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
);

-- create index to optimize queries searching by insurance_claim_id
CREATE INDEX IF NOT EXISTS idx_insurance_and_claims_insurance_claims_id
ON insuranceandclaims.insurance_claims (insurance_claim_id);
