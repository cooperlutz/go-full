
-- name: GetProduct :one
SELECT * FROM inventoryandproducts.products
WHERE product_id = $1;

-- name: AddProduct :exec
INSERT INTO inventoryandproducts.products (
    product_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
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
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateProduct :exec
UPDATE inventoryandproducts.products
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
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
    -- TODO
WHERE product_id = $1;

-- name: FindOneProduct :one
SELECT * FROM inventoryandproducts.products
WHERE product_id = $1;

-- name: FindAllProducts :many
SELECT * FROM inventoryandproducts.products;


-- name: GetInventoryItem :one
SELECT * FROM inventoryandproducts.inventory_items
WHERE inventory_item_id = $1;

-- name: AddInventoryItem :exec
INSERT INTO inventoryandproducts.inventory_items (
    inventory_item_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
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
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateInventoryItem :exec
UPDATE inventoryandproducts.inventory_items
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
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
    -- TODO
WHERE inventory_item_id = $1;

-- name: FindOneInventoryItem :one
SELECT * FROM inventoryandproducts.inventory_items
WHERE inventory_item_id = $1;

-- name: FindAllInventoryItems :many
SELECT * FROM inventoryandproducts.inventory_items;


-- name: GetSupplier :one
SELECT * FROM inventoryandproducts.suppliers
WHERE supplier_id = $1;

-- name: AddSupplier :exec
INSERT INTO inventoryandproducts.suppliers (
    supplier_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
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
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateSupplier :exec
UPDATE inventoryandproducts.suppliers
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
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
    -- TODO
WHERE supplier_id = $1;

-- name: FindOneSupplier :one
SELECT * FROM inventoryandproducts.suppliers
WHERE supplier_id = $1;

-- name: FindAllSuppliers :many
SELECT * FROM inventoryandproducts.suppliers;


-- name: GetPurchaseOrder :one
SELECT * FROM inventoryandproducts.purchase_orders
WHERE purchase_order_id = $1;

-- name: AddPurchaseOrder :exec
INSERT INTO inventoryandproducts.purchase_orders (
    purchase_order_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
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
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdatePurchaseOrder :exec
UPDATE inventoryandproducts.purchase_orders
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
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
    -- TODO
WHERE purchase_order_id = $1;

-- name: FindOnePurchaseOrder :one
SELECT * FROM inventoryandproducts.purchase_orders
WHERE purchase_order_id = $1;

-- name: FindAllPurchaseOrders :many
SELECT * FROM inventoryandproducts.purchase_orders;

