
-- name: GetSalesOrder :one
SELECT * FROM retailsales.sales_orders
WHERE sales_order_id = $1;

-- name: AddSalesOrder :exec
INSERT INTO retailsales.sales_orders (
    sales_order_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
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
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateSalesOrder :exec
UPDATE retailsales.sales_orders
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
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
    -- TODO
WHERE sales_order_id = $1;

-- name: FindOneSalesOrder :one
SELECT * FROM retailsales.sales_orders
WHERE sales_order_id = $1;

-- name: FindAllSalesOrders :many
SELECT * FROM retailsales.sales_orders;


-- name: GetShoppingCart :one
SELECT * FROM retailsales.shopping_carts
WHERE shopping_cart_id = $1;

-- name: AddShoppingCart :exec
INSERT INTO retailsales.shopping_carts (
    shopping_cart_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
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
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateShoppingCart :exec
UPDATE retailsales.shopping_carts
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
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
    -- TODO
WHERE shopping_cart_id = $1;

-- name: FindOneShoppingCart :one
SELECT * FROM retailsales.shopping_carts
WHERE shopping_cart_id = $1;

-- name: FindAllShoppingCarts :many
SELECT * FROM retailsales.shopping_carts;

