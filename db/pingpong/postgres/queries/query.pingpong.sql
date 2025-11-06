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


/* STEP 2.1. Implement Data Access Layer
- Write the associated sql logic in accordance with the specifics defined in SQLC
- once defined, run `make queries`, this will run sqlc and generate the relevant Go code
*/

-- name: FindOneByID :one
SELECT *
FROM pingpong
WHERE pingpong_id = $1;