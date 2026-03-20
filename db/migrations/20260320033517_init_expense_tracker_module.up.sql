CREATE SCHEMA IF NOT EXISTS expensetracker;


CREATE TABLE IF NOT EXISTS expensetracker.expenses (
    expense_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --description
    --,
    --amount
    --,
    --category
    --
);

-- create index to optimize queries searching by expense_id
CREATE INDEX IF NOT EXISTS idx_expense_tracker_expenses_id
ON expensetracker.expenses (expense_id);


