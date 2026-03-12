package command

import (
	"context"

	"github.com/cooperlutz/go-full/internal/inventoryandproducts/domain/inventoryandproducts"
)

type DeductInventory struct {
	//
	//ProductId string,
	//
	//QuantityDeducted int32,
	//
	//Reason string,
	//
	// TODO
}

type DeductInventoryHandler struct {
	ProductRepo inventoryandproducts.ProductRepository

	InventoryItemRepo inventoryandproducts.InventoryItemRepository

	SupplierRepo inventoryandproducts.SupplierRepository

	PurchaseOrderRepo inventoryandproducts.PurchaseOrderRepository
}

func NewDeductInventoryHandler(
	productRepo inventoryandproducts.ProductRepository,

	inventoryitemRepo inventoryandproducts.InventoryItemRepository,

	supplierRepo inventoryandproducts.SupplierRepository,

	purchaseorderRepo inventoryandproducts.PurchaseOrderRepository,
) DeductInventoryHandler {
	return DeductInventoryHandler{
		ProductRepo: productRepo,

		InventoryItemRepo: inventoryitemRepo,

		SupplierRepo: supplierRepo,

		PurchaseOrderRepo: purchaseorderRepo,
	}
}

func (h DeductInventoryHandler) Handle(ctx context.Context, cmd DeductInventory) error {
	// ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.command.deduct_inventory.handle")
	// defer span.End()

	// TODO
	//err = h.ProductRepo.UpdateProduct(ctx, uuid.MustParse(cmd.ProductId), func(p *inventoryandproducts.Product) (*inventoryandproducts.Product, error) {
	//
	//	 err := p.DeductInventory(
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
	//	 err := i.DeductInventory(
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
	//	 err := s.DeductInventory(
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
	//	 err := p.DeductInventory(
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
