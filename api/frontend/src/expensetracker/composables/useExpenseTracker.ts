import { ref } from "vue";

import { BackendConfig } from "../config";
import { 
  DefaultApi,
  
  type Expense, 
  type FindOneExpenseRequest,
  
} from "../services";

const expensetrackerAPI = new DefaultApi(BackendConfig);


export function useFindAllExpenses() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const expenses = ref<Array<Expense> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await expensetrackerAPI.findAllExpenses();
      expenses.value = response;
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
    expenses,
    getFindAll,
  };
}


export function useFindOneExpenses() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const expense = ref<Expense | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneExpenseRequest = {
         expenseId: id,
      };
      const response = await expensetrackerAPI.findOneExpense(req);
      expense.value = response;
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
    expense,
    getFindOne,
  };
}



//export function useSubmitExpenses() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const submitExpense = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: SubmitExpenseRequest = {
//         Id: id,
//      };
//      await expensetrackerAPI.submitExpense(req);
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
//    submitExpense,
//  };
//}
