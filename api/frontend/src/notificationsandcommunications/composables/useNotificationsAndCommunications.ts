import { ref } from "vue";

import { BackendConfig } from "../config";
import {
  DefaultApi,
  type Notification,
  type FindOneNotificationRequest,
  type NotificationTemplate,
  type FindOneNotificationTemplateRequest,
} from "../services";

const notificationsandcommunicationsAPI = new DefaultApi(BackendConfig);

export function useFindAllNotifications() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const notifications = ref<Array<Notification> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response =
        await notificationsandcommunicationsAPI.findAllNotifications();
      notifications.value = response;
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
    notifications,
    getFindAll,
  };
}

export function useFindOneNotifications() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const notification = ref<Notification | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneNotificationRequest = {
        notificationId: id,
      };
      const response =
        await notificationsandcommunicationsAPI.findOneNotification(req);
      notification.value = response;
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
    notification,
    getFindOne,
  };
}

export function useFindAllNotificationTemplates() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const notificationtemplates = ref<Array<NotificationTemplate> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response =
        await notificationsandcommunicationsAPI.findAllNotificationTemplates();
      notificationtemplates.value = response;
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
    notificationtemplates,
    getFindAll,
  };
}

export function useFindOneNotificationTemplates() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const notificationtemplate = ref<NotificationTemplate | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneNotificationTemplateRequest = {
        notificationTemplateId: id,
      };
      const response =
        await notificationsandcommunicationsAPI.findOneNotificationTemplate(
          req,
        );
      notificationtemplate.value = response;
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
    notificationtemplate,
    getFindOne,
  };
}

// export function useSendNotifications() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const sendNotification = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: SendNotificationRequest = {
//         Id: id,
//      };
//      await notificationsandcommunicationsAPI.sendNotification(req);
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
//    sendNotification,
//  };
// }

// export function useScheduleNotifications() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const scheduleNotification = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: ScheduleNotificationRequest = {
//         Id: id,
//      };
//      await notificationsandcommunicationsAPI.scheduleNotification(req);
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
//    scheduleNotification,
//  };
// }

// export function useCancelScheduledNotifications() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const cancelScheduledNotification = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: CancelScheduledNotificationRequest = {
//         Id: id,
//      };
//      await notificationsandcommunicationsAPI.cancelScheduledNotification(req);
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
//    cancelScheduledNotification,
//  };
// }

// export function useCreateNotificationTemplates() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const createNotificationTemplate = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: CreateNotificationTemplateRequest = {
//         Id: id,
//      };
//      await notificationsandcommunicationsAPI.createNotificationTemplate(req);
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
//    createNotificationTemplate,
//  };
// }
