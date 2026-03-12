CREATE SCHEMA IF NOT EXISTS retailsales;


CREATE TABLE IF NOT EXISTS retailsales.sales_orders (
    sales_order_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --order_id
    --,
    --owner_id
    --,
    --order_date
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
    --channel
    --,
    --status
    --
);

-- create index to optimize queries searching by sales_order_id
CREATE INDEX IF NOT EXISTS idx_retail_sales_sales_orders_id
ON retailsales.sales_orders (sales_order_id);

CREATE TABLE IF NOT EXISTS retailsales.shopping_carts (
    shopping_cart_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
    --,
    --cart_id
    --,
    --owner_id
    --,
    --items
    --,
    --created_at
    --,
    --updated_at
    --
);

-- create index to optimize queries searching by shopping_cart_id
CREATE INDEX IF NOT EXISTS idx_retail_sales_shopping_carts_id
ON retailsales.shopping_carts (shopping_cart_id);
