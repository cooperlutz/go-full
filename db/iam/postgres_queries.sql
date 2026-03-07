-- name: CreateUser :one
INSERT INTO iam.users (
    id,
    email,
    password_hash,
    created_at,
    last_login
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: UpdateUser :one
UPDATE iam.users
SET email = $2,
    password_hash = $3,
    last_login = $4,
    created_at = $5
WHERE id = $1
RETURNING *;

-- name: FindUserByEmail :one
SELECT * FROM iam.users
WHERE email = $1;

-- name: GetUser :one
SELECT * FROM iam.users
WHERE id = $1;

-- name: FindUserByID :one
SELECT * FROM iam.users
WHERE id = $1;

-- name: CreateRefreshToken :one
INSERT INTO iam.refresh_tokens (
    id,
    user_id,
    token,
    expires_at,
    created_at,
    revoked
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetRefreshToken :one
SELECT * FROM iam.refresh_tokens
WHERE id = $1;

-- name: RevokeRefreshToken :exec
UPDATE iam.refresh_tokens
SET revoked = TRUE
WHERE id = $1;

-- name: UpdateRefreshToken :one
UPDATE iam.refresh_tokens
SET user_id = $2,
    token = $3,
    expires_at = $4,
    created_at = $5,
    revoked = $6
WHERE id = $1
RETURNING *;