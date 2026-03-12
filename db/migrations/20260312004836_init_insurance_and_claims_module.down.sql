
DROP TABLE IF EXISTS insuranceandclaims.insurance_providers;
DROP INDEX CONCURRENTLY idx_insurance_and_claims_insurance_providers_id;

DROP TABLE IF EXISTS insuranceandclaims.insurance_claims;
DROP INDEX CONCURRENTLY idx_insurance_and_claims_insurance_claims_id;

DROP SCHEMA IF EXISTS insuranceandclaims;