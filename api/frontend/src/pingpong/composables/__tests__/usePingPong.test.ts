import { describe, it, expect, vi } from "vitest";
import { useFindAllPingPongs } from "../usePingPong";

// Mock BackendConfig (not used directly, but required for MetricsApi instantiation)
vi.mock("~/pingpong/config", () => ({
  BackendConfig: {},
}));

describe("usePingPong", () => {
  it("should initialize refs with default values", async () => {
    const { allPingPongs, loading, error } = useFindAllPingPongs();
    expect(allPingPongs.value).toBe(null);
    expect(loading.value).toBe(false);
    expect(error.value).toEqual(null);
  });

  it("should test successful fetchData", async () => {
    const { allPingPongs, loading, fetchData } = useFindAllPingPongs();

    await fetchData();

    expect(allPingPongs.value).toBeDefined();
    expect(loading.value).toBe(false);
  });

  it("should handle API errors in fetchData", async () => {
    const { allPingPongs, loading, error, fetchData } = useFindAllPingPongs();

    await fetchData();

    expect(allPingPongs.value).toBe(null);
    expect(loading.value).toBe(false);
    expect(error.value).toBeDefined();
  });

  it("should set loading to true during fetchData execution", async () => {
    const { loading, fetchData } = useFindAllPingPongs();

    const fetchPromise = fetchData();
    expect(loading.value).toBe(true);

    await fetchPromise;
    expect(loading.value).toBe(false);
  });
});
