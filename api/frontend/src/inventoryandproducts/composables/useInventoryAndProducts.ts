import { ref } from "vue";

import { BackendConfig } from "../config";
import {
  DefaultApi,
  type Product,
  type FindOneProductRequest,
  type InventoryItem,
  type FindOneInventoryItemRequest,
  type Supplier,
  type FindOneSupplierRequest,
  type PurchaseOrder,
  type FindOnePurchaseOrderRequest,
} from "../services";

const inventoryandproductsAPI = new DefaultApi(BackendConfig);

export function useFindAllProducts() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const products = ref<Array<Product> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await inventoryandproductsAPI.findAllProducts();
      products.value = response;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    error,
    loading,
    products,
    getFindAll,
  };
}

export function useFindOneProducts() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const product = ref<Product | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneProductRequest = {
        productId: id,
      };
      const response = await inventoryandproductsAPI.findOneProduct(req);
      product.value = response;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    error,
    loading,
    product,
    getFindOne,
  };
}

export function useFindAllInventoryItems() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const inventoryitems = ref<Array<InventoryItem> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await inventoryandproductsAPI.findAllInventoryItems();
      inventoryitems.value = response;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    error,
    loading,
    inventoryitems,
    getFindAll,
  };
}

export function useFindOneInventoryItems() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const inventoryitem = ref<InventoryItem | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneInventoryItemRequest = {
        inventoryItemId: id,
      };
      const response = await inventoryandproductsAPI.findOneInventoryItem(req);
      inventoryitem.value = response;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    error,
    loading,
    inventoryitem,
    getFindOne,
  };
}

export function useFindAllSuppliers() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const suppliers = ref<Array<Supplier> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await inventoryandproductsAPI.findAllSuppliers();
      suppliers.value = response;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    error,
    loading,
    suppliers,
    getFindAll,
  };
}

export function useFindOneSuppliers() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const supplier = ref<Supplier | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneSupplierRequest = {
        supplierId: id,
      };
      const response = await inventoryandproductsAPI.findOneSupplier(req);
      supplier.value = response;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    error,
    loading,
    supplier,
    getFindOne,
  };
}

export function useFindAllPurchaseOrders() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const purchaseorders = ref<Array<PurchaseOrder> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await inventoryandproductsAPI.findAllPurchaseOrders();
      purchaseorders.value = response;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    error,
    loading,
    purchaseorders,
    getFindAll,
  };
}

export function useFindOnePurchaseOrders() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const purchaseorder = ref<PurchaseOrder | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOnePurchaseOrderRequest = {
        purchaseOrderId: id,
      };
      const response = await inventoryandproductsAPI.findOnePurchaseOrder(req);
      purchaseorder.value = response;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    error,
    loading,
    purchaseorder,
    getFindOne,
  };
}

// export function useAddProductToCatalogs() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const addProductToCatalog = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: AddProductToCatalogRequest = {
//         Id: id,
//      };
//      await inventoryandproductsAPI.addProductToCatalog(req);
//    } catch (err) {
//      if (err instanceof Error) {
//        error.value = err;
//      }
//    } finally {
//      loading.value = false;
//    }
//  };
//
//  return {
//    error,
//    loading,
//    addProductToCatalog,
//  };
// }

// export function useUpdateProductPrices() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const updateProductPrice = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: UpdateProductPriceRequest = {
//         Id: id,
//      };
//      await inventoryandproductsAPI.updateProductPrice(req);
//    } catch (err) {
//      if (err instanceof Error) {
//        error.value = err;
//      }
//    } finally {
//      loading.value = false;
//    }
//  };
//
//  return {
//    error,
//    loading,
//    updateProductPrice,
//  };
// }

// export function useRestockInventorys() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const restockInventory = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: RestockInventoryRequest = {
//         Id: id,
//      };
//      await inventoryandproductsAPI.restockInventory(req);
//    } catch (err) {
//      if (err instanceof Error) {
//        error.value = err;
//      }
//    } finally {
//      loading.value = false;
//    }
//  };
//
//  return {
//    error,
//    loading,
//    restockInventory,
//  };
// }

// export function useDeductInventorys() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const deductInventory = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: DeductInventoryRequest = {
//         Id: id,
//      };
//      await inventoryandproductsAPI.deductInventory(req);
//    } catch (err) {
//      if (err instanceof Error) {
//        error.value = err;
//      }
//    } finally {
//      loading.value = false;
//    }
//  };
//
//  return {
//    error,
//    loading,
//    deductInventory,
//  };
// }

// export function usePlacePurchaseOrders() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const placePurchaseOrder = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: PlacePurchaseOrderRequest = {
//         Id: id,
//      };
//      await inventoryandproductsAPI.placePurchaseOrder(req);
//    } catch (err) {
//      if (err instanceof Error) {
//        error.value = err;
//      }
//    } finally {
//      loading.value = false;
//    }
//  };
//
//  return {
//    error,
//    loading,
//    placePurchaseOrder,
//  };
// }

// export function useReceivePurchaseOrders() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const receivePurchaseOrder = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: ReceivePurchaseOrderRequest = {
//         Id: id,
//      };
//      await inventoryandproductsAPI.receivePurchaseOrder(req);
//    } catch (err) {
//      if (err instanceof Error) {
//        error.value = err;
//      }
//    } finally {
//      loading.value = false;
//    }
//  };
//
//  return {
//    error,
//    loading,
//    receivePurchaseOrder,
//  };
// }
