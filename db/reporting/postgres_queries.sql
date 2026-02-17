
-- name: GetReportByName :one
SELECT * FROM reporting.reports
WHERE name = $1;

-- name: AddReport :exec
INSERT INTO reporting.reports (
    name,
    description,
    component_definitions
) VALUES (
    $1,
    $2,
    $3
);

-- name: UpdateReport :exec
UPDATE reporting.reports
SET
    name = $1,
    description = $2,
    component_definitions = $3
WHERE name = $1;

-- name: GetMetricByName :one
SELECT * FROM reporting.metrics
WHERE name = $1;

-- name: AddMetric :exec
INSERT INTO reporting.metrics (
    name,
    value
) VALUES (
    $1,
    $2
);

-- name: UpdateMetric :exec
UPDATE reporting.metrics
SET
    name = $1,
    value = $2
WHERE name = $1;