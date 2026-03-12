//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllSalesOrdersReadModel interface {
	FindAllSalesOrders(ctx context.Context) ([]SalesOrder, error)
}

type FindAllSalesOrdersHandler struct {
	readModel FindAllSalesOrdersReadModel
}

func NewFindAllSalesOrdersHandler(
	readModel FindAllSalesOrdersReadModel,
) FindAllSalesOrdersHandler {
	return FindAllSalesOrdersHandler{readModel: readModel}
}

func (h FindAllSalesOrdersHandler) Handle(ctx context.Context) ([]SalesOrder, error) {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.app.query.find_all_sales_orders.handle")
	defer span.End()

	exams, err := h.readModel.FindAllSalesOrders(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
