CREATE SCHEMA IF NOT EXISTS ownermanagement;


CREATE TABLE IF NOT EXISTS ownermanagement.owners (
    owner_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --owner_id
    --,
    --first_name
    --,
    --last_name
    --,
    --email
    --,
    --phone_number
    --,
    --address
    --,
    --communication_preference
    --,
    --loyalty_member
    --,
    --loyalty_points
    --,
    --status
    --
);

-- create index to optimize queries searching by owner_id
CREATE INDEX IF NOT EXISTS idx_owner_management_owners_id
ON ownermanagement.owners (owner_id);

CREATE TABLE IF NOT EXISTS ownermanagement.loyalty_accounts (
    loyalty_account_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --loyalty_account_id
    --,
    --owner_id
    --,
    --points_balance
    --,
    --tier
    --,
    --enrolled_date
    --
);

-- create index to optimize queries searching by loyalty_account_id
CREATE INDEX IF NOT EXISTS idx_owner_management_loyalty_accounts_id
ON ownermanagement.loyalty_accounts (loyalty_account_id);
