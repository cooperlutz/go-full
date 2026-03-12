import { ref } from "vue";

import { BackendConfig } from "../config";
import {
  DefaultApi,
  type SalesOrder,
  type FindOneSalesOrderRequest,
  type ShoppingCart,
  type FindOneShoppingCartRequest,
} from "../services";

const retailsalesAPI = new DefaultApi(BackendConfig);

export function useFindAllSalesOrders() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const salesorders = ref<Array<SalesOrder> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await retailsalesAPI.findAllSalesOrders();
      salesorders.value = response;
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
    salesorders,
    getFindAll,
  };
}

export function useFindOneSalesOrders() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const salesorder = ref<SalesOrder | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneSalesOrderRequest = {
        salesOrderId: id,
      };
      const response = await retailsalesAPI.findOneSalesOrder(req);
      salesorder.value = response;
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
    salesorder,
    getFindOne,
  };
}

export function useFindAllShoppingCarts() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const shoppingcarts = ref<Array<ShoppingCart> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await retailsalesAPI.findAllShoppingCarts();
      shoppingcarts.value = response;
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
    shoppingcarts,
    getFindAll,
  };
}

export function useFindOneShoppingCarts() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const shoppingcart = ref<ShoppingCart | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneShoppingCartRequest = {
        shoppingCartId: id,
      };
      const response = await retailsalesAPI.findOneShoppingCart(req);
      shoppingcart.value = response;
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
    shoppingcart,
    getFindOne,
  };
}

// export function useAddItemToCarts() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const addItemToCart = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: AddItemToCartRequest = {
//         Id: id,
//      };
//      await retailsalesAPI.addItemToCart(req);
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
//    addItemToCart,
//  };
// }

// export function useRemoveItemFromCarts() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const removeItemFromCart = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: RemoveItemFromCartRequest = {
//         Id: id,
//      };
//      await retailsalesAPI.removeItemFromCart(req);
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
//    removeItemFromCart,
//  };
// }

// export function usePlaceSalesOrders() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const placeSalesOrder = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: PlaceSalesOrderRequest = {
//         Id: id,
//      };
//      await retailsalesAPI.placeSalesOrder(req);
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
//    placeSalesOrder,
//  };
// }

// export function useFulfillSalesOrders() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const fulfillSalesOrder = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: FulfillSalesOrderRequest = {
//         Id: id,
//      };
//      await retailsalesAPI.fulfillSalesOrder(req);
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
//    fulfillSalesOrder,
//  };
// }

// export function useCancelSalesOrders() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const cancelSalesOrder = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: CancelSalesOrderRequest = {
//         Id: id,
//      };
//      await retailsalesAPI.cancelSalesOrder(req);
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
//    cancelSalesOrder,
//  };
// }
