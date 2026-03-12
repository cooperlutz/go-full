
-- name: GetOwner :one
SELECT * FROM ownermanagement.owners
WHERE owner_id = $1;

-- name: AddOwner :exec
INSERT INTO ownermanagement.owners (
    owner_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
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
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateOwner :exec
UPDATE ownermanagement.owners
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
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
    -- TODO
WHERE owner_id = $1;

-- name: FindOneOwner :one
SELECT * FROM ownermanagement.owners
WHERE owner_id = $1;

-- name: FindAllOwners :many
SELECT * FROM ownermanagement.owners;


-- name: GetLoyaltyAccount :one
SELECT * FROM ownermanagement.loyalty_accounts
WHERE loyalty_account_id = $1;

-- name: AddLoyaltyAccount :exec
INSERT INTO ownermanagement.loyalty_accounts (
    loyalty_account_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
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
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateLoyaltyAccount :exec
UPDATE ownermanagement.loyalty_accounts
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
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
    -- TODO
WHERE loyalty_account_id = $1;

-- name: FindOneLoyaltyAccount :one
SELECT * FROM ownermanagement.loyalty_accounts
WHERE loyalty_account_id = $1;

-- name: FindAllLoyaltyAccounts :many
SELECT * FROM ownermanagement.loyalty_accounts;

