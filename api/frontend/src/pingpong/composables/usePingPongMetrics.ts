import { ref } from "vue";
import { BackendConfig } from "~/pingpong/config";
import { MetricsApi } from "~/pingpong/services";

const metricsAPI = new MetricsApi(BackendConfig);

function usePingPongMetrics() {
  const totalNumberOfPingPongs = ref(0);
  const totalNumberOfPings = ref(0);
  const totalNumberOfPongs = ref(0);
  const totalNumberOfPingPongsPerDay = ref<{
    labels: string[];
    values: number[];
  }>({ labels: [], values: [] });

  const state = ref({
    // Define your state properties here
    error: null as Error | null,
    loading: false,
  });

  const actions = {
    // Define your actions here
    fetchData: async () => {
      state.value.loading = true;
      state.value.error = null;
      try {
        // getTotalPingPongs
        const getTotalPingPongsResponse = await metricsAPI.getTotalPingPongs();
        totalNumberOfPingPongs.value = getTotalPingPongsResponse;
        // getTotalPings
        const getTotalPingsResponse = await metricsAPI.getTotalPings();
        totalNumberOfPings.value = getTotalPingsResponse;
        // getTotalPongs
        const getTotalPongsResponse = await metricsAPI.getTotalPongs();
        totalNumberOfPongs.value = getTotalPongsResponse;
        // getDailyDistribution
        const getdailyDistributionResponse =
          await metricsAPI.getDailyDistribution();
        totalNumberOfPingPongsPerDay.value = {
          labels: getdailyDistributionResponse.dimensionKeys || [],
          values:
            getdailyDistributionResponse.dimensionValues?.map(Number) || [],
        };
      } catch (err) {
        if (err instanceof Error) {
          state.value.error = err;
        }
      } finally {
        state.value.loading = false;
      }
    },
  };

  return {
    totalNumberOfPingPongs,
    totalNumberOfPings,
    totalNumberOfPongs,
    totalNumberOfPingPongsPerDay,
    state,
    actions,
  };
}

export { usePingPongMetrics };
