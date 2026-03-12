CREATE SCHEMA IF NOT EXISTS billingandpayments;


CREATE TABLE IF NOT EXISTS billingandpayments.invoices (
    invoice_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
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
);

-- create index to optimize queries searching by invoice_id
CREATE INDEX IF NOT EXISTS idx_billing_and_payments_invoices_id
ON billingandpayments.invoices (invoice_id);

CREATE TABLE IF NOT EXISTS billingandpayments.payments (
    payment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
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
);

-- create index to optimize queries searching by payment_id
CREATE INDEX IF NOT EXISTS idx_billing_and_payments_payments_id
ON billingandpayments.payments (payment_id);

CREATE TABLE IF NOT EXISTS billingandpayments.refunds (
    refund_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
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
);

-- create index to optimize queries searching by refund_id
CREATE INDEX IF NOT EXISTS idx_billing_and_payments_refunds_id
ON billingandpayments.refunds (refund_id);
