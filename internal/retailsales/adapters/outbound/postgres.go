//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package outbound

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/cooperlutz/go-full/internal/retailsales/app/query"
	"github.com/cooperlutz/go-full/internal/retailsales/domain/retailsales"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// PostgresAdapter implements the examination repository using Postgres as the data store.
type PostgresAdapter struct {
	Handler IQuerierRetailSales
}

// NewPostgresAdapter creates a new instance of PostgresAdapter.
func NewPostgresAdapter(db deebee.IDatabase) PostgresAdapter {
	return PostgresAdapter{
		Handler: NewQueriesWrapper(db),
	}
}

func (p PostgresAdapter) FindAllSalesOrders(ctx context.Context) ([]query.SalesOrder, error) {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.adapters.outbound.postgres.find_all_sales_order")
	defer span.End()

	salesorders, err := p.Handler.FindAllSalesOrders(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return retailsalesSalesOrdersToQuery(salesorders)
}

func (p PostgresAdapter) FindOneSalesOrder(ctx context.Context, id uuid.UUID) (query.SalesOrder, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_sales_order")
	defer span.End()

	salesorder, err := p.GetSalesOrder(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.SalesOrder{}, err
	}

	return mapEntitySalesOrderToQuery(salesorder), nil
}

// AddSalesOrder adds a new exam to the database.
func (p PostgresAdapter) AddSalesOrder(ctx context.Context, salesorder *retailsales.SalesOrder) error {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.adapters.outbound.postgres.add_sales_order")
	defer span.End()

	dbSalesOrder := mapEntitySalesOrderToDB(salesorder)

	err := p.Handler.AddSalesOrder(ctx, AddSalesOrderParams(dbSalesOrder))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetSalesOrder(ctx context.Context, id uuid.UUID) (*retailsales.SalesOrder, error) {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.adapters.outbound.postgres.get_sales_order")
	defer span.End()

	salesorder, err := p.Handler.GetSalesOrder(
		ctx,
		GetSalesOrderParams{SalesOrderID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return salesorder.toDomain()
}

func (p PostgresAdapter) UpdateSalesOrder(
	ctx context.Context,
	salesorderId uuid.UUID,
	updateFn func(e *retailsales.SalesOrder) (*retailsales.SalesOrder, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.adapters.outbound.postgres.update_salesorder")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	salesorder, err := p.GetSalesOrder(ctx, salesorderId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedSalesOrder, err := updateFn(salesorder)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbSalesOrder := mapEntitySalesOrderToDB(updatedSalesOrder)

	err = p.Handler.UpdateSalesOrder(ctx, UpdateSalesOrderParams(dbSalesOrder))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) FindAllShoppingCarts(ctx context.Context) ([]query.ShoppingCart, error) {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.adapters.outbound.postgres.find_all_shopping_cart")
	defer span.End()

	shoppingcarts, err := p.Handler.FindAllShoppingCarts(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return retailsalesShoppingCartsToQuery(shoppingcarts)
}

func (p PostgresAdapter) FindOneShoppingCart(ctx context.Context, id uuid.UUID) (query.ShoppingCart, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_shopping_cart")
	defer span.End()

	shoppingcart, err := p.GetShoppingCart(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.ShoppingCart{}, err
	}

	return mapEntityShoppingCartToQuery(shoppingcart), nil
}

// AddShoppingCart adds a new exam to the database.
func (p PostgresAdapter) AddShoppingCart(ctx context.Context, shoppingcart *retailsales.ShoppingCart) error {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.adapters.outbound.postgres.add_shopping_cart")
	defer span.End()

	dbShoppingCart := mapEntityShoppingCartToDB(shoppingcart)

	err := p.Handler.AddShoppingCart(ctx, AddShoppingCartParams(dbShoppingCart))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetShoppingCart(ctx context.Context, id uuid.UUID) (*retailsales.ShoppingCart, error) {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.adapters.outbound.postgres.get_shopping_cart")
	defer span.End()

	shoppingcart, err := p.Handler.GetShoppingCart(
		ctx,
		GetShoppingCartParams{ShoppingCartID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return shoppingcart.toDomain()
}

func (p PostgresAdapter) UpdateShoppingCart(
	ctx context.Context,
	shoppingcartId uuid.UUID,
	updateFn func(e *retailsales.ShoppingCart) (*retailsales.ShoppingCart, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.adapters.outbound.postgres.update_shoppingcart")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	shoppingcart, err := p.GetShoppingCart(ctx, shoppingcartId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedShoppingCart, err := updateFn(shoppingcart)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbShoppingCart := mapEntityShoppingCartToDB(updatedShoppingCart)

	err = p.Handler.UpdateShoppingCart(ctx, UpdateShoppingCartParams(dbShoppingCart))
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
