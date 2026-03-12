//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllPurchaseOrdersReadModel interface {
	FindAllPurchaseOrders(ctx context.Context) ([]PurchaseOrder, error)
}

type FindAllPurchaseOrdersHandler struct {
	readModel FindAllPurchaseOrdersReadModel
}

func NewFindAllPurchaseOrdersHandler(
	readModel FindAllPurchaseOrdersReadModel,
) FindAllPurchaseOrdersHandler {
	return FindAllPurchaseOrdersHandler{readModel: readModel}
}

func (h FindAllPurchaseOrdersHandler) Handle(ctx context.Context) ([]PurchaseOrder, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.query.find_all_purchase_orders.handle")
	defer span.End()

	exams, err := h.readModel.FindAllPurchaseOrders(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
