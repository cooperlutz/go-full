package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/inventoryandproducts/domain/inventoryandproducts"
)

type RestockInventory struct {
	//
	//ProductId string,
	//
	//QuantityAdded int32,
	//
	//PurchaseOrderId *string,
	//
	// TODO
}

type RestockInventoryHandler struct {
	ProductRepo inventoryandproducts.ProductRepository

	InventoryItemRepo inventoryandproducts.InventoryItemRepository

	SupplierRepo inventoryandproducts.SupplierRepository

	PurchaseOrderRepo inventoryandproducts.PurchaseOrderRepository
}

func NewRestockInventoryHandler(
	productRepo inventoryandproducts.ProductRepository,

	inventoryitemRepo inventoryandproducts.InventoryItemRepository,

	supplierRepo inventoryandproducts.SupplierRepository,

	purchaseorderRepo inventoryandproducts.PurchaseOrderRepository,
) RestockInventoryHandler {
	return RestockInventoryHandler{
		ProductRepo: productRepo,

		InventoryItemRepo: inventoryitemRepo,

		SupplierRepo: supplierRepo,

		PurchaseOrderRepo: purchaseorderRepo,
	}
}

func (h RestockInventoryHandler) Handle(ctx context.Context, cmd RestockInventory) error {
	// ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.command.restock_inventory.handle")
	// defer span.End()

	// TODO
	//err = h.ProductRepo.UpdateProduct(ctx, uuid.MustParse(cmd.ProductId), func(p *inventoryandproducts.Product) (*inventoryandproducts.Product, error) {
	//
	//	 err := p.RestockInventory(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return p, nil
	//})
	//if err != nil {
	//	return err
	//}

	// TODO
	//err = h.InventoryItemRepo.UpdateInventoryItem(ctx, uuid.MustParse(cmd.InventoryItemId), func(i *inventoryandproducts.InventoryItem) (*inventoryandproducts.InventoryItem, error) {
	//
	//	 err := i.RestockInventory(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return i, nil
	//})
	//if err != nil {
	//	return err
	//}

	// TODO
	//err = h.SupplierRepo.UpdateSupplier(ctx, uuid.MustParse(cmd.SupplierId), func(s *inventoryandproducts.Supplier) (*inventoryandproducts.Supplier, error) {
	//
	//	 err := s.RestockInventory(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return s, nil
	//})
	//if err != nil {
	//	return err
	//}

	// TODO
	//err = h.PurchaseOrderRepo.UpdatePurchaseOrder(ctx, uuid.MustParse(cmd.PurchaseOrderId), func(p *inventoryandproducts.PurchaseOrder) (*inventoryandproducts.PurchaseOrder, error) {
	//
	//	 err := p.RestockInventory(
	//	 	)
	//	 if err != nil {
	//	 	telemetree.RecordError(ctx, err)
	//
	//	 	return nil, err
	//	 }
	//
	//	return p, nil
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}
