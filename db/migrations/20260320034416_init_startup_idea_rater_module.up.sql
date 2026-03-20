CREATE SCHEMA IF NOT EXISTS startupidearater;


CREATE TABLE IF NOT EXISTS startupidearater.pitchs (
    pitch_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --content
    --,
    --score
    --,
    --critique
    --
);

-- create index to optimize queries searching by pitch_id
CREATE INDEX IF NOT EXISTS idx_startup_idea_rater_pitchs_id
ON startupidearater.pitchs (pitch_id);


