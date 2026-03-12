
DROP TABLE IF EXISTS inventoryandproducts.products;
DROP INDEX CONCURRENTLY idx_inventory_and_products_products_id;

DROP TABLE IF EXISTS inventoryandproducts.inventory_items;
DROP INDEX CONCURRENTLY idx_inventory_and_products_inventory_items_id;

DROP TABLE IF EXISTS inventoryandproducts.suppliers;
DROP INDEX CONCURRENTLY idx_inventory_and_products_suppliers_id;

DROP TABLE IF EXISTS inventoryandproducts.purchase_orders;
DROP INDEX CONCURRENTLY idx_inventory_and_products_purchase_orders_id;

DROP SCHEMA IF EXISTS inventoryandproducts;