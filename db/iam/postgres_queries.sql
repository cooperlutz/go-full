-- name: CreateUser :one
INSERT INTO iam.users (
    id,
    email,
    password_hash,
    last_login
) VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM iam.users
WHERE email = $1;

-- name: GetUserByID :one
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
WHERE token = $1;

-- name: RevokeRefreshToken :exec
UPDATE iam.refresh_tokens
SET revoked = TRUE
WHERE token = $1;