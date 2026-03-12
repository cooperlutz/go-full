package query

type Product struct {
	//
	//productId: productId,
	//
	//name: name,
	//
	//description: description,
	//
	//category: category,
	//
	//unitPrice: unitPrice,
	//
	//sku: sku,
	//
	//supplierId: supplierId,
	//
	//isActive: isActive,
	//
	// TODO
}

type InventoryItem struct {
	//
	//inventoryItemId: inventoryItemId,
	//
	//productId: productId,
	//
	//quantityOnHand: quantityOnHand,
	//
	//reorderThreshold: reorderThreshold,
	//
	//reorderQuantity: reorderQuantity,
	//
	//lastRestockedDate: lastRestockedDate,
	//
	// TODO
}

type Supplier struct {
	//
	//supplierId: supplierId,
	//
	//name: name,
	//
	//contactName: contactName,
	//
	//email: email,
	//
	//phoneNumber: phoneNumber,
	//
	//productCategories: productCategories,
	//
	//status: status,
	//
	// TODO
}

type PurchaseOrder struct {
	//
	//purchaseOrderId: purchaseOrderId,
	//
	//supplierId: supplierId,
	//
	//orderDate: orderDate,
	//
	//expectedDeliveryDate: expectedDeliveryDate,
	//
	//lineItems: lineItems,
	//
	//totalCost: totalCost,
	//
	//status: status,
	//
	// TODO
}
