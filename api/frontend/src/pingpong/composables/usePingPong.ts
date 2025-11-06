import { ref } from "vue";

import { BackendConfig } from "~/pingpong/config";
import {
  PingpongApi,
  type GetFindOneByIDRequest,
  type PingPongRequest,
  type PingPong,
  // type PingPongRaw,
  type PingPongsRaw,
} from "~/pingpong/services";

const pingpongAPI = new PingpongApi(BackendConfig);

// Composable for sending a ping pong message
export function useSendPingPong() {
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const sendPingPong = async (msg: string) => {
    loading.value = true;
    error.value = null;
    const data: PingPong = {
      message: msg,
    };
    const req: PingPongRequest = {
      pingPong: data,
    };
    try {
      const response = await pingpongAPI.pingPong(req);
      return response;
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
    sendPingPong,
  };
}

// Composable for fetching all ping pong entities
export function useFindAllPingPongs() {
  const allPingPongs = ref<PingPongsRaw | null>(null);
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const fetchData = async () => {
    loading.value = true;
    error.value = null;
    try {
      // getAll
      const getAll = await pingpongAPI.getFindAllPingPongs();
      allPingPongs.value = getAll;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    allPingPongs,
    error,
    loading,
    fetchData,
  };
}

/* STEP 5.2. Implement Frontend Composable
here we define the composable which will be used to interact with our REST API
the composable will be tested via the component tests in conjunction with the MSW endpoint
*/
export function useFindPingPongByID() {
  // initialize a null error and set our loading state to false
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const lookup = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: GetFindOneByIDRequest = {
        pingPongID: id,
      };
      const getByID = await pingpongAPI.getFindOneByID(req);
      return getByID;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    lookup,
    error,
    loading,
  };
}
