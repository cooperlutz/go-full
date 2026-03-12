import { ref } from "vue";

import { BackendConfig } from "../config";
import {
  DefaultApi,
  type Pet,
  type FindOnePetRequest,
  type MedicalRecord,
  type FindOneMedicalRecordRequest,
  type VaccinationRecord,
  type FindOneVaccinationRecordRequest,
} from "../services";

const patientmanagementAPI = new DefaultApi(BackendConfig);

export function useFindAllPets() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const pets = ref<Array<Pet> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await patientmanagementAPI.findAllPets();
      pets.value = response;
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
    pets,
    getFindAll,
  };
}

export function useFindOnePets() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const pet = ref<Pet | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOnePetRequest = {
        petId: id,
      };
      const response = await patientmanagementAPI.findOnePet(req);
      pet.value = response;
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
    pet,
    getFindOne,
  };
}

export function useFindAllMedicalRecords() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const medicalrecords = ref<Array<MedicalRecord> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await patientmanagementAPI.findAllMedicalRecords();
      medicalrecords.value = response;
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
    medicalrecords,
    getFindAll,
  };
}

export function useFindOneMedicalRecords() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const medicalrecord = ref<MedicalRecord | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneMedicalRecordRequest = {
        medicalRecordId: id,
      };
      const response = await patientmanagementAPI.findOneMedicalRecord(req);
      medicalrecord.value = response;
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
    medicalrecord,
    getFindOne,
  };
}

export function useFindAllVaccinationRecords() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const vaccinationrecords = ref<Array<VaccinationRecord> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await patientmanagementAPI.findAllVaccinationRecords();
      vaccinationrecords.value = response;
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
    vaccinationrecords,
    getFindAll,
  };
}

export function useFindOneVaccinationRecords() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const vaccinationrecord = ref<VaccinationRecord | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneVaccinationRecordRequest = {
        vaccinationRecordId: id,
      };
      const response = await patientmanagementAPI.findOneVaccinationRecord(req);
      vaccinationrecord.value = response;
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
    vaccinationrecord,
    getFindOne,
  };
}

// export function useRegisterPets() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const registerPet = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: RegisterPetRequest = {
//         Id: id,
//      };
//      await patientmanagementAPI.registerPet(req);
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
//    registerPet,
//  };
// }

// export function useUpdatePetDetailss() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const updatePetDetails = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: UpdatePetDetailsRequest = {
//         Id: id,
//      };
//      await patientmanagementAPI.updatePetDetails(req);
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
//    updatePetDetails,
//  };
// }

// export function useAddMedicalRecords() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const addMedicalRecord = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: AddMedicalRecordRequest = {
//         Id: id,
//      };
//      await patientmanagementAPI.addMedicalRecord(req);
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
//    addMedicalRecord,
//  };
// }

// export function useRecordVaccinations() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const recordVaccination = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: RecordVaccinationRequest = {
//         Id: id,
//      };
//      await patientmanagementAPI.recordVaccination(req);
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
//    recordVaccination,
//  };
// }

// export function useDeactivatePets() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const deactivatePet = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: DeactivatePetRequest = {
//         Id: id,
//      };
//      await patientmanagementAPI.deactivatePet(req);
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
//    deactivatePet,
//  };
// }
