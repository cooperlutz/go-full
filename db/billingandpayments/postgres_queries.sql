
-- name: GetInvoice :one
SELECT * FROM billingandpayments.invoices
WHERE invoice_id = $1;

-- name: AddInvoice :exec
INSERT INTO billingandpayments.invoices (
    invoice_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
    --,
    --invoice_id
    --,
    --owner_id
    --,
    --pet_id
    --,
    --appointment_id
    --,
    --line_items
    --,
    --subtotal
    --,
    --discount_amount
    --,
    --tax_amount
    --,
    --total_amount
    --,
    --status
    --,
    --issued_date
    --,
    --due_date
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

-- name: UpdateInvoice :exec
UPDATE billingandpayments.invoices
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
    --,
    --invoice_id
    --,
    --owner_id
    --,
    --pet_id
    --,
    --appointment_id
    --,
    --line_items
    --,
    --subtotal
    --,
    --discount_amount
    --,
    --tax_amount
    --,
    --total_amount
    --,
    --status
    --,
    --issued_date
    --,
    --due_date
    --
    -- TODO
WHERE invoice_id = $1;

-- name: FindOneInvoice :one
SELECT * FROM billingandpayments.invoices
WHERE invoice_id = $1;

-- name: FindAllInvoices :many
SELECT * FROM billingandpayments.invoices;


-- name: GetPayment :one
SELECT * FROM billingandpayments.payments
WHERE payment_id = $1;

-- name: AddPayment :exec
INSERT INTO billingandpayments.payments (
    payment_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
    --,
    --payment_id
    --,
    --invoice_id
    --,
    --owner_id
    --,
    --amount_paid
    --,
    --payment_method
    --,
    --payment_date
    --,
    --status
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

-- name: UpdatePayment :exec
UPDATE billingandpayments.payments
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
    --,
    --payment_id
    --,
    --invoice_id
    --,
    --owner_id
    --,
    --amount_paid
    --,
    --payment_method
    --,
    --payment_date
    --,
    --status
    --
    -- TODO
WHERE payment_id = $1;

-- name: FindOnePayment :one
SELECT * FROM billingandpayments.payments
WHERE payment_id = $1;

-- name: FindAllPayments :many
SELECT * FROM billingandpayments.payments;


-- name: GetRefund :one
SELECT * FROM billingandpayments.refunds
WHERE refund_id = $1;

-- name: AddRefund :exec
INSERT INTO billingandpayments.refunds (
    refund_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
    --,
    --refund_id
    --,
    --payment_id
    --,
    --owner_id
    --,
    --refund_amount
    --,
    --reason
    --,
    --refund_date
    --,
    --status
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

-- name: UpdateRefund :exec
UPDATE billingandpayments.refunds
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
    --,
    --refund_id
    --,
    --payment_id
    --,
    --owner_id
    --,
    --refund_amount
    --,
    --reason
    --,
    --refund_date
    --,
    --status
    --
    -- TODO
WHERE refund_id = $1;

-- name: FindOneRefund :one
SELECT * FROM billingandpayments.refunds
WHERE refund_id = $1;

-- name: FindAllRefunds :many
SELECT * FROM billingandpayments.refunds;

