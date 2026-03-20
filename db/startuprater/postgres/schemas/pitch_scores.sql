CREATE SCHEMA IF NOT EXISTS startup_rater;

CREATE TABLE IF NOT EXISTS startup_rater.pitch_scores (
    pitch_score_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    --
    pitch TEXT NOT NULL,
    score INT NOT NULL,
    critique TEXT NOT NULL
);
