-- name: CreatePingPong :one
INSERT INTO pingpong (pingpong_id, ping_or_pong, created_at, deleted_at, deleted) 
VALUES ($1, $2, $3, $4, $5) RETURNING *;