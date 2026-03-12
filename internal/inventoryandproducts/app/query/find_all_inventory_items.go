//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllInventoryItemsReadModel interface {
	FindAllInventoryItems(ctx context.Context) ([]InventoryItem, error)
}

type FindAllInventoryItemsHandler struct {
	readModel FindAllInventoryItemsReadModel
}

func NewFindAllInventoryItemsHandler(
	readModel FindAllInventoryItemsReadModel,
) FindAllInventoryItemsHandler {
	return FindAllInventoryItemsHandler{readModel: readModel}
}

func (h FindAllInventoryItemsHandler) Handle(ctx context.Context) ([]InventoryItem, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.query.find_all_inventory_items.handle")
	defer span.End()

	exams, err := h.readModel.FindAllInventoryItems(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
