
DROP TABLE IF EXISTS billingandpayments.invoices;
DROP INDEX CONCURRENTLY idx_billing_and_payments_invoices_id;

DROP TABLE IF EXISTS billingandpayments.payments;
DROP INDEX CONCURRENTLY idx_billing_and_payments_payments_id;

DROP TABLE IF EXISTS billingandpayments.refunds;
DROP INDEX CONCURRENTLY idx_billing_and_payments_refunds_id;

DROP SCHEMA IF EXISTS billingandpayments;