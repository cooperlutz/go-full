import { ref } from "vue";

import { BackendConfig } from "../config";
import {
  DefaultApi,
  type Appointment,
  type FindOneAppointmentRequest,
  type TelemedicineSession,
  type FindOneTelemedicineSessionRequest,
} from "../services";

const appointmentschedulingAPI = new DefaultApi(BackendConfig);

export function useFindAllAppointments() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const appointments = ref<Array<Appointment> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await appointmentschedulingAPI.findAllAppointments();
      appointments.value = response;
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
    appointments,
    getFindAll,
  };
}

export function useFindOneAppointments() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const appointment = ref<Appointment | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneAppointmentRequest = {
        appointmentId: id,
      };
      const response = await appointmentschedulingAPI.findOneAppointment(req);
      appointment.value = response;
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
    appointment,
    getFindOne,
  };
}

export function useFindAllTelemedicineSessions() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const telemedicinesessions = ref<Array<TelemedicineSession> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response =
        await appointmentschedulingAPI.findAllTelemedicineSessions();
      telemedicinesessions.value = response;
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
    telemedicinesessions,
    getFindAll,
  };
}

export function useFindOneTelemedicineSessions() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const telemedicinesession = ref<TelemedicineSession | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneTelemedicineSessionRequest = {
        telemedicineSessionId: id,
      };
      const response =
        await appointmentschedulingAPI.findOneTelemedicineSession(req);
      telemedicinesession.value = response;
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
    telemedicinesession,
    getFindOne,
  };
}

// export function useScheduleAppointments() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const scheduleAppointment = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: ScheduleAppointmentRequest = {
//         Id: id,
//      };
//      await appointmentschedulingAPI.scheduleAppointment(req);
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
//    scheduleAppointment,
//  };
// }

// export function useRescheduleAppointments() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const rescheduleAppointment = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: RescheduleAppointmentRequest = {
//         Id: id,
//      };
//      await appointmentschedulingAPI.rescheduleAppointment(req);
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
//    rescheduleAppointment,
//  };
// }

// export function useCancelAppointments() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const cancelAppointment = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: CancelAppointmentRequest = {
//         Id: id,
//      };
//      await appointmentschedulingAPI.cancelAppointment(req);
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
//    cancelAppointment,
//  };
// }

// export function useConfirmAppointments() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const confirmAppointment = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: ConfirmAppointmentRequest = {
//         Id: id,
//      };
//      await appointmentschedulingAPI.confirmAppointment(req);
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
//    confirmAppointment,
//  };
// }

// export function useCompleteAppointments() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const completeAppointment = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: CompleteAppointmentRequest = {
//         Id: id,
//      };
//      await appointmentschedulingAPI.completeAppointment(req);
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
//    completeAppointment,
//  };
// }

// export function useStartTelemedicineSessions() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const startTelemedicineSession = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: StartTelemedicineSessionRequest = {
//         Id: id,
//      };
//      await appointmentschedulingAPI.startTelemedicineSession(req);
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
//    startTelemedicineSession,
//  };
// }

// export function useEndTelemedicineSessions() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const endTelemedicineSession = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: EndTelemedicineSessionRequest = {
//         Id: id,
//      };
//      await appointmentschedulingAPI.endTelemedicineSession(req);
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
//    endTelemedicineSession,
//  };
// }
