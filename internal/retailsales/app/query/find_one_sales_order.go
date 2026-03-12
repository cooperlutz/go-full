//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneSalesOrder struct {
	SalesOrderID string
}

type FindOneSalesOrderReadModel interface {
	FindOneSalesOrder(ctx context.Context, salesorderId uuid.UUID) (SalesOrder, error)
}

type FindOneSalesOrderHandler struct {
	readModel FindOneSalesOrderReadModel
}

func NewFindOneSalesOrderHandler(
	readModel FindOneSalesOrderReadModel,
) FindOneSalesOrderHandler {
	return FindOneSalesOrderHandler{readModel: readModel}
}

func (h FindOneSalesOrderHandler) Handle(ctx context.Context, qry FindOneSalesOrder) (SalesOrder, error) {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.app.query.find_one_sales_order.handle")
	defer span.End()

	salesorder, err := h.readModel.FindOneSalesOrder(ctx, uuid.MustParse(qry.SalesOrderID))
	if err != nil {
		return SalesOrder{}, err
	}

	return salesorder, nil
}
