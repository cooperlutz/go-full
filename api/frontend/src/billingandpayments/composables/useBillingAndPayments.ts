import { ref } from "vue";

import { BackendConfig } from "../config";
import {
  DefaultApi,
  type Invoice,
  type FindOneInvoiceRequest,
  type Payment,
  type FindOnePaymentRequest,
  type Refund,
  type FindOneRefundRequest,
} from "../services";

const billingandpaymentsAPI = new DefaultApi(BackendConfig);

export function useFindAllInvoices() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const invoices = ref<Array<Invoice> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await billingandpaymentsAPI.findAllInvoices();
      invoices.value = response;
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
    invoices,
    getFindAll,
  };
}

export function useFindOneInvoices() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const invoice = ref<Invoice | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneInvoiceRequest = {
        invoiceId: id,
      };
      const response = await billingandpaymentsAPI.findOneInvoice(req);
      invoice.value = response;
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
    invoice,
    getFindOne,
  };
}

export function useFindAllPayments() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const payments = ref<Array<Payment> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await billingandpaymentsAPI.findAllPayments();
      payments.value = response;
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
    payments,
    getFindAll,
  };
}

export function useFindOnePayments() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const payment = ref<Payment | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOnePaymentRequest = {
        paymentId: id,
      };
      const response = await billingandpaymentsAPI.findOnePayment(req);
      payment.value = response;
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
    payment,
    getFindOne,
  };
}

export function useFindAllRefunds() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const refunds = ref<Array<Refund> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await billingandpaymentsAPI.findAllRefunds();
      refunds.value = response;
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
    refunds,
    getFindAll,
  };
}

export function useFindOneRefunds() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const refund = ref<Refund | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneRefundRequest = {
        refundId: id,
      };
      const response = await billingandpaymentsAPI.findOneRefund(req);
      refund.value = response;
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
    refund,
    getFindOne,
  };
}

// export function useGenerateInvoices() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const generateInvoice = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: GenerateInvoiceRequest = {
//         Id: id,
//      };
//      await billingandpaymentsAPI.generateInvoice(req);
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
//    generateInvoice,
//  };
// }

// export function useApplyDiscountToInvoices() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const applyDiscountToInvoice = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: ApplyDiscountToInvoiceRequest = {
//         Id: id,
//      };
//      await billingandpaymentsAPI.applyDiscountToInvoice(req);
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
//    applyDiscountToInvoice,
//  };
// }

// export function useProcessPayments() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const processPayment = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: ProcessPaymentRequest = {
//         Id: id,
//      };
//      await billingandpaymentsAPI.processPayment(req);
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
//    processPayment,
//  };
// }

// export function useIssueRefunds() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const issueRefund = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: IssueRefundRequest = {
//         Id: id,
//      };
//      await billingandpaymentsAPI.issueRefund(req);
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
//    issueRefund,
//  };
// }

// export function useVoidInvoices() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const voidInvoice = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: VoidInvoiceRequest = {
//         Id: id,
//      };
//      await billingandpaymentsAPI.voidInvoice(req);
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
//    voidInvoice,
//  };
// }
