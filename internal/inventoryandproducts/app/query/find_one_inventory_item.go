//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneInventoryItem struct {
	InventoryItemID string
}

type FindOneInventoryItemReadModel interface {
	FindOneInventoryItem(ctx context.Context, inventoryitemId uuid.UUID) (InventoryItem, error)
}

type FindOneInventoryItemHandler struct {
	readModel FindOneInventoryItemReadModel
}

func NewFindOneInventoryItemHandler(
	readModel FindOneInventoryItemReadModel,
) FindOneInventoryItemHandler {
	return FindOneInventoryItemHandler{readModel: readModel}
}

func (h FindOneInventoryItemHandler) Handle(ctx context.Context, qry FindOneInventoryItem) (InventoryItem, error) {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.query.find_one_inventory_item.handle")
	defer span.End()

	inventoryitem, err := h.readModel.FindOneInventoryItem(ctx, uuid.MustParse(qry.InventoryItemID))
	if err != nil {
		return InventoryItem{}, err
	}

	return inventoryitem, nil
}
