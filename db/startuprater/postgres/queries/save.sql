-- name: SavePitchScore :exec
INSERT INTO startup_rater.pitch_scores (
    pitch_score_id,
    created_at,
    updated_at,
    deleted_at,
    deleted,
    pitch,
    score,
    critique
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
);
