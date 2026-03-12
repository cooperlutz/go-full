//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOnePurchaseOrder struct {
	PurchaseOrderID string
}

type FindOnePurchaseOrderReadModel interface {
	FindOnePurchaseOrder(ctx context.Context, purchaseorderId uuid.UUID) (PurchaseOrder, error)
}

type FindOnePurchaseOrderHandler struct {
	readModel FindOnePurchaseOrderReadModel
}

func NewFindOnePurchaseOrderHandler(
	readModel FindOnePurchaseOrderReadModel,
) FindOnePurchaseOrderHandler {
	return FindOnePurchaseOrderHandler{readModel: readModel}
}

func (h FindOnePurchaseOrderHandler) Handle(ctx context.Context, qry FindOnePurchaseOrder) (PurchaseOrder, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.query.find_one_purchase_order.handle")
	defer span.End()

	purchaseorder, err := h.readModel.FindOnePurchaseOrder(ctx, uuid.MustParse(qry.PurchaseOrderID))
	if err != nil {
		return PurchaseOrder{}, err
	}

	return purchaseorder, nil
}
