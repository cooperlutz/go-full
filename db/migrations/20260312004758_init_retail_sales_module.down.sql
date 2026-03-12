
DROP TABLE IF EXISTS retailsales.sales_orders;
DROP INDEX CONCURRENTLY idx_retail_sales_sales_orders_id;

DROP TABLE IF EXISTS retailsales.shopping_carts;
DROP INDEX CONCURRENTLY idx_retail_sales_shopping_carts_id;

DROP SCHEMA IF EXISTS retailsales;