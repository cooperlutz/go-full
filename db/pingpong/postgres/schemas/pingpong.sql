CREATE TABLE IF NOT EXISTS pingpong (
    pingpong_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ping_or_pong VARCHAR(4),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
);
