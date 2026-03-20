-- name: FindAllPitchScores :many
SELECT
    pitch_score_id,
    created_at,
    updated_at,
    deleted_at,
    deleted,
    pitch,
    score,
    critique
FROM startup_rater.pitch_scores
ORDER BY created_at DESC;
