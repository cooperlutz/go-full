//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package outbound

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/cooperlutz/go-full/internal/inventoryandproducts/app/query"
	"github.com/cooperlutz/go-full/internal/inventoryandproducts/domain/inventoryandproducts"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// PostgresAdapter implements the examination repository using Postgres as the data store.
type PostgresAdapter struct {
	Handler IQuerierInventoryAndProducts
}

// NewPostgresAdapter creates a new instance of PostgresAdapter.
func NewPostgresAdapter(db deebee.IDatabase) PostgresAdapter {
	return PostgresAdapter{
		Handler: NewQueriesWrapper(db),
	}
}

func (p PostgresAdapter) FindAllProducts(ctx context.Context) ([]query.Product, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.find_all_product")
	defer span.End()

	products, err := p.Handler.FindAllProducts(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return inventoryandproductsProductsToQuery(products)
}

func (p PostgresAdapter) FindOneProduct(ctx context.Context, id uuid.UUID) (query.Product, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_product")
	defer span.End()

	product, err := p.GetProduct(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.Product{}, err
	}

	return mapEntityProductToQuery(product), nil
}

// AddProduct adds a new exam to the database.
func (p PostgresAdapter) AddProduct(ctx context.Context, product *inventoryandproducts.Product) error {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.add_product")
	defer span.End()

	dbProduct := mapEntityProductToDB(product)

	err := p.Handler.AddProduct(ctx, AddProductParams(dbProduct))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetProduct(ctx context.Context, id uuid.UUID) (*inventoryandproducts.Product, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.get_product")
	defer span.End()

	product, err := p.Handler.GetProduct(
		ctx,
		GetProductParams{ProductID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return product.toDomain()
}

func (p PostgresAdapter) UpdateProduct(
	ctx context.Context,
	productId uuid.UUID,
	updateFn func(e *inventoryandproducts.Product) (*inventoryandproducts.Product, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.update_product")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	product, err := p.GetProduct(ctx, productId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedProduct, err := updateFn(product)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbProduct := mapEntityProductToDB(updatedProduct)

	err = p.Handler.UpdateProduct(ctx, UpdateProductParams(dbProduct))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) FindAllInventoryItems(ctx context.Context) ([]query.InventoryItem, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.find_all_inventory_item")
	defer span.End()

	inventoryitems, err := p.Handler.FindAllInventoryItems(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return inventoryandproductsInventoryItemsToQuery(inventoryitems)
}

func (p PostgresAdapter) FindOneInventoryItem(ctx context.Context, id uuid.UUID) (query.InventoryItem, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_inventory_item")
	defer span.End()

	inventoryitem, err := p.GetInventoryItem(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.InventoryItem{}, err
	}

	return mapEntityInventoryItemToQuery(inventoryitem), nil
}

// AddInventoryItem adds a new exam to the database.
func (p PostgresAdapter) AddInventoryItem(ctx context.Context, inventoryitem *inventoryandproducts.InventoryItem) error {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.add_inventory_item")
	defer span.End()

	dbInventoryItem := mapEntityInventoryItemToDB(inventoryitem)

	err := p.Handler.AddInventoryItem(ctx, AddInventoryItemParams(dbInventoryItem))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetInventoryItem(ctx context.Context, id uuid.UUID) (*inventoryandproducts.InventoryItem, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.get_inventory_item")
	defer span.End()

	inventoryitem, err := p.Handler.GetInventoryItem(
		ctx,
		GetInventoryItemParams{InventoryItemID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return inventoryitem.toDomain()
}

func (p PostgresAdapter) UpdateInventoryItem(
	ctx context.Context,
	inventoryitemId uuid.UUID,
	updateFn func(e *inventoryandproducts.InventoryItem) (*inventoryandproducts.InventoryItem, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.update_inventoryitem")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	inventoryitem, err := p.GetInventoryItem(ctx, inventoryitemId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedInventoryItem, err := updateFn(inventoryitem)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbInventoryItem := mapEntityInventoryItemToDB(updatedInventoryItem)

	err = p.Handler.UpdateInventoryItem(ctx, UpdateInventoryItemParams(dbInventoryItem))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) FindAllSuppliers(ctx context.Context) ([]query.Supplier, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.find_all_supplier")
	defer span.End()

	suppliers, err := p.Handler.FindAllSuppliers(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return inventoryandproductsSuppliersToQuery(suppliers)
}

func (p PostgresAdapter) FindOneSupplier(ctx context.Context, id uuid.UUID) (query.Supplier, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_supplier")
	defer span.End()

	supplier, err := p.GetSupplier(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.Supplier{}, err
	}

	return mapEntitySupplierToQuery(supplier), nil
}

// AddSupplier adds a new exam to the database.
func (p PostgresAdapter) AddSupplier(ctx context.Context, supplier *inventoryandproducts.Supplier) error {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.add_supplier")
	defer span.End()

	dbSupplier := mapEntitySupplierToDB(supplier)

	err := p.Handler.AddSupplier(ctx, AddSupplierParams(dbSupplier))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetSupplier(ctx context.Context, id uuid.UUID) (*inventoryandproducts.Supplier, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.get_supplier")
	defer span.End()

	supplier, err := p.Handler.GetSupplier(
		ctx,
		GetSupplierParams{SupplierID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return supplier.toDomain()
}

func (p PostgresAdapter) UpdateSupplier(
	ctx context.Context,
	supplierId uuid.UUID,
	updateFn func(e *inventoryandproducts.Supplier) (*inventoryandproducts.Supplier, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.update_supplier")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	supplier, err := p.GetSupplier(ctx, supplierId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedSupplier, err := updateFn(supplier)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbSupplier := mapEntitySupplierToDB(updatedSupplier)

	err = p.Handler.UpdateSupplier(ctx, UpdateSupplierParams(dbSupplier))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) FindAllPurchaseOrders(ctx context.Context) ([]query.PurchaseOrder, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.find_all_purchase_order")
	defer span.End()

	purchaseorders, err := p.Handler.FindAllPurchaseOrders(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return inventoryandproductsPurchaseOrdersToQuery(purchaseorders)
}

func (p PostgresAdapter) FindOnePurchaseOrder(ctx context.Context, id uuid.UUID) (query.PurchaseOrder, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_purchase_order")
	defer span.End()

	purchaseorder, err := p.GetPurchaseOrder(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.PurchaseOrder{}, err
	}

	return mapEntityPurchaseOrderToQuery(purchaseorder), nil
}

// AddPurchaseOrder adds a new exam to the database.
func (p PostgresAdapter) AddPurchaseOrder(ctx context.Context, purchaseorder *inventoryandproducts.PurchaseOrder) error {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.add_purchase_order")
	defer span.End()

	dbPurchaseOrder := mapEntityPurchaseOrderToDB(purchaseorder)

	err := p.Handler.AddPurchaseOrder(ctx, AddPurchaseOrderParams(dbPurchaseOrder))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetPurchaseOrder(ctx context.Context, id uuid.UUID) (*inventoryandproducts.PurchaseOrder, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.get_purchase_order")
	defer span.End()

	purchaseorder, err := p.Handler.GetPurchaseOrder(
		ctx,
		GetPurchaseOrderParams{PurchaseOrderID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return purchaseorder.toDomain()
}

func (p PostgresAdapter) UpdatePurchaseOrder(
	ctx context.Context,
	purchaseorderId uuid.UUID,
	updateFn func(e *inventoryandproducts.PurchaseOrder) (*inventoryandproducts.PurchaseOrder, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.adapters.outbound.postgres.update_purchaseorder")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	purchaseorder, err := p.GetPurchaseOrder(ctx, purchaseorderId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedPurchaseOrder, err := updateFn(purchaseorder)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbPurchaseOrder := mapEntityPurchaseOrderToDB(updatedPurchaseOrder)

	err = p.Handler.UpdatePurchaseOrder(ctx, UpdatePurchaseOrderParams(dbPurchaseOrder))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

// finishTransaction commits or rolls back the transaction based on the error state.
func (p PostgresAdapter) finishTransaction(ctx context.Context, err error, tx pgx.Tx) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			telemetree.RecordError(ctx, rollbackErr, "failed to rollback tx")

			return rollbackErr
		}

		return err
	} else {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			telemetree.RecordError(ctx, commitErr, "failed to commit tx")

			return commitErr
		}

		return nil
	}
}
