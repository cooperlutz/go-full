
DROP TABLE IF EXISTS ownermanagement.owners;
DROP INDEX CONCURRENTLY idx_owner_management_owners_id;

DROP TABLE IF EXISTS ownermanagement.loyalty_accounts;
DROP INDEX CONCURRENTLY idx_owner_management_loyalty_accounts_id;

DROP SCHEMA IF EXISTS ownermanagement;