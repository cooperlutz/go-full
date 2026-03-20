import { ref } from "vue";

import { BackendConfig } from "../config";
import { 
  DefaultApi,
  
  type StartupPitch, 
  type FindOneStartupPitchRequest,
  
} from "../services";

const startupidearaterAPI = new DefaultApi(BackendConfig);


export function useFindAllStartupPitchs() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const startuppitchs = ref<Array<StartupPitch> | null>(null);

  const getFindAll = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await startupidearaterAPI.findAllStartupPitchs();
      startuppitchs.value = response;
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
    startuppitchs,
    getFindAll,
  };
}


export function useFindOneStartupPitchs() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const startuppitch = ref<StartupPitch | null>(null);

  const getFindOne = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: FindOneStartupPitchRequest = {
         startupPitchId: id,
      };
      const response = await startupidearaterAPI.findOneStartupPitch(req);
      startuppitch.value = response;
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
    startuppitch,
    getFindOne,
  };
}



//export function useRateStartupIdeas() {
//  const error = ref<Error | null>(null);
//  const loading = ref(false);
//
//  const rateStartupIdea = async () => {
//    loading.value = true;
//    error.value = null;
//    try {
//      const req: RateStartupIdeaRequest = {
//         Id: id,
//      };
//      await startupidearaterAPI.rateStartupIdea(req);
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
//    rateStartupIdea,
//  };
//}
