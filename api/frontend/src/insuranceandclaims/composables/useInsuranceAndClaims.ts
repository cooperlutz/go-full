import { ref } from "vue";

import { BackendConfig } from "../config";
import {
  DefaultApi,
  type InsuranceProvider,
  type FindOneInsuranceProviderRequest,
  type InsuranceClaim,
  type FindOneInsuranceClaimRequest,
} from "../services";

const insuranceandclaimsAPI = new DefaultApi(BackendConfig);

export function useFindAllInsuranceProviders() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const insuranceproviders = ref<Array<InsuranceProvider> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await insuranceandclaimsAPI.findAllInsuranceProviders();
      insuranceproviders.value = response;
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
    insuranceproviders,
    getFindAll,
  };
}

export function useFindOneInsuranceProviders() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const insuranceprovider = ref<InsuranceProvider | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneInsuranceProviderRequest = {
        insuranceProviderId: id,
      };
      const response =
        await insuranceandclaimsAPI.findOneInsuranceProvider(req);
      insuranceprovider.value = response;
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
    insuranceprovider,
    getFindOne,
  };
}

export function useFindAllInsuranceClaims() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const insuranceclaims = ref<Array<InsuranceClaim> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await insuranceandclaimsAPI.findAllInsuranceClaims();
      insuranceclaims.value = response;
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
    insuranceclaims,
    getFindAll,
  };
}

export function useFindOneInsuranceClaims() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const insuranceclaim = ref<InsuranceClaim | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneInsuranceClaimRequest = {
        insuranceClaimId: id,
      };
      const response = await insuranceandclaimsAPI.findOneInsuranceClaim(req);
      insuranceclaim.value = response;
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
    insuranceclaim,
    getFindOne,
  };
}

// export function useRegisterInsuranceProviders() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const registerInsuranceProvider = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: RegisterInsuranceProviderRequest = {
//         Id: id,
//      };
//      await insuranceandclaimsAPI.registerInsuranceProvider(req);
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
//    registerInsuranceProvider,
//  };
// }

// export function useSubmitInsuranceClaims() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const submitInsuranceClaim = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: SubmitInsuranceClaimRequest = {
//         Id: id,
//      };
//      await insuranceandclaimsAPI.submitInsuranceClaim(req);
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
//    submitInsuranceClaim,
//  };
// }

// export function useApproveInsuranceClaims() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const approveInsuranceClaim = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: ApproveInsuranceClaimRequest = {
//         Id: id,
//      };
//      await insuranceandclaimsAPI.approveInsuranceClaim(req);
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
//    approveInsuranceClaim,
//  };
// }

// export function useRejectInsuranceClaims() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const rejectInsuranceClaim = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: RejectInsuranceClaimRequest = {
//         Id: id,
//      };
//      await insuranceandclaimsAPI.rejectInsuranceClaim(req);
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
//    rejectInsuranceClaim,
//  };
// }
