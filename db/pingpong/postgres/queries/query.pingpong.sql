-- name: FindAllPing :many
SELECT *
FROM pingpong
WHERE ping_or_pong = 'ping';

-- name: FindAllPong :many
SELECT * FROM pingpong
WHERE ping_or_pong = 'pong';

-- name: FindAll :many
SELECT *
FROM pingpong;