-- name: TotalNumberOfPingPongs :one
SELECT COUNT(*) FROM pingpong;

-- name: TotalNumberOfPings :one
SELECT COUNT(*) FROM pingpong WHERE ping_or_pong = 'ping';

-- name: TotalNumberOfPongs :one
SELECT COUNT(*) FROM pingpong WHERE ping_or_pong = 'pong';

-- name: CountPerDay :many
SELECT DATE(created_at) AS creation_date, COUNT(*) AS count_created
FROM pingpong
GROUP BY creation_date
ORDER BY creation_date ASC;

-- name: FrequencyDistribution :many
SELECT ping_or_pong, COUNT(*) AS frequency
FROM pingpong
GROUP BY ping_or_pong
ORDER BY frequency DESC;

-- name: FrequencyDistributionByDay :many
SELECT DATE(created_at) AS creation_date, ping_or_pong, COUNT(*) AS frequency
FROM pingpong
GROUP BY creation_date, ping_or_pong
ORDER BY creation_date ASC, frequency DESC;

-- name: FrequencyDistributionByDayPong :many
SELECT DATE(created_at) AS creation_date, ping_or_pong AS pong, COUNT(*) AS frequency
FROM pingpong
WHERE pong = 'pong'
GROUP BY creation_date, pong
ORDER BY creation_date ASC, frequency DESC;

-- name: FrequencyDistributionByDayPing :many
SELECT DATE(created_at) AS creation_date, ping_or_pong AS ping, COUNT(*) AS frequency
FROM pingpong
WHERE ping = 'ping'
GROUP BY creation_date, ping
ORDER BY creation_date ASC, frequency DESC;

