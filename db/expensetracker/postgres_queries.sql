
-- name: GetExpense :one
SELECT * FROM expensetracker.expenses
WHERE expense_id = $1;

-- name: AddExpense :exec
INSERT INTO expensetracker.expenses (
    expense_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
    --,
    --description
    --,
    --amount
    --,
    --category
    --
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateExpense :exec
UPDATE expensetracker.expenses
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
    --,
    --description
    --,
    --amount
    --,
    --category
    --
    -- TODO
WHERE expense_id = $1;

-- name: FindOneExpense :one
SELECT * FROM expensetracker.expenses
WHERE expense_id = $1;

-- name: FindAllExpenses :many
SELECT * FROM expensetracker.expenses;


