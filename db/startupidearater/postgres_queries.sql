
-- name: GetStartupPitch :one
SELECT * FROM startupidearater.startup_pitchs
WHERE startup_pitch_id = $1;

-- name: AddStartupPitch :exec
INSERT INTO startupidearater.startup_pitchs (
    startup_pitch_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
    --,
    --pitch_content
    --,
    --score
    --,
    --critique
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

-- name: UpdateStartupPitch :exec
UPDATE startupidearater.startup_pitchs
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
    --,
    --pitch_content
    --,
    --score
    --,
    --critique
    --
    -- TODO
WHERE startup_pitch_id = $1;

-- name: FindOneStartupPitch :one
SELECT * FROM startupidearater.startup_pitchs
WHERE startup_pitch_id = $1;

-- name: FindAllStartupPitchs :many
SELECT * FROM startupidearater.startup_pitchs;


