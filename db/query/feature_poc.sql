CREATE TABLE pitch_scores (
  id SERIAL PRIMARY KEY,
  pitch TEXT NOT NULL,
  score INTEGER NOT NULL,
  critique TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- name: LogPitchScore :one
INSERT INTO pitch_scores (pitch, score, critique) VALUES ($1, $2, $3)
RETURNING id;

-- name: GetAllPitchScores :many
SELECT id, pitch, score, critique, created_at FROM pitch_scores;