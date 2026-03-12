import { ref } from "vue";

import { BackendConfig } from "../config";
import {
  DefaultApi,
  type Owner,
  type FindOneOwnerRequest,
  type LoyaltyAccount,
  type FindOneLoyaltyAccountRequest,
} from "../services";

const ownermanagementAPI = new DefaultApi(BackendConfig);

export function useFindAllOwners() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const owners = ref<Array<Owner> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await ownermanagementAPI.findAllOwners();
      owners.value = response;
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
    owners,
    getFindAll,
  };
}

export function useFindOneOwners() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const owner = ref<Owner | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneOwnerRequest = {
        ownerId: id,
      };
      const response = await ownermanagementAPI.findOneOwner(req);
      owner.value = response;
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
    owner,
    getFindOne,
  };
}

export function useFindAllLoyaltyAccounts() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const loyaltyaccounts = ref<Array<LoyaltyAccount> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await ownermanagementAPI.findAllLoyaltyAccounts();
      loyaltyaccounts.value = response;
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
    loyaltyaccounts,
    getFindAll,
  };
}

export function useFindOneLoyaltyAccounts() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const loyaltyaccount = ref<LoyaltyAccount | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneLoyaltyAccountRequest = {
        loyaltyAccountId: id,
      };
      const response = await ownermanagementAPI.findOneLoyaltyAccount(req);
      loyaltyaccount.value = response;
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
    loyaltyaccount,
    getFindOne,
  };
}

// export function useRegisterOwners() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const registerOwner = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: RegisterOwnerRequest = {
//         Id: id,
//      };
//      await ownermanagementAPI.registerOwner(req);
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
//    registerOwner,
//  };
// }

// export function useUpdateOwnerProfiles() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const updateOwnerProfile = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: UpdateOwnerProfileRequest = {
//         Id: id,
//      };
//      await ownermanagementAPI.updateOwnerProfile(req);
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
//    updateOwnerProfile,
//  };
// }

// export function useEnrollInLoyaltyPrograms() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const enrollInLoyaltyProgram = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: EnrollInLoyaltyProgramRequest = {
//         Id: id,
//      };
//      await ownermanagementAPI.enrollInLoyaltyProgram(req);
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
//    enrollInLoyaltyProgram,
//  };
// }

// export function useAwardLoyaltyPointss() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const awardLoyaltyPoints = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: AwardLoyaltyPointsRequest = {
//         Id: id,
//      };
//      await ownermanagementAPI.awardLoyaltyPoints(req);
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
//    awardLoyaltyPoints,
//  };
// }

// export function useRedeemLoyaltyPointss() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const redeemLoyaltyPoints = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: RedeemLoyaltyPointsRequest = {
//         Id: id,
//      };
//      await ownermanagementAPI.redeemLoyaltyPoints(req);
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
//    redeemLoyaltyPoints,
//  };
// }

// export function useDeactivateOwners() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const deactivateOwner = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: DeactivateOwnerRequest = {
//         Id: id,
//      };
//      await ownermanagementAPI.deactivateOwner(req);
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
//    deactivateOwner,
//  };
// }
