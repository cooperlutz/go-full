import { describe, it, expect, vi } from "vitest";
import { usePingPongMetrics } from "../usePingPongMetrics";

// Mock MetricsApi methods
vi.mock("~/pingpong/services", () => {
  return {
    MetricsApi: vi.fn().mockImplementation(() => ({
      getTotalPingPongs: vi.fn(),
      getTotalPings: vi.fn(),
      getTotalPongs: vi.fn(),
      getDailyDistribution: vi.fn(),
    })),
  };
});

// Mock BackendConfig (not used directly, but required for MetricsApi instantiation)
vi.mock("~/pingpong/config", () => ({
  BackendConfig: {},
}));

describe("usePingPongMetrics", () => {
  it("should initialize refs with default values", () => {
    const composable = usePingPongMetrics();
    expect(composable.totalNumberOfPingPongs.value).toBe(0);
    expect(composable.totalNumberOfPings.value).toBe(0);
    expect(composable.totalNumberOfPongs.value).toBe(0);
    expect(composable.totalNumberOfPingPongsPerDay.value).toEqual({
      labels: [],
      values: [],
    });
    expect(composable.state.value).toEqual({ error: null, loading: false });
  });
});
