import { ref } from "vue";

import { BackendConfig } from "~/pingpong/config";
import {
  PingpongApi,
  PingpongsApi,
  type PingPongRequest,
  type PingPong,
  type PingPongsRaw,
} from "~/pingpong/services";

const pingpongAPI = new PingpongApi(BackendConfig);
const pingpongsAPI = new PingpongsApi(BackendConfig);

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
  const allPingPongs = ref<PingPongsRaw>();
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const fetchData = async () => {
    loading.value = true;
    error.value = null;
    try {
      // getAll
      const getAll = await pingpongsAPI.getFindAllPingPongs();
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
