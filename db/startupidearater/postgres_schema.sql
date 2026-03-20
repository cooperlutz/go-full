CREATE SCHEMA IF NOT EXISTS startupidearater;


CREATE TABLE IF NOT EXISTS startupidearater.startup_pitchs (
    startup_pitch_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --pitch_content TEXT NOT NULL
    --,
    --score SMALLINT
    --,
    --critique TEXT
    --
);

-- create index to optimize queries searching by startup_pitch_id
CREATE INDEX IF NOT EXISTS idx_startup_idea_rater_startup_pitchs_id
ON startupidearater.startup_pitchs (startup_pitch_id);


