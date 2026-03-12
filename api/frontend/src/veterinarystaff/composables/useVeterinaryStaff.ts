import { ref } from "vue";

import { BackendConfig } from "../config";
import {
  DefaultApi,
  type Veterinarian,
  type FindOneVeterinarianRequest,
  type StaffMember,
  type FindOneStaffMemberRequest,
  type AvailabilitySchedule,
  type FindOneAvailabilityScheduleRequest,
} from "../services";

const veterinarystaffAPI = new DefaultApi(BackendConfig);

export function useFindAllVeterinarians() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const veterinarians = ref<Array<Veterinarian> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await veterinarystaffAPI.findAllVeterinarians();
      veterinarians.value = response;
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
    veterinarians,
    getFindAll,
  };
}

export function useFindOneVeterinarians() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const veterinarian = ref<Veterinarian | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneVeterinarianRequest = {
        veterinarianId: id,
      };
      const response = await veterinarystaffAPI.findOneVeterinarian(req);
      veterinarian.value = response;
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
    veterinarian,
    getFindOne,
  };
}

export function useFindAllStaffMembers() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const staffmembers = ref<Array<StaffMember> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await veterinarystaffAPI.findAllStaffMembers();
      staffmembers.value = response;
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
    staffmembers,
    getFindAll,
  };
}

export function useFindOneStaffMembers() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const staffmember = ref<StaffMember | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneStaffMemberRequest = {
        staffMemberId: id,
      };
      const response = await veterinarystaffAPI.findOneStaffMember(req);
      staffmember.value = response;
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
    staffmember,
    getFindOne,
  };
}

export function useFindAllAvailabilitySchedules() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const availabilityschedules = ref<Array<AvailabilitySchedule> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await veterinarystaffAPI.findAllAvailabilitySchedules();
      availabilityschedules.value = response;
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
    availabilityschedules,
    getFindAll,
  };
}

export function useFindOneAvailabilitySchedules() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const availabilityschedule = ref<AvailabilitySchedule | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneAvailabilityScheduleRequest = {
        availabilityScheduleId: id,
      };
      const response =
        await veterinarystaffAPI.findOneAvailabilitySchedule(req);
      availabilityschedule.value = response;
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
    availabilityschedule,
    getFindOne,
  };
}

// export function useOnboardVeterinarians() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const onboardVeterinarian = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: OnboardVeterinarianRequest = {
//         Id: id,
//      };
//      await veterinarystaffAPI.onboardVeterinarian(req);
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
//    onboardVeterinarian,
//  };
// }

// export function useUpdateVeterinarianProfiles() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const updateVeterinarianProfile = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: UpdateVeterinarianProfileRequest = {
//         Id: id,
//      };
//      await veterinarystaffAPI.updateVeterinarianProfile(req);
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
//    updateVeterinarianProfile,
//  };
// }

// export function useSetStaffAvailabilitys() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const setStaffAvailability = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: SetStaffAvailabilityRequest = {
//         Id: id,
//      };
//      await veterinarystaffAPI.setStaffAvailability(req);
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
//    setStaffAvailability,
//  };
// }

// export function useDeactivateStaffMembers() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const deactivateStaffMember = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: DeactivateStaffMemberRequest = {
//         Id: id,
//      };
//      await veterinarystaffAPI.deactivateStaffMember(req);
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
//    deactivateStaffMember,
//  };
// }
