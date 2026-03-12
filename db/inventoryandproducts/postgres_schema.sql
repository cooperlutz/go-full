CREATE SCHEMA IF NOT EXISTS inventoryandproducts;


CREATE TABLE IF NOT EXISTS inventoryandproducts.products (
    product_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --product_id
    --,
    --name
    --,
    --description
    --,
    --category
    --,
    --unit_price
    --,
    --sku
    --,
    --supplier_id
    --,
    --is_active
    --
);

-- create index to optimize queries searching by product_id
CREATE INDEX IF NOT EXISTS idx_inventory_and_products_products_id
ON inventoryandproducts.products (product_id);

CREATE TABLE IF NOT EXISTS inventoryandproducts.inventory_items (
    inventory_item_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --inventory_item_id
    --,
    --product_id
    --,
    --quantity_on_hand
    --,
    --reorder_threshold
    --,
    --reorder_quantity
    --,
    --last_restocked_date
    --
);

-- create index to optimize queries searching by inventory_item_id
CREATE INDEX IF NOT EXISTS idx_inventory_and_products_inventory_items_id
ON inventoryandproducts.inventory_items (inventory_item_id);

CREATE TABLE IF NOT EXISTS inventoryandproducts.suppliers (
    supplier_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --supplier_id
    --,
    --name
    --,
    --contact_name
    --,
    --email
    --,
    --phone_number
    --,
    --product_categories
    --,
    --status
    --
);

-- create index to optimize queries searching by supplier_id
CREATE INDEX IF NOT EXISTS idx_inventory_and_products_suppliers_id
ON inventoryandproducts.suppliers (supplier_id);

CREATE TABLE IF NOT EXISTS inventoryandproducts.purchase_orders (
    purchase_order_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --purchase_order_id
    --,
    --supplier_id
    --,
    --order_date
    --,
    --expected_delivery_date
    --,
    --line_items
    --,
    --total_cost
    --,
    --status
    --
);

-- create index to optimize queries searching by purchase_order_id
CREATE INDEX IF NOT EXISTS idx_inventory_and_products_purchase_orders_id
ON inventoryandproducts.purchase_orders (purchase_order_id);
